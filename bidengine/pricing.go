package bidengine

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/akash-network/akash-api/go/node/types/unit"
	"github.com/akash-network/node/sdl"

	"github.com/akash-network/provider/cluster/util"
	"github.com/akash-network/provider/spheron/entities"
)

type Request struct {
	Owner          string `json:"owner"`
	DeploymentSpec *entities.DeploymentSpec
	PricePrecision int
}

const (
	DefaultPricePrecision = 6
)

type BidPricingStrategy interface {
	CalculatePrice(ctx context.Context, req Request) (uint64, error)
}

var (
	errAllScalesZero               = errors.New("at least one bid price must be a non-zero number")
	errNoPriceScaleForStorageClass = errors.New("no pricing configured for storage class")
	errScaleNegative               = errors.New("scale price cannot be negative")
)

type Storage map[string]decimal.Decimal

func (ss Storage) IsAnyZero() bool {
	if len(ss) == 0 {
		return true
	}

	for _, val := range ss {
		if val.IsZero() {
			return true
		}
	}

	return false
}

func (ss Storage) IsAnyNegative() bool {
	for _, val := range ss {
		if val.IsNegative() {
			return true
		}
	}

	return false
}

// AllLessThenOrEqual check all storage classes fit into max limits
// note better have dedicated limits for each class
func (ss Storage) AllLessThenOrEqual(val decimal.Decimal) bool {
	for _, storage := range ss {
		if !storage.LessThanOrEqual(val) {
			return false
		}
	}

	return true
}

type scalePricing struct {
	cpuScale      decimal.Decimal
	memoryScale   decimal.Decimal
	storageScale  Storage
	endpointScale decimal.Decimal
	ipScale       decimal.Decimal
}

func MakeScalePricing(
	cpuScale decimal.Decimal,
	memoryScale decimal.Decimal,
	storageScale Storage,
	endpointScale decimal.Decimal,
	ipScale decimal.Decimal) (BidPricingStrategy, error) {

	if cpuScale.IsZero() && memoryScale.IsZero() && storageScale.IsAnyZero() && endpointScale.IsZero() && ipScale.IsZero() {
		return nil, errAllScalesZero
	}

	if cpuScale.IsNegative() || memoryScale.IsNegative() || storageScale.IsAnyNegative() || endpointScale.IsNegative() ||
		ipScale.IsNegative() {
		return nil, errScaleNegative
	}

	result := scalePricing{
		cpuScale:      cpuScale,
		memoryScale:   memoryScale,
		storageScale:  storageScale,
		endpointScale: endpointScale,
		ipScale:       ipScale,
	}

	return result, nil
}

var (
	ErrBidQuantityInvalid = errors.New("A bid quantity is invalid")
	ErrBidZero            = errors.New("A bid of zero was produced")
)

func ceilBigRatToBigInt(v *big.Rat) *big.Int {
	numerator := v.Num()
	denom := v.Denom()

	result := big.NewInt(0).Div(numerator, denom)
	if !v.IsInt() {
		result.Add(result, big.NewInt(1))
	}

	return result
}

func (fp scalePricing) CalculatePrice(_ context.Context, req Request) (uint64, error) {
	// Use unlimited precision math here.
	// Otherwise, a correctly crafted order could create a cost of '1' given
	// a possible configuration
	cpuTotal := decimal.NewFromInt(0)
	memoryTotal := decimal.NewFromInt(0)
	storageTotal := make(Storage)
	denom := req.DeploymentSpec.Price()

	for k := range fp.storageScale {
		storageTotal[k] = decimal.NewFromInt(0)
	}

	endpointTotal := decimal.NewFromInt(0)
	ipTotal := decimal.NewFromInt(0).Add(fp.ipScale)
	ipTotal = ipTotal.Mul(decimal.NewFromInt(int64(util.GetEndpointQuantityOfResourceGroup(req.DeploymentSpec, entities.Endpoint_LEASED_IP))))

	// iterate over everything & sum it up
	for _, group := range req.DeploymentSpec.Resources {
		groupCount := decimal.NewFromInt(int64(group.Count)) // Expand uint32 to int64

		cpuQuantity := decimal.NewFromInt(int64(group.Resources.CPU.Units))
		cpuQuantity = cpuQuantity.Mul(groupCount)
		cpuTotal = cpuTotal.Add(cpuQuantity)

		memoryQuantity := decimal.NewFromInt(int64(group.Resources.Memory.Units))
		memoryQuantity = memoryQuantity.Mul(groupCount)
		memoryTotal = memoryTotal.Add(memoryQuantity)

		for _, storage := range group.Resources.Storage {
			storageQuantity := decimal.NewFromInt(int64(storage.Units))
			storageQuantity = storageQuantity.Mul(groupCount)

			storageClass := sdl.StorageEphemeral
			attr := storage.Attributes.Find(sdl.StorageAttributePersistent)
			if isPersistent, _ := attr.AsBool(); isPersistent {
				attr = storage.Attributes.Find(sdl.StorageAttributeClass)
				if class, set := attr.AsString(); set {
					storageClass = class
				}
			}

			total, exists := storageTotal[storageClass]

			if !exists {
				return 0, errors.Wrapf(errNoPriceScaleForStorageClass, storageClass)
			}

			total = total.Add(storageQuantity)

			storageTotal[storageClass] = total
		}

		endpointQuantity := decimal.NewFromInt(int64(len(group.Resources.Endpoints)))
		endpointTotal = endpointTotal.Add(endpointQuantity)
	}

	cpuTotal = cpuTotal.Mul(fp.cpuScale)

	mebibytes := decimal.NewFromInt(unit.Mi)

	memoryTotal = memoryTotal.Div(mebibytes)
	memoryTotal = memoryTotal.Mul(fp.memoryScale)

	for class, total := range storageTotal {
		total = total.Div(mebibytes)

		// at this point presence of class in storageScale has been validated
		total = total.Mul(fp.storageScale[class])

		storageTotal[class] = total
	}

	endpointTotal = endpointTotal.Mul(fp.endpointScale)

	// Each quantity must be non-negative
	// and fit into an Int64
	if cpuTotal.IsNegative() ||
		memoryTotal.IsNegative() ||
		storageTotal.IsAnyNegative() ||
		endpointTotal.IsNegative() ||
		ipTotal.IsNegative() {
		return 0, ErrBidQuantityInvalid
	}

	totalCost := cpuTotal
	totalCost = totalCost.Add(memoryTotal)
	for _, total := range storageTotal {
		totalCost = totalCost.Add(total)
	}
	totalCost = totalCost.Add(endpointTotal)
	totalCost = totalCost.Add(ipTotal)

	if totalCost.IsNegative() {
		return 0, ErrBidQuantityInvalid
	}

	if totalCost.IsZero() {
		// Return an error indicating we can't bid with a cost of zero
		return 0, ErrBidZero
	}

	costDec, err := strconv.ParseUint(totalCost.String(), 10, 64)

	if err != nil {
		return 0, err
	}

	// if !costDec.LTE(sdk.MaxSortableDec) {
	// 	return 0, ErrBidQuantityInvalid
	// }

	return costDec, nil
}

