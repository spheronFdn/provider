package spheron

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/akash-network/provider/spheron/gen/NodeProviderRegistry"
	"github.com/akash-network/provider/spheron/gen/OrderMatching"

	"github.com/akash-network/node/pubsub"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/gen/requestLogger"

	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/akash-network/akash-api/go/sdkutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type EventRequestBody struct {
	EventType string `json:"event_type"`
	Body      string `json:"body"`
}

func (client *Client) SubscribeEvents(ctx context.Context, bus pubsub.Bus) error {

	err := client.subToRequestLogger(ctx, bus)
	if err != nil {
		return err
	}
	err = client.subToOrderMatching(ctx, bus)
	if err != nil {
		return err
	}

	go client.handleChainEvents(ctx, bus)
	return nil
}

func (client *Client) subToRequestLogger(ctx context.Context, bus pubsub.Bus) error {
	contractAddress := common.HexToAddress(requestLoggerContract)

	contract, err := requestLogger.NewRequestLogger(contractAddress, client.EthClient)
	if err != nil {
		return err
	}
	txch := make(chan *requestLogger.RequestLoggerRequestStored)

	subscription, err := contract.WatchRequestStored(nil, txch)
	if err != nil {
		return err
	}

	client.Logger.Debug("Listening for requests")

	go func() {
		select {
		case <-ctx.Done():
			subscription.Unsubscribe()
		case ev := <-txch:
			client.ChainEventCh <- ev
		}
	}()

	return nil
}

func (client *Client) subToOrderMatching(ctx context.Context, bus pubsub.Bus) error {
	// Todo:(spheron) change contract addr
	contractAddress := common.HexToAddress(requestLoggerContract)

	contract, err := OrderMatching.NewOrderMatching(contractAddress, client.EthClient)
	if err != nil {
		return err
	}
	// Create a channel to receive events
	txch := make(chan *OrderMatching.OrderMatchingOrderCreated)

	// Subscribe to the event
	subscription, err := contract.WatchOrderCreated(nil, txch)
	if err != nil {
		return err
	}

	client.Logger.Debug("Listening for requests")

	go func() {
		select {
		case <-ctx.Done():
			subscription.Unsubscribe()
		case ev := <-txch:
			client.ChainEventCh <- ev
		}
	}()

	return nil
}

func (client *Client) handleChainEvents(ctx context.Context, bus pubsub.Bus) {
	// Todo:(spheron) Decouple the channels.
loop:
	for {
		select {
		case <-ctx.Done():
			// sub.Unsubscribe()
			// close(eventCh)
			break loop // this thing might break
		case ev := <-client.ChainEventCh:
			switch ev.(type) {
			case *OrderMatching.OrderMatchingOrderCreated:
				// handle order create
			case *OrderMatching.OrderMatchingOrderMatched:
				// handle order matching
			case *requestLogger.RequestLoggerRequestStored:
				go client.processRequestLoggerEvents(ev.(*requestLogger.RequestLoggerRequestStored), bus)
			}
		}
	}
}

func (client *Client) handleOrderCreated(event *OrderMatching.OrderMatchingOrderCreated, bus pubsub.Bus) {

}

func (client *Client) handleOrderMatched(event *OrderMatching.OrderMatchingOrderMatched, bus pubsub.Bus) {

}

func (client *Client) processRequestLoggerEvents(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
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
	case "DeploymentCreated":
		evt := &entities.Deployment{}
		if err := json.Unmarshal([]byte(rawEvent.Body), evt); err != nil {
			return
		}
		msg := mtypes.EventOrderCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "bid-created"},
			ID: mtypes.OrderID{
				Owner: evt.ID.Owner,
				DSeq:  evt.ID.DSeq,
				GSeq:  1,
				OSeq:  1,
			}}
		internalEvent = msg
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
	contractAddress := common.HexToAddress(requestLoggerContract)

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

func (client *Client) GenerateTx(msg interface{}, eventType string) (*EventRequestBody, error) {

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

func (client *Client) AddNodeProvider(ctx context.Context, region string, paymentTokens []string) (string, error) {

	chainId, err := client.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(client.Context.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	// Todo:(spheron) change this to registerNode Contract
	contractAddress := common.HexToAddress(providerRegistryContract)

	instance, err := NodeProviderRegistry.NewNodeProviderRegistry(contractAddress, client.EthClient)
	if err != nil {
		return "", err
	}
	tx, err := instance.AddNodeProvider(auth, region, client.Context.Key.Address, paymentTokens)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (client *Client) RemoveNodeProvider(ctx context.Context, id *big.Int) (string, error) {

	chainId, err := client.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(client.Context.Key.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	// Todo:(spheron) change this to registerNode Contract
	contractAddress := common.HexToAddress(providerRegistryContract)

	instance, err := NodeProviderRegistry.NewNodeProviderRegistry(contractAddress, client.EthClient)
	if err != nil {
		return "", err
	}
	tx, err := instance.RemoveNodeProvider(auth, id)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (client *Client) CheckBalance() {

}
