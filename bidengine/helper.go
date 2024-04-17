package bidengine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
)

func sendRequest(endpoint string) ([]byte, error) {
	url := "http://localhost:8088" + endpoint
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

func getGroup(dseq uint64) (dtypes.Group, error) {
	endpoint := fmt.Sprintf("/groups/%d", dseq)

	responseData, err := sendRequest(endpoint)
	if err != nil {
		return dtypes.Group{}, fmt.Errorf("error sending request JSON: %v", err)

	}

	var group dtypes.Group
	if err := json.Unmarshal(responseData, &group); err != nil {
		return dtypes.Group{}, fmt.Errorf("error decoding JSON: %v", err)
	}
	return group, nil
}
