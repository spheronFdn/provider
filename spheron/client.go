package spheron

// TODO(spheron): add logging
import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"io"
	"net/http"
	"time"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	"github.com/akash-network/akash-api/go/node/market/v1beta4"
	"google.golang.org/grpc"

	"github.com/akash-network/provider/tools/fromctx"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

type AuthJson struct {
	PubKey          string `json:"pub_key"`
	Timestamp       int64  `json:"timestamp"`
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

func SignMessage(key *keystore.Key, msg string) (interface{}, error) {
	// Convert the message to a hash to be signed
	hash := crypto.Keccak256Hash([]byte(msg)).Bytes()

	// Sign the hash using the private key
	signature, err := crypto.Sign(hash, key.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// Convert the signature to a hex string
	signatureHex := fmt.Sprintf("0x%x", signature)
	return signatureHex, nil
}

func validateSignature(pubKey *ecdsa.PublicKey, signedMsg string, originalMsg string) (bool, error) {
	publicKeyBytes := crypto.FromECDSAPub(pubKey)
	// Hash the original message
	msgHash := crypto.Keccak256Hash([]byte(originalMsg))

	// Convert the signature hex string to a byte slice
	signatureBytes, err := hex.DecodeString(signedMsg[2:]) // assuming the signature is prefixed with "0x"
	if err != nil {
		return false, fmt.Errorf("invalid signature hex: %v", err)
	}

	sigPublicKey, err := crypto.Ecrecover(msgHash.Bytes(), signatureBytes)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)

	return matches, nil
}

func validateTimestamp(originalMsg string) (bool, error) {
	timestamp, err := strconv.ParseInt(originalMsg, 10, 64)
	if err != nil {
		return false, fmt.Errorf("timestamp parsing failed: %v", err)
	}
	currentTime := time.Now().Unix()
	if currentTime-timestamp > 20 {
		return false, fmt.Errorf("timestamp is older than 20 seconds")
	}
	return true, nil
}

func ValidateAuthToken(pubKey *ecdsa.PublicKey, signedMsg string, originalMsg string) (bool, error) {
	validSig, err := validateSignature(pubKey, signedMsg, originalMsg)
	if err != nil {
		return false, err
	}

	validTimestamp, err := validateTimestamp(originalMsg)
	if err != nil {
		return false, err
	}

	return validSig && validTimestamp, nil
}

func EncodePublicKey(pubKey *ecdsa.PublicKey) string {
	if pubKey == nil {
		return ""
	}
	// Convert the public key to a byte slice
	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	// Encode the byte slice to a hex string
	pubKeyHex := hex.EncodeToString(pubKeyBytes)
	return pubKeyHex
}

func DecodePublicKey(pubKeyHex string) (*ecdsa.PublicKey, error) {
	// Decode the hex string to a byte slice
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key hex: %v", err)
	}

	// Unmarshal the byte slice to an ECDSA public key
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %v", err)
	}

	return pubKey, nil
}

func CreateAuthorizationToken(ctx context.Context, cctx *Context) (string, error) {
	ts := time.Now().Unix()
	signedTimestamp, err := SignMessage(cctx.Key, strconv.FormatInt(ts, 10))
	publicKeyHex := EncodePublicKey(&cctx.Key.PrivateKey.PublicKey)

	if err != nil {
		return "", err
	}

	body := AuthJson{
		Timestamp:       ts,
		PubKey:          publicKeyHex,
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
