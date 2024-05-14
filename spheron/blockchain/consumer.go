package blockchain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akash-network/akash-api/go/sdkutil"
	"github.com/akash-network/node/pubsub"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/blockchain/gen/requestLogger"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/events"
	"github.com/ethereum/go-ethereum/common"

	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
)

func (b *BlockChainClient) SubscribeEvents(ctx context.Context, bus pubsub.Bus) error {

	err := b.subscribeToOrderMatching(ctx)
	if err != nil {
		return err
	}

	go b.handleChainEvents(ctx, bus)
	return nil
}

func (b *BlockChainClient) subscribeToOrderMatching(ctx context.Context) error {
	contractAddress := common.HexToAddress(orderMatchingContract)

	contract, err := OrderMatching.NewOrderMatching(contractAddress, b.EthClient)
	if err != nil {
		return err
	}

	// Create a channel to receive events
	txch1 := make(chan *OrderMatching.OrderMatchingOrderCreated)
	txch2 := make(chan *OrderMatching.OrderMatchingOrderMatched)

	// Subscribe to chain events
	subscription1, err := contract.WatchOrderCreated(nil, txch1)
	if err != nil {
		return err
	}

	subscription2, err := contract.WatchOrderMatched(nil, txch2)
	if err != nil {
		return err
	}

	go func() {
		select {
		case <-ctx.Done():
			subscription1.Unsubscribe()
			subscription2.Unsubscribe()
		case ev := <-txch1:
			b.ChainEventCh <- ev
		case ev := <-txch2:
			b.ChainEventCh <- ev
		}
	}()
	return nil
}

func (b *BlockChainClient) handleChainEvents(ctx context.Context, bus pubsub.Bus) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop // this thing might break
		case ev := <-b.ChainEventCh:
			switch ev.(type) {
			case *OrderMatching.OrderMatchingOrderCreated:
				go b.handleOrderCreated(ev.(*OrderMatching.OrderMatchingOrderCreated), bus)
			case *OrderMatching.OrderMatchingOrderMatched:
				go b.handleOrderMatched(ev.(*OrderMatching.OrderMatchingOrderMatched), bus)
			case *requestLogger.RequestLoggerRequestStored:
				go b.handleRequestLogger(ev.(*requestLogger.RequestLoggerRequestStored), bus) // proxy for all events we don't have currently
			}
		}
	}
}

func (b *BlockChainClient) handleOrderCreated(event *OrderMatching.OrderMatchingOrderCreated, bus pubsub.Bus) {

	ev := MapOrderCreated(event)
	msg := events.MapOrderCreated(ev)

	if err := bus.Publish(msg); err != nil {
		bus.Close()
		return
	}
}

func (b *BlockChainClient) handleOrderMatched(event *OrderMatching.OrderMatchingOrderMatched, bus pubsub.Bus) {
	ev := MapOrderMatched(event)
	msg := events.MapOrderMatched(ev)

	if err := bus.Publish(msg); err != nil {
		bus.Close()
		return
	}
}

// hack function for handling all events we don't have in contracts right now
func (b *BlockChainClient) handleRequestLogger(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
	fmt.Printf("Received event: %v\n", event)
	// TODO(spheron) -> untill we have all contracts available take event.Request as if it's a json representation of akash event
	var internalEvent interface{}
	jsonByte := []byte(event.Request)
	rawEvent := &EventRequestBody{}
	if err := json.Unmarshal(jsonByte, rawEvent); err != nil {
		b.Logger.Error("unable to unmarshal event", err)
		return
	}

	switch rawEvent.EventType {
	case "DeploymentCreated":
		evt := &entities.Order{}
		if err := json.Unmarshal([]byte(rawEvent.Body), evt); err != nil {
			return
		}
		msg := mtypes.EventOrderCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "bid-created"},
			ID: mtypes.OrderID{
				Owner: "evt.ID.Owner",
				DSeq:  1, //
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

func (b *BlockChainClient) handleOrderUpdateRequest(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
	ev := MapOrderOrderUpdateRequest(event)
	msg := events.MapOrderUpdateRequested(ev)

	if err := bus.Publish(msg); err != nil {
		bus.Close()
		return
	}
}

func (b *BlockChainClient) handleOrderClosed(event *requestLogger.RequestLoggerRequestStored, bus pubsub.Bus) {
	ev := MapOrderClosed(event)
	msg1, msg2 := events.MapOrderClosed(ev)

	if err := bus.Publish(msg1); err != nil {
		bus.Close()
		return
	}
	if err := bus.Publish(msg2); err != nil {
		bus.Close()
		return
	}
}
