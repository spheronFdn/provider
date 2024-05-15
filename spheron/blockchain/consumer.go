package blockchain

import (
	"context"

	"github.com/akash-network/node/pubsub"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/blockchain/gen/requestLogger"
	"github.com/akash-network/provider/spheron/events"
	"github.com/ethereum/go-ethereum/common"
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
	// TODO(spheron): Add handling of order created

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
