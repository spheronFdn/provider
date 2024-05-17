package blockchain

import (
	"context"

	"github.com/akash-network/node/pubsub"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/events"
	"github.com/ethereum/go-ethereum/common"
)

func (b *BlockChainClient) SubscribeEvents(ctx context.Context, bus pubsub.Bus) error {

	err := b.subscribeToOrderMatching(ctx)
	if err != nil {
		return err
	}

	go b.run(ctx, bus)
	return nil
}

func (b *BlockChainClient) subscribeToOrderMatching(ctx context.Context) error {
	contractAddress := common.HexToAddress(orderMatchingContract)

	contract, err := OrderMatching.NewOrderMatching(contractAddress, b.EthClient)
	if err != nil {
		return err
	}

	// Create a channel to receive events
	ocreatedch := make(chan *OrderMatching.OrderMatchingOrderCreated)
	omatchedch := make(chan *OrderMatching.OrderMatchingOrderMatched)
	oclosedch := make(chan *OrderMatching.OrderMatchingOrderClosed)

	// TODO(spheron): Add handling of order created

	// Subscribe to chain events
	screated, err := contract.WatchOrderCreated(nil, ocreatedch)
	if err != nil {
		return err
	}

	smatched, err := contract.WatchOrderMatched(nil, omatchedch)
	if err != nil {
		return err
	}

	sclosed, err := contract.WatchOrderClosed(nil, oclosedch)
	if err != nil {
		return err
	}

	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				screated.Unsubscribe()
				smatched.Unsubscribe()
				sclosed.Unsubscribe()
				break loop
			case ev := <-ocreatedch:
				b.ChainEventCh <- ev
			case ev := <-omatchedch:
				b.ChainEventCh <- ev
			case ev := <-oclosedch:
				b.ChainEventCh <- ev
			}
		}
	}()
	return nil
}

func (b *BlockChainClient) run(ctx context.Context, bus pubsub.Bus) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop // this thing might break
		case ev := <-b.ChainEventCh:
			switch ev := ev.(type) {
			case *OrderMatching.OrderMatchingOrderCreated:
				go b.handleOrderCreated(ev, bus)
			case *OrderMatching.OrderMatchingOrderMatched:
				go b.handleOrderMatched(ev, bus)
			case *OrderMatching.OrderMatchingOrderClosed:
				go b.handleOrderClosed(ev, bus)
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

func (b *BlockChainClient) handleOrderClosed(event *OrderMatching.OrderMatchingOrderClosed, bus pubsub.Bus) {
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