type randomRangePricing int

func MakeRandomRangePricing() (BidPricingStrategy, error) {
	return randomRangePricing(0), nil
}

func (randomRangePricing) CalculatePrice(_ context.Context, req Request) (uint64, error) {
	min, max := calculatePriceRange(req.DeploymentSpec)
	if min.IsEqual(max) {
		return max, nil
	}

	const scale = 10000

	delta := max.Amount.Sub(min.Amount).Mul(sdk.NewDec(scale))

	minbid := delta.TruncateInt64()
	if minbid < 1 {
		minbid = 1
	}
	val, err := rand.Int(rand.Reader, big.NewInt(minbid))
	if err != nil {
		return sdk.DecCoin{}, err
	}

	scaledValue := sdk.NewDecFromBigInt(val).QuoInt64(scale).QuoInt64(100)
	amount := min.Amount.Add(scaledValue)
	return sdk.NewDecCoinFromDec(min.Denom, amount), nil
}

func calculatePriceRange(gspec *entities.DeploymentSpec) (sdk.DecCoin, sdk.DecCoin) {
	// memory-based pricing:
	//   min: requested memory * configured min price per Gi
	//   max: requested memory * configured max price per Gi

	// assumption: group.Count > 0
	// assumption: all same denom (returned by gspec.Price())
	// assumption: gspec.Price() > 0

	mem := sdk.NewInt(0)

	for _, group := range gspec.Resources {
		mem = mem.Add(
			sdk.NewIntFromUint64(group.Resources.Memory.Units).
				MulRaw(int64(group.Count)))
	}

	rmax := gspec.Price()

	const minGroupMemPrice = int64(50)
	const maxGroupMemPrice = int64(1048576)

	cmin := sdk.NewDecFromInt(mem.MulRaw(
		minGroupMemPrice).
		Quo(sdk.NewInt(unit.Gi)))

	cmax := sdk.NewDecFromInt(mem.MulRaw(
		maxGroupMemPrice).
		Quo(sdk.NewInt(unit.Gi)))

	if cmax.GT(rmax.Amount) {
		cmax = rmax.Amount
	}

	if cmin.IsZero() {
		cmin = sdk.NewDec(1)
	}

	if cmax.IsZero() {
		cmax = sdk.NewDec(1)
	}

	return sdk.NewDecCoinFromDec(rmax.Denom, cmin), sdk.NewDecCoinFromDec(rmax.Denom, cmax)
}

var errPathEmpty = errors.New("script path cannot be the empty string")
var errProcessLimitZero = errors.New("process limit must be greater than zero")
var errProcessRuntimeLimitZero = errors.New("process runtime limit must be greater than zero")

type storageElement struct {
	Class string `json:"class"`
	Size  uint64 `json:"size"`
}

type gpuVendorAttributes struct {
	Model string  `json:"model"`
	RAM   *string `json:"ram,omitempty"`
}

type gpuAttributes struct {
	Vendor map[string]gpuVendorAttributes `json:"vendor,omitempty"`
}

type gpuElement struct {
	Units      uint64        `json:"units"`
	Attributes gpuAttributes `json:"attributes"`
}

type dataForScriptElement struct {
	Memory           uint64           `json:"memory"`
	CPU              uint64           `json:"cpu"`
	GPU              gpuElement       `json:"gpu"`
	Storage          []storageElement `json:"storage"`
	Count            uint32           `json:"count"`
	EndpointQuantity int              `json:"endpoint_quantity"`
	IPLeaseQuantity  uint             `json:"ip_lease_quantity"`
}

type dataForScript struct {
	Resources      []dataForScriptElement `json:"resources"`
	Price          uint64                 `json:"price"`
	PricePrecision *int                   `json:"price_precision,omitempty"`
}
