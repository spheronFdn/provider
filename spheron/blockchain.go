package spheron

import (
	"context"
	"fmt"
	"log"
	"os"

	requestLogger "github.com/akash-network/provider/spheron/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func (client *Client) SubscribeEvents() error {
	// Define the contract address and ABI
	contractAddress := common.HexToAddress("0xfffaf1762a1fa569f639abe1c05f38f4745c4976")
	//contractABI := []byte(requestLogger.RequestLoggerMetaData.ABI)

	// Instantiate the contract
	contract, err := requestLogger.NewRequestLogger(contractAddress, client.EthClient)
	if err != nil {
		return err
	}

	// Create a channel to receive events
	eventChannel := make(chan *requestLogger.RequestLoggerRequestStored)
	defer close(eventChannel)

	// Subscribe to the event
	subscription, err := contract.WatchRequestStored(nil, eventChannel)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	fmt.Println("Listening for requests")

	// Launch a go routine and start listing to events
	go func() {
		for {
			select {
			case event := <-eventChannel:
				client.processEvents(event)
			case err := <-subscription.Err():
				log.Fatal("unable to subscribe to events: ", err)
			}
		}
	}()
	return nil
}

func (client *Client) processEvents(event *requestLogger.RequestLoggerRequestStored) {
	fmt.Printf("Received event: %v\n", event)
}

func (client *Client) PublishEvent() string {

	b, err := os.ReadFile("spheron/keys/wallet1.json") // where wallets
	if err != nil {
		log.Fatal(err)
	}

	const password = "testPassword"
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Create a new transactor with your private key
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// Define the contract address
	contractAddress := common.HexToAddress("0xfffaf1762a1fa569f639abe1c05f38f4745c4976")

	// Bind the contract with the client and auth
	instance, err := requestLogger.NewRequestLogger(contractAddress, client.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	// Trigger the event by calling the contract function
	tx, err := instance.StoreRequest(auth, "test request") // Pass the value for newValue
	if err != nil {
		log.Fatal(err)
	}

	// Print the transaction hash
	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (client *Client) CheckBalance() {

}
