package spheron

// TODO(spheron): add logging
import (
	"context"
	"fmt"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"

	"github.com/akash-network/akash-api/go/node/market/v1beta4"
	ptypes "github.com/akash-network/akash-api/go/node/provider/v1beta3"

	"google.golang.org/grpc"

	"github.com/akash-network/provider/spheron/blockchain"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/tools/fromctx"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/libs/log"
)

type AuthJson struct {
	PubKey          string `json:"pub_key"`
	Timestamp       int64  `json:"timestamp"`
	SignedTimestamp string `json:"signed_timestamp"`
}

// Client defines the structure for our client to interact with the API.
type Client struct {
	Context  Context
	BcClient *blockchain.BlockChainClient
	Logger   log.Logger
}

type ClientConfig struct {
	HomeDir string
	Key     *keystore.Key
}

func NewClient(config ClientConfig) *Client {
	logger := fromctx.LogcFromCtx(context.Background())
	b, err := blockchain.NewBlockChainClient(config.Key)
	if err != nil {
		logger.Error("unable to initiate blockchain client ", err)
		panic("spheron init failed")
	}

	context := &Context{
		Key:     config.Key,
		HomeDir: config.HomeDir,
	}

	return &Client{
		Context:  *context,
		Logger:   logger,
		BcClient: b,
	}
}

func NewClientWithContext(cctx Context) *Client {
	logger := fromctx.LogcFromCtx(context.Background())
	b, err := blockchain.NewBlockChainClient(cctx.Key)
	if err != nil {
		logger.Error("unable to initiate blockchain client ", err)
		panic("spheron init failed")
	}

	return &Client{
		Context:  cctx,
		Logger:   logger,
		BcClient: b,
	}
}

func (client *Client) GetGroup(ctx context.Context, dseq uint64) (dtypes.Group, error) {

	o, err := client.BcClient.GetOrderById(ctx, dseq)
	if err != nil {
		return dtypes.Group{}, fmt.Errorf("error fetching order from chain: %v", err)

	}
	// transform order into group
	return entities.MapOrderToGroup(o), nil
}

func (client *Client) CreateBid(ctx context.Context, bid *entities.Bid) (string, error) {
	tx, err := client.BcClient.CreateBid(ctx, bid)
	if err != nil {
		return "", err
	}
	return tx, err
}

func (client *Client) GetBid(ctx context.Context, dseq uint64) (*v1beta4.QueryBidResponse, error) {

	_, err := client.BcClient.GetBid(ctx, dseq)
	if err != nil {
		return nil, fmt.Errorf("error fetching bid from contract: %v", err)
	}

	var response v1beta4.QueryBidResponse
	return &response, nil
}

func (client *Client) GetOrdersByProvider(ctx context.Context, provider string) (*v1beta4.QueryOrdersResponse, error) {
	// TODO remap chain response to fit response type
	responseOrders := v1beta4.Orders{}

	orders, err := client.BcClient.GetOrdersByProvider(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("error fetching Orders from chain: %v", err)
	}

	for _, o := range orders {
		o := entities.MapOrderToV1Order(o)

		responseOrders = append(responseOrders, o)
	}

	response := v1beta4.QueryOrdersResponse{
		Orders: responseOrders,
	}
	return &response, nil
}

func (client *Client) GetOrdersWithFilter(ctx context.Context, in *v1beta4.QueryLeasesRequest, opts ...grpc.CallOption) (*v1beta4.QueryLeasesResponse, error) {
	// TODO(spheron): fetch this information from our chain
	return nil, nil
}

func (client *Client) GetProviderByAddress(ctx context.Context, address string) (*ptypes.Provider, []string, error) {

	provider, err := client.BcClient.GetProviderByAddress(ctx, common.HexToAddress(address))
	if err != nil {
		return nil, nil, err
	}
	p := entities.MapProviderToV3Provider(provider)

	return p, provider.Tokens, nil
}
