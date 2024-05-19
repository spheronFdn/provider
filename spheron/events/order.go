package events

import "github.com/akash-network/provider/spheron/entities"

type OrderCreated struct {
	ID      uint64
	Creator string
}

type OrderMatched struct {
	ID       uint64
	Provider string
	Creator  string
}

type OrderUpdateRequest struct {
	ID       uint64
	NewPrice uint64
	Specs    entities.DeploymentSpec
}

type OrderUpdateConfirm struct {
	ID uint64
}

type OrderUpdated struct {
	ID uint64
}

type OrderClosed struct {
	ID       uint64
	Provider string
	Creator  string
}
