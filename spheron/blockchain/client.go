package blockchain

import (
	"context"
	"math/big"

	"github.com/akash-network/provider/spheron/blockchain/gen/NodeProviderRegistry"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/tools/fromctx"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

type BlockChainClient struct {
	EthClient    *ethclient.Client
	Logger       log.Logger
	ChainEventCh chan interface{}
	Key          *keystore.Key
}

type EventRequestBody struct {
	EventType string `json:"event_type"`
	Body      string `json:"body"`
}

func NewBlockChainClient(key *keystore.Key) *BlockChainClient {
	logger := fromctx.LogcFromCtx(context.Background())
	client, err := ethclient.DialContext(context.Background(), "wss://spheron-devnet.rpc.caldera.xyz/ws") // Use WebSocket RPC endpoint
	if err != nil {
		logger.Error("unable to connect to spheron-devnet")
	}

	return &BlockChainClient{
		EthClient:    client,
		Logger:       logger,
		ChainEventCh: make(chan interface{}),
		Key:          key,
	}
}

// Provider contract

func (b *BlockChainClient) AddNodeProvider(ctx context.Context, region string, paymentTokens []string) (string, error) {

	chainId, err := b.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(b.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	contractAddress := common.HexToAddress(providerRegistryContract)

	instance, err := NodeProviderRegistry.NewNodeProviderRegistry(contractAddress, b.EthClient)
	if err != nil {
		return "", err
	}
	tx, err := instance.AddNodeProvider(auth, region, b.Key.Address, paymentTokens)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (b *BlockChainClient) RemoveNodeProvider(ctx context.Context, id *big.Int) (string, error) {

	chainId, err := b.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(b.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	// Todo:(spheron) change this to registerNode Contract
	contractAddress := common.HexToAddress(providerRegistryContract)

	instance, err := NodeProviderRegistry.NewNodeProviderRegistry(contractAddress, b.EthClient)
	if err != nil {
		return "", err
	}
	tx, err := instance.RemoveNodeProvider(auth, id)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// Order contract
func (b *BlockChainClient) CreateOrder(ctx context.Context, order *entities.Order) (string, error) {

	chainId, err := b.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(orderMatchingContract)

	instance, err := OrderMatching.NewOrderMatching(contractAddress, b.EthClient)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(b.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}

	tx, err := instance.CreateOrder(auth, order.Region, order.Uptime, order.Reputation, order.Slashes,
		big.NewInt(int64(order.MaxPrice)), order.Token,
		stringSliceToAddressSlice(order.Specs.PlacementsRequirement.ProviderWallets), 1, "test",
		getMatchingResourceAttribute(order.Specs.Resources[0].Resources))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (b *BlockChainClient) GetOrderById(ctx context.Context, id uint64) (*entities.Order, error) {
	// TODO(spheron): interact with blockchain

	return &entities.Order{
		ID:         id,
		Region:     "us-east",
		Uptime:     0,
		Reputation: 0,
		Slashes:    0,
		MaxPrice:   10,
		Token:      "USDC",
		Creator:    "owner",
		State:      entities.OrderActive,
		Specs:      entities.DeploymentSpec{},
	}, nil
}

func (b *BlockChainClient) GetOrdersByProvider(ctx context.Context, provider string) ([]*entities.Order, error) {
	// TODO(spheron): interact with blockchain
	return []*entities.Order{
		{
			ID:         1,
			Region:     "us-east",
			Uptime:     0,
			Reputation: 0,
			Slashes:    0,
			MaxPrice:   10,
			Token:      "USDC",
			Creator:    "owner",
			State:      entities.OrderActive,
			Specs:      entities.DeploymentSpec{},
		},
	}, nil
}

func (b *BlockChainClient) CloseOrder(ctx context.Context, id uint64) (string, error) {
	// TODO(spheron): interact with blockchain
	return "", nil
}

func (b *BlockChainClient) CreateBid(ctx context.Context, bid *entities.Bid) (string, error) {
	// TODO(spheron): interact with blockchain
	return "", nil
}

func (b *BlockChainClient) GetBid(ctx context.Context, id uint64) (*entities.Bid, error) {
	// TODO(spheron): interact with blockchain

	return &entities.Bid{
		OrderID:  id,
		BidPrice: 1,
		Bidder:   "owner",
	}, nil
}

// Token registry
func (b *BlockChainClient) GetRegistedTokens(ctx context.Context) ([]string, error) {
	// TODO(spheron): interact with blockchain
	tokens := []string{"USDT", "USDC"}
	return tokens, nil
}

// Common operations
func (b *BlockChainClient) CheckBalance() {
}
