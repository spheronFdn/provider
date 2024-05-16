package blockchain

import (
	"context"
	"math/big"

	"github.com/akash-network/provider/spheron/blockchain/gen/NodeProviderRegistry"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/blockchain/gen/TokenRegistry"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/tools/fromctx"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

type BlockChainClient struct {
	EthClient            *ethclient.Client
	Logger               log.Logger
	ChainEventCh         chan interface{}
	Key                  *keystore.Key
	Auth                 *bind.TransactOpts
	NodeProviderRegistry *NodeProviderRegistry.NodeProviderRegistry
	OrderMatching        *OrderMatching.OrderMatching
	TokenRegistry        *TokenRegistry.TokenRegistry
}

type EventRequestBody struct {
	EventType string `json:"event_type"`
	Body      string `json:"body"`
}

func NewBlockChainClient(key *keystore.Key) (*BlockChainClient, error) {
	logger := fromctx.LogcFromCtx(context.Background())
	client, err := ethclient.DialContext(context.Background(), "wss://spheron-devnet.rpc.caldera.xyz/ws") // Use WebSocket RPC endpoint
	if err != nil {
		return nil, err
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	npr, err := NodeProviderRegistry.NewNodeProviderRegistry(common.HexToAddress(providerRegistryContract), client)
	if err != nil {
		return nil, err
	}

	om, err := OrderMatching.NewOrderMatching(common.HexToAddress(orderMatchingContract), client)

	tkrg, err := TokenRegistry.NewTokenRegistry(common.HexToAddress(tokenRegistryContract), client)
	return &BlockChainClient{
		EthClient:            client,
		Logger:               logger,
		ChainEventCh:         make(chan interface{}),
		Key:                  key,
		Auth:                 auth,
		NodeProviderRegistry: npr,
		OrderMatching:        om,
		TokenRegistry:        tkrg,
	}, nil
}

// Provider contract
func (b *BlockChainClient) AddNodeProvider(ctx context.Context, region string, paymentTokens []string) (string, error) {
	tx, err := b.NodeProviderRegistry.AddNodeProvider(b.Auth, region, b.Key.Address, paymentTokens)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (b *BlockChainClient) RemoveNodeProvider(ctx context.Context, id *big.Int) (string, error) {
	tx, err := b.NodeProviderRegistry.RemoveNodeProvider(b.Auth, id)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}



func (b *BlockChainClient) GetProviderByAddress(ctx context.Context, address common.Address) (*entities.Provider, error) {
	opts := &bind.CallOpts{
		From: b.Key.Address, //TODO(spheron): check on this
	}
	_, region, paymentAccepted, isActive, err := b.NodeProviderRegistry.GetNodeProviderByAddress(opts, address )
	if err != nil {
		return nil, err
	}
	// TODO:(spheron) replace domain mock for provider
	return &entities.Provider{
		 WalletAddress: address.Hex(),
		 Region: region,
		 IsActive: isActive,
		 Tokens: paymentAccepted,
		 Domain: "https://localhost:8443",
	}, nil
}

// Order contract
func (b *BlockChainClient) CreateOrder(ctx context.Context, order *entities.Order) (string, error) {
	tx, err := b.OrderMatching.CreateOrder(b.Auth, order.Region, order.Uptime, order.Reputation, order.Slashes,
		big.NewInt(int64(order.MaxPrice)), order.Token, getOrderSpec(order.Specs), "test")
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (b *BlockChainClient) GetOrderById(ctx context.Context, id uint64) (*entities.Order, error) {
	opts := &bind.CallOpts{
		From: b.Key.Address, //TODO(spheron): check on this
	}

	o, err := b.OrderMatching.GetOrderById(opts, id)
	if err != nil {
		return nil, err
	}

	order, err := entities.MapOrderMatchingOrderToOrder(o)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (b *BlockChainClient) GetOrdersByProvider(ctx context.Context, provider string) ([]*entities.Order, error) {
	orders := []*entities.Order{}
	opts := &bind.CallOpts{
		From: b.Key.Address, //TODO(spheron): check on this
	}

	ids, err := b.OrderMatching.GetOrderByProvider(opts, common.HexToAddress(provider))
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		o, err := b.OrderMatching.GetOrderById(opts, id)
		if err != nil {
			return nil, err
		}

		order, err := entities.MapOrderMatchingOrderToOrder(o)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (b *BlockChainClient) CloseOrder(ctx context.Context, id uint64) (string, error) {
	tx, err := b.OrderMatching.CloseOrder(b.Auth, id)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (b *BlockChainClient) CreateBid(ctx context.Context, bid *entities.Bid) (string, error) {

	tx, err := b.OrderMatching.PlaceBid(b.Auth, bid.OrderID, big.NewInt(int64(bid.BidPrice)))
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
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
	// opts := &bind.CallOpts{
	// 	From: b.Key.Address, //TODO(spheron): check on this
	// }
	// ids, err := b.TokenRegistry())

	tokens := []string{"USDT", "USDC"}
	return tokens, nil
}

// Common operations
func (b *BlockChainClient) CheckBalance() {
}
