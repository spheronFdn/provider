package cluster

import (
	ctypes "github.com/akash-network/provider/cluster/types/v1beta3"
	"github.com/akash-network/provider/cluster/util"
	"github.com/akash-network/provider/spheron/entities"
)

func newReservation(deploymentID entities.DeploymentID, resources entities.ResourceGroup) *reservation {
	return &reservation{
		deploymentID:     deploymentID,
		resources:        resources,
		endpointQuantity: util.GetEndpointQuantityOfResourceGroup(resources, entities.Endpoint_LEASED_IP)}
}

type reservation struct {
	deploymentID      entities.DeploymentID
	resources         entities.ResourceGroup
	adjustedResources entities.ResourceUnits
	clusterParams     interface{}
	endpointQuantity  uint
	allocated         bool
	ipsConfirmed      bool
}

var _ ctypes.Reservation = (*reservation)(nil) //TODO (spheron): why this??????

func (r *reservation) DeploymentID() entities.DeploymentID {
	return r.deploymentID
}

func (r *reservation) Resources() entities.ResourceGroup {
	return r.resources
}

func (r *reservation) SetAllocatedResources(val entities.ResourceUnits) {
	r.adjustedResources = val
}

func (r *reservation) GetAllocatedResources() entities.ResourceUnits {
	return r.adjustedResources
}

func (r *reservation) SetClusterParams(val interface{}) {
	r.clusterParams = val
}

func (r *reservation) ClusterParams() interface{} {
	return r.clusterParams
}

func (r *reservation) Allocated() bool {
	return r.allocated
}
