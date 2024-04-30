package spheron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	"github.com/akash-network/akash-api/go/node/market/v1beta4"
)

// HelperClient defines the structure for our client to interact with the API.
type HelperClient struct {
	BaseURL string
}

// NewHelperClient creates a new HelperClient with the specified base URL.
func NewHelperClient(baseURL string) *HelperClient {
	return &HelperClient{BaseURL: baseURL}
}

func (client *HelperClient) SendRequest(ctx context.Context, endpoint string) ([]byte, error) {
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

func (client *HelperClient) SendPostRequest(ctx context.Context, endpoint string, data interface{}) ([]byte, error) {
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

func (client *HelperClient) GetGroup(ctx context.Context, dseq uint64) (dtypes.Group, error) {
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

func (client *HelperClient) CreateBid(ctx context.Context, bidMsg v1beta4.MsgCreateBid) (interface{}, error) {
	resp, err := client.SendPostRequest(ctx, "/bid", bidMsg)

	var respObj interface{}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return respObj, err
}

func (client *HelperClient) CloseBid(ctx context.Context, bidMsg v1beta4.MsgCloseBid) (interface{}, error) {
	resp, err := client.SendPostRequest(ctx, "/bid/close", bidMsg)

	var respObj interface{}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return respObj, err
}

func (client *HelperClient) GetBid(ctx context.Context, dseq uint64) (*v1beta4.QueryBidResponse, error) {
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

func (client *HelperClient) GetDeployment(ctx context.Context, dseq uint64) (*dtypes.QueryDeploymentResponse, error) {
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

func (client *HelperClient) GetLeases(ctx context.Context, dseq uint64) (*v1beta4.QueryLeasesResponse, error) {
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

func (client *HelperClient) GetOrders(ctx context.Context, provider string) (*v1beta4.QueryOrdersResponse, error) {
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
