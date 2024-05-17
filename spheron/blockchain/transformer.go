package blockchain

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/events"
	"github.com/ethereum/go-ethereum/common"
)

func MapOrderCreated(event *OrderMatching.OrderMatchingOrderCreated) *events.OrderCreated {
	ev := events.OrderCreated{
		ID:      event.OrderId,
		Creator: "",
	}
	return &ev
}

func MapOrderMatched(event *OrderMatching.OrderMatchingOrderMatched) *events.OrderMatched {
	ev := events.OrderMatched{
		ID:       event.OrderId,
		Provider: event.ProviderAddress.Hex(),
		Creator:  "",
	}
	return &ev
}

func MapOrderClosed(event *OrderMatching.OrderMatchingOrderClosed) *events.OrderClosed {
	ev := events.OrderClosed{
		ID:       event.OrderId,
		Provider: "",
		Creator:  "",
	}
	return &ev
}

func MapChainOrderToOrder(initialOrder *struct {
	Id            uint64
	Region        string
	Uptime        uint64
	Reputation    uint64
	Slashes       uint64
	MaxPrice      *big.Int
	Token         string
	Creator       common.Address
	State         uint8
	Specs         string
	Version       string
	CreationBlock *big.Int
}) (*entities.Order, error) {
	// Unmarshal the Resources JSON string
	var deploymentSpec entities.DeploymentSpec
	err := json.Unmarshal([]byte(initialOrder.Specs), &deploymentSpec)
	if err != nil {
		return &entities.Order{}, fmt.Errorf("failed to unmarshal resources: %w", err)
	}

	// Map the order state
	var state entities.OrderState
	switch initialOrder.State {
	case 1:
		state = entities.OrderOpen
	case 2:
		state = entities.OrderActive
	case 3:
		state = entities.OrderClosed
	default:
		state = entities.OrderInvalid
	}

	// Map the initial order to the new order
	order := entities.Order{
		ID:         initialOrder.Id,
		Region:     initialOrder.Region,
		Uptime:     initialOrder.Uptime,
		Reputation: initialOrder.Reputation,
		Slashes:    initialOrder.Slashes,
		MaxPrice:   initialOrder.MaxPrice.Uint64(),
		Token:      initialOrder.Token,
		Creator:    initialOrder.Creator.Hex(),
		State:      state,
		Specs:      deploymentSpec,
	}

	return &order, nil
}

func MapChainLeaseToLease(lease *struct {
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	StartBlock      *big.Int
	StartTime       *big.Int
	EndBlock        *big.Int
	EndTime         *big.Int
	State           uint8
}, orderId uint64) (*entities.Lease, error) {
	// Map the order state
	var state entities.OrderState
	switch lease.State {
	case 1:
		state = entities.OrderOpen
	case 2:
		state = entities.OrderActive
	case 3:
		state = entities.OrderClosed
	default:
		state = entities.OrderInvalid
	}

	// Map the initial order to the new order
	l := entities.Lease{
		OrderID:         orderId,
		ProviderAddress: lease.ProviderAddress.Hex(),
		AcceptedPrice:   lease.AcceptedPrice.Uint64(),
		State:           state,
	}
	return &l, nil
}

func MapChainProviderToProvider(address string, region string, tokens []string, isActive bool, domain string) *entities.Provider {
	return &entities.Provider{
		WalletAddress: address,
		Region:        region,
		IsActive:      isActive,
		Tokens:        tokens,
		Domain:        domain,
	}
}
