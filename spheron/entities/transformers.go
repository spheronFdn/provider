package entities

import (
	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	"github.com/akash-network/akash-api/go/node/types/v1beta3"
)

func TransformGroupToDeployment(gs *dtypes.GroupSpec) *Deployment {
	// Map the GroupSpec to a DeploymentSpec
	ds := DeploymentSpec{
		Name:         gs.Name,
		Requirements: mapPlacementRequirements(gs.Requirements),
		Resources:    mapResourceUnits(gs.Resources),
	}

	// Return a Deployment with a placeholder DeploymentID and default state
	return &Deployment{
		ID:    DeploymentID{}, // Adjust as necessary
		State: DeploymentOpen,
		Spec:  ds,
	}
}

// Helper function to map PlacementRequirements from repo 1 to repo 2
func mapPlacementRequirements(pr v1beta3.PlacementRequirements) PlacementRequirements {
	return PlacementRequirements{
		SignedBy:   mapSignedBy(pr.SignedBy),
		Attributes: mapAttributes(pr.Attributes),
	}
}

// Helper function to map SignedBy from repo 1 to repo 2
func mapSignedBy(sb v1beta3.SignedBy) SignedBy {
	return SignedBy{
		AllOf: sb.AllOf,
		AnyOf: sb.AnyOf,
	}
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
func mapResourceUnits(rus dtypes.ResourceUnits) ResourceUnits {
	mappedRus := make(ResourceUnits, len(rus))
	for i, ru := range rus {
		mappedRus[i] = ResourceUnit{
			Resources: mapResources(ru.Resources),
			Count:     ru.Count,
			Price:     ru.Price.Amount.BigInt().Uint64(),
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
		mappedVols[i] = Storage{
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
			Kind:           int32(ep.Kind),
			SequenceNumber: ep.SequenceNumber,
		}
	}
	return mappedEps
}
