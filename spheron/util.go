package spheron

import (
	"os"

	types3 "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	types "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/spf13/pflag"
)

type MarketOptions struct {
	Owner    string
	Provider string
}

type MarketOption func(*MarketOptions)

func WithOwner(val string) MarketOption {
	return func(opt *MarketOptions) {
		opt.Owner = val
	}
}

func WithProvider(val string) MarketOption {
	return func(opt *MarketOptions) {
		opt.Provider = val
	}
}

func BidIDFromFlags(flags *pflag.FlagSet, opts ...MarketOption) (types.BidID, error) {
	prev, err := OrderIDFromFlags(flags, opts...)
	if err != nil {
		return types.BidID{}, err
	}

	opt := &MarketOptions{}

	for _, o := range opts {
		o(opt)
	}

	if opt.Provider == "" {
		provider, err := flags.GetString("provider")
		if err != nil {
			return types.BidID{}, err
		}

		return MakeBidID(prev, provider), nil
	}

	return MakeBidID(prev, opt.Provider), nil
}

func OrderIDFromFlags(flags *pflag.FlagSet, opts ...MarketOption) (types.OrderID, error) {
	prev, err := GroupIDFromFlags(flags, opts...)
	if err != nil {
		return types.OrderID{}, err
	}
	val, err := flags.GetUint32("oseq")
	if err != nil {
		return types.OrderID{}, err
	}
	return types.MakeOrderID(prev, val), nil
}

// GroupIDFromFlags returns GroupID with given flags and error if occurred
func GroupIDFromFlags(flags *pflag.FlagSet, opts ...MarketOption) (types3.GroupID, error) {
	var id types3.GroupID
	prev, err := DeploymentIDFromFlags(flags, opts...)
	if err != nil {
		return id, err
	}

	gseq, err := flags.GetUint32("gseq")
	if err != nil {
		return id, err
	}
	return types3.MakeGroupID(prev, gseq), nil
}

// DeploymentIDFromFlags returns DeploymentID with given flags, owner and error if occurred
func DeploymentIDFromFlags(flags *pflag.FlagSet, opts ...MarketOption) (types3.DeploymentID, error) {
	var id types3.DeploymentID
	opt := &MarketOptions{}

	for _, o := range opts {
		o(opt)
	}

	var owner string
	if flag := flags.Lookup("owner"); flag != nil {
		owner = flag.Value.String()
	}

	// if --owner flag was explicitly provided, use that.
	var err error
	if owner != "" {
		id.Owner = opt.Owner
	}

	if id.DSeq, err = flags.GetUint64("dseq"); err != nil {
		return id, err
	}
	return id, nil
}

func MakeBidID(id types.OrderID, provider string) types.BidID {
	return types.BidID{
		Owner:    id.Owner,
		DSeq:     id.DSeq,
		GSeq:     id.GSeq,
		OSeq:     id.OSeq,
		Provider: provider,
	}
}

func ReadKey(walletPath string, password string) *keystore.Key {
	b, _ := os.ReadFile(walletPath)
	key, _ := keystore.DecryptKey(b, password)
	if key != nil {
		return key
	}
	return nil
}

func ArrayContainsString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
