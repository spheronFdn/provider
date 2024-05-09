package spheron

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/akash-network/node/pubsub"
	requestLogger "github.com/akash-network/provider/spheron/gen"

	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

func (client *Client) SubscribeEvents(ctx context.Context, bus pubsub.Bus) error {
	// Define the contract address and ABI
	contractAddress := common.HexToAddress("0xfffaf1762a1fa569f639abe1c05f38f4745c4976")
	//contractABI := []byte(requestLogger.RequestLoggerMetaData.ABI)

	// Instantiate the contract
	contract, err := requestLogger.NewRequestLogger(contractAddress, client.EthClient)
	if err != nil {
		return err
	}
	// Create a channel to receive events
	txch := make(chan *requestLogger.RequestLoggerRequestStored)

	// Subscribe to the event
	subscription, err := contract.WatchRequestStored(nil, txch)
	if err != nil {
		return err
	}

	client.Logger.Debug("Listening for requests")

	go client.publishEvents(ctx, subscription, txch, bus)

	return nil
}

func (client *Client) publishEvents(ctx context.Context, sub event.Subscription, txchan chan *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) error {
	var err error
loop:
	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			close(txchan)
			break loop
		case ed := <-txchan:
			client.processEvents(ed, bus)
		}
	}

	return err
}

type EventRequestBody struct {
	EventType string `json:"event_type"`
	Body      string `json:"body"`
}

func (client *Client) processEvents(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
	fmt.Printf("Received event: %v\n", event)
	// TODO(spheron) -> untill we have all contracts available take event.Request as if it's a json representation of akash event
	var internalEvent interface{}
	jsonByte := []byte(event.Request)
	rawEvent := &EventRequestBody{}
	if err := json.Unmarshal(jsonByte, rawEvent); err != nil {
		client.Logger.Error("unable to unmarshal event", err)
		return
	}

	switch rawEvent.EventType {
	case "EventOrderCreated":
		e := &mtypes.EventOrderCreated{}
		if err := json.Unmarshal([]byte(rawEvent.Body), e); err != nil {
			return
		}
		internalEvent = *e
	}
	if err := bus.Publish(internalEvent); err != nil {
		bus.Close()
		return
	}
}

func (client *Client) SendTx(body *EventRequestBody) (string, error) {

	chainId, err := client.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	// Create a new transactor with your private key
	auth, err := bind.NewKeyedTransactorWithChainID(client.Context.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	// Define the contract address
	contractAddress := common.HexToAddress("0xfffaf1762a1fa569f639abe1c05f38f4745c4976")

	// Bind the contract with the client and auth
	instance, err := requestLogger.NewRequestLogger(contractAddress, client.EthClient)
	if err != nil {
		return "", err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	tx, err := instance.StoreRequest(auth, string(jsonBody)) // Pass the value for newValue
	// Trigger the event by calling the contract function
	if err != nil {
		return "", err
	}

	// Print the transaction hash
	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func (clinet *Client) GenerateTx(msg interface{}, eventType string) (*EventRequestBody, error) {

	msgStr, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Error while marshaling tx msg")
		return nil, err

	}
	tx := &EventRequestBody{
		EventType: eventType,
		Body:      string(msgStr),
	}

	return tx, nil
}

func (client *Client) CheckBalance() {

}
