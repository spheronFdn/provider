package spheron

// TODO(spheron): add logging
import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"time"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	"github.com/akash-network/akash-api/go/node/market/v1beta4"
	"google.golang.org/grpc"

	"github.com/akash-network/provider/tools/fromctx"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

type AuthJson struct {
	PubKey          string `json:"pub_key"`
	Timestamp       uint64 `json:"timestamp"`
	SignedTimestamp string `json:"signed_timestamp"`
}

// Client defines the structure for our client to interact with the API.
type Client struct {
	BaseURL   string
	Context   Context
	EthClient *ethclient.Client
	Logger    log.Logger
	ChainEventCh   chan interface{}
}

type ClientConfig struct {
	HomeDir string
	Key     *keystore.Key
}

func NewClient(config ClientConfig) *Client {
	logger := fromctx.LogcFromCtx(context.Background())

	client, err := ethclient.DialContext(context.Background(), "wss://spheron-devnet.rpc.caldera.xyz/ws") // Use WebSocket RPC endpoint
	if err != nil {
		logger.Error("unable to connect to spheron-devnet")
	}

	context := &Context{
		Key:     config.Key,
		HomeDir: config.HomeDir,
	}
	return &Client{
		BaseURL:   "http://localhost:8088",
		Context:   *context,
		EthClient: client,
		Logger:    logger,
		ChainEventCh: make(chan interface{}),
	}
}

func NewClientWithContext(cctx Context) *Client {
	logger := fromctx.LogcFromCtx(context.Background())

	client, err := ethclient.DialContext(context.Background(), "wss://spheron-devnet.rpc.caldera.xyz/ws") // Use WebSocket RPC endpoint
	if err != nil {
		logger.Error("unable to connect to spheron-devnet")
	}

	return &Client{
		BaseURL:   "http://localhost:8088",
		Context:   cctx,
		EthClient: client,
		Logger:    logger,
	}
}

func (client *Client) SendRequest(ctx context.Context, endpoint string) ([]byte, error) {
	url := client.BaseURL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func (client *Client) SendPostRequest(ctx context.Context, endpoint string, data interface{}) ([]byte, error) {
	url := client.BaseURL + endpoint
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func (client *Client) GetGroup(ctx context.Context, dseq uint64) (dtypes.Group, error) {
	endpoint := fmt.Sprintf("/groups/%d", dseq)

	responseData, err := client.SendRequest(ctx, endpoint)
	if err != nil {
		return dtypes.Group{}, fmt.Errorf("error sending request JSON: %v", err)

	}

	var group dtypes.Group
	if err := json.Unmarshal(responseData, &group); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}
	return group, nil
}

func (client *Client) CreateBid(ctx context.Context, bidMsg v1beta4.MsgCreateBid) (interface{}, error) {
	resp, err := client.SendPostRequest(ctx, "/bid", bidMsg)

	var respObj interface{}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return respObj, err
}

func (client *Client) CloseBid(ctx context.Context, bidMsg v1beta4.MsgCloseBid) (interface{}, error) {
	resp, err := client.SendPostRequest(ctx, "/bid/close", bidMsg)

	var respObj interface{}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return respObj, err
}

func (client *Client) GetBid(ctx context.Context, dseq uint64) (*v1beta4.QueryBidResponse, error) {
	endpoint := fmt.Sprintf("/bid/%d", dseq)

	responseData, err := client.SendRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error sending request JSON: %v", err)
	}

	var response v1beta4.QueryBidResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func (client *Client) GetDeployment(ctx context.Context, dseq uint64) (*dtypes.QueryDeploymentResponse, error) {
	endpoint := fmt.Sprintf("/deployment/%d", dseq)

	responseData, err := client.SendRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error sending request JSON: %v", err)
	}

	var response dtypes.QueryDeploymentResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func (client *Client) GetLeases(ctx context.Context, dseq uint64) (*v1beta4.QueryLeasesResponse, error) {
	endpoint := fmt.Sprintf("/leases?dseq=%d&gseq=%d&oseq=%d", dseq, 1, 1)

	responseData, err := client.SendRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error sending request JSON: %v", err)
	}

	var response v1beta4.QueryLeasesResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func (client *Client) GetOrders(ctx context.Context, provider string) (*v1beta4.QueryOrdersResponse, error) {
	endpoint := fmt.Sprintf("/orders?provider=%s", provider)

	responseData, err := client.SendRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error sending request JSON: %v", err)
	}

	var response v1beta4.QueryOrdersResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func (client *Client) Leases(ctx context.Context, in *v1beta4.QueryLeasesRequest, opts ...grpc.CallOption) (*v1beta4.QueryLeasesResponse, error) {
	// TODO(spheron): fetch this information from our chain
	return nil, nil
}

func SignMessage(ctx context.Context, msg string) (interface{}, error) {
	// TODO(spheron): add signature with wallet
	signedMessage := msg
	return signedMessage, nil
}

func CreateAuthorizationToken(ctx context.Context, cctx *Context) (string, error) {
	ts := time.Now().Unix()
	tsStr := fmt.Sprintf("%v", ts)
	publicKey := cctx.Key.Address.Hex() // TODO(spheron) -> extract this data properly
	signedTimestamp, err := SignMessage(ctx, tsStr)
	if err != nil {
		return "", err
	}
	body := AuthJson{
		Timestamp:       uint64(ts),
		PubKey:          publicKey,
		SignedTimestamp: signedTimestamp.(string),
	}
	// Convert authToken to a base64-encoded string
	authTokenBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("unable to marshal auth token: %v", err.Error())
	}
	res := base64.StdEncoding.EncodeToString(authTokenBytes)
	return res, nil
}
