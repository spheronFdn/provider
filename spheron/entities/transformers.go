package entities

import (
	"math/big"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	types "github.com/akash-network/akash-api/go/node/types/v1beta3"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/akash-network/akash-api/go/node/types/v1beta3"
)

func TransformGroupToOrder(gs *dtypes.GroupSpec) *Order {
	// Map the GroupSpec to a DeploymentSpec
	ds := DeploymentSpec{
		PlacementsRequirement: mapPlacementRequirements(gs.Requirements),
		Resources:             mapResourceUnits(gs.Resources),
	}

	// Return a Deployment with a placeholder DeploymentID and default state
	return &Order{
		ID:    0, // Adjust as necessary
		State: OrderOpen,
		Specs: ds,
	}
}

// Helper function to map PlacementRequirements from repo 1 to repo 2
func mapPlacementRequirements(pr v1beta3.PlacementRequirements) PlacementRequirements {
	return PlacementRequirements{
		ProviderWallets: mapProviderWallets(pr.SignedBy),
		Attributes:      mapAttributes(pr.Attributes),
	}
}

// Helper function to map SignedBy from repo 1 to repo 2
func mapProviderWallets(sb v1beta3.SignedBy) []string {
	return sb.AnyOf
}

// Helper function to map Attributes from repo 1 to repo 2
func mapAttributes(attrs v1beta3.Attributes) Attributes {
	mappedAttrs := make(Attributes, len(attrs))
	for i, attr := range attrs {
		mappedAttrs[i] = Attribute{
			Key:   attr.Key,
			Value: attr.Value,
		}
	}
	return mappedAttrs
}

// Helper function to map ResourceUnits from repo 1 to repo 2
func mapResourceUnits(rus dtypes.ResourceUnits) ServiceResources {
	mappedRus := make(ServiceResources, len(rus))
	for i, ru := range rus {
		mappedRus[i] = ServiceResource{
			Resources:    mapResources(ru.Resources),
			ReplicaCount: ru.Count,
		}
	}
	return mappedRus
}

// Helper function to map Resources from repo 1 to repo 2
func mapResources(res v1beta3.Resources) Resources {
	return Resources{
		ID:        res.ID,
		CPU:       mapCPU(res.CPU),
		Memory:    mapMemory(res.Memory),
		Storage:   mapVolumes(res.Storage),
		GPU:       mapGPU(res.GPU),
		Endpoints: mapEndpoints(res.Endpoints),
	}
}

// Helper function to map CPU from repo 1 to repo 2
func mapCPU(cpu *v1beta3.CPU) *CPU {
	if cpu == nil {
		return nil
	}
	return &CPU{
		Units:      cpu.Units.Value(),
		Attributes: mapAttributes(cpu.Attributes),
	}
}

// Helper function to map Memory from repo 1 to repo 2
func mapMemory(mem *v1beta3.Memory) *Memory {
	if mem == nil {
		return nil
	}
	return &Memory{
		Units:      mem.Quantity.Value(),
		Attributes: mapAttributes(mem.Attributes),
	}
}

// Helper function to map Volumes from repo 1 to repo 2
func mapVolumes(vols v1beta3.Volumes) Volumes {
	mappedVols := make(Volumes, len(vols))
	for i, vol := range vols {
		mappedVols[i] = Volume{
			Name:       vol.Name,
			Units:      vol.Quantity.Value(),
			Attributes: mapAttributes(vol.Attributes),
		}
	}
	return mappedVols
}

// Helper function to map GPU from repo 1 to repo 2
func mapGPU(gpu *v1beta3.GPU) *GPU {
	if gpu == nil {
		return nil
	}
	return &GPU{
		Units:      gpu.Units.Value(),
		Attributes: mapAttributes(gpu.Attributes),
	}
}

// Helper function to map Endpoints from repo 1 to repo 2
func mapEndpoints(eps v1beta3.Endpoints) Endpoints {
	mappedEps := make(Endpoints, len(eps))
	for i, ep := range eps {
		mappedEps[i] = Endpoint{
			Kind:           EndpointKind(ep.Kind),
			SequenceNumber: ep.SequenceNumber,
		}
	}
	return mappedEps
}

func TransformToResourceValue(value uint64) types.ResourceValue {
	bigIntValue := new(big.Int).SetUint64(value)
	resourceInt := cosmostypes.NewIntFromBigInt(bigIntValue)
	return types.ResourceValue{
		Val: resourceInt,
	}
}
