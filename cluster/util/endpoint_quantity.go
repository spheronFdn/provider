package util

import (
	"github.com/akash-network/provider/spheron/entities"
)

func GetEndpointQuantityOfResourceGroup(resources entities.ResourceGroup, kind entities.Endpoint_Kind) uint {
	endpoints := make(map[uint32]struct{})
	for _, resource := range resources.GetResourceUnits() {
		accumEndpointsOfResources(resource.Resources, kind, endpoints)

	}
	return uint(len(endpoints))
}

func accumEndpointsOfResources(r entities.Resources, kind entities.Endpoint_Kind, accum map[uint32]struct{}) {
	for _, endpoint := range r.Endpoints {
		if entities.Endpoint_Kind(endpoint.Kind) == kind { //TODO (spheron): check why this requires conversion
			accum[endpoint.SequenceNumber] = struct{}{}
		}
	}
}

func GetEndpointQuantityOfResourceUnits(r entities.Resources, kind entities.Endpoint_Kind) uint {
	endpoints := make(map[uint32]struct{})
	accumEndpointsOfResources(r, kind, endpoints)

	return uint(len(endpoints))
}
