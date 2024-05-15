package blockchain

import (
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/blockchain/gen/requestLogger"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/events"
)

func MapOrderCreated(event *OrderMatching.OrderMatchingOrderCreated) *events.OrderCreated {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	// fetch from sc order and populate additional details
	ev := events.OrderCreated{
		ID: event.OrderId,
	}
	return &ev
}

func MapOrderMatched(event *OrderMatching.OrderMatchingOrderMatched) *events.OrderMatched {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	ev := events.OrderMatched{
		ID:       1,
		Provider: "provider",
	}
	return &ev
}

func MapOrderOrderUpdateRequest(event *requestLogger.RequestLoggerRequestStored) *events.OrderUpdateRequest {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	ev := events.OrderUpdateRequest{
		ID:       1,
		NewPrice: 10,
		Specs:    entities.DeploymentSpec{},
	}
	return &ev
}

func MapOrderUpdateConfirm(event *requestLogger.RequestLoggerRequestStored) *events.OrderUpdated {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	ev := events.OrderUpdated{
		ID: 1,
	}
	return &ev
}

func MapOrderClosed(event *requestLogger.RequestLoggerRequestStored) *events.OrderClosed {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	ev := events.OrderClosed{
		ID: 1,
	}
	return &ev
}

func MapBidPlaced(event *requestLogger.RequestLoggerRequestStored) *events.BidPlaced {
	// Spheron(TODO): when event is properly written in the contract replace id properly
	ev := events.BidPlaced{
		ID:       1,
		BidPrice: 10,
		Bidder:   "someone",
	}
	return &ev
}
