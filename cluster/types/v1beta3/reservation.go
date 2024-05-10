package v1beta3

import (
	"github.com/akash-network/provider/spheron/entities"
)

//go:generate mockery --name ReservationGroup --output ./mocks
type ReservationGroup interface {
	Resources() entities.ResourceGroup
	SetAllocatedResources(entities.ResourceUnits)
	GetAllocatedResources() entities.ResourceUnits
	SetClusterParams(interface{})
	ClusterParams() interface{}
}

// Reservation interface implements orders and resources
//
//go:generate mockery --name Reservation --output ./mocks
type Reservation interface {
	DeploymentID() entities.DeploymentID
	Allocated() bool
	ReservationGroup
}
