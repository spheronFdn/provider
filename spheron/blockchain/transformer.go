package blockchain

import (
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/events"
)

func MapOrderCreated(event *OrderMatching.OrderMatchingOrderCreated) *events.OrderCreated {
	ev := events.OrderCreated{
		ID: event.OrderId,
	}
	return &ev
}

func MapOrderMatched(event *OrderMatching.OrderMatchingOrderMatched) *events.OrderMatched {
	ev := events.OrderMatched{
		ID:       event.OrderId,
		Provider: event.ProviderAddress.Hex(),
	}
	return &ev
}

func MapOrderClosed(event *OrderMatching.OrderMatchingOrderClosed) *events.OrderClosed {
	ev := events.OrderClosed{
		ID: event.OrderId,
	}
	return &ev
}

