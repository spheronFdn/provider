package spheron

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akash-network/node/pubsub"
	requestLogger "github.com/akash-network/provider/spheron/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
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

func (client *Client) processEvents(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
	fmt.Printf("Received event: %v\n", event)
}

func (client *Client) PublishEvent() (string, error) {

	b, err := os.ReadFile("spheron/keys/wallet1.json") // where wallets
	if err != nil {
		log.Fatal(err)
	}

	const password = "testPassword"
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		return "", err
	}

	chainId, err := client.EthClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	// Create a new transactor with your private key
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
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

	// Trigger the event by calling the contract function
	tx, err := instance.StoreRequest(auth, "test request") // Pass the value for newValue
	if err != nil {
		return "", err
	}

	// Print the transaction hash
	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func (client *Client) CheckBalance() {

}
