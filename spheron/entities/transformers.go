package entities

import (
	"math/big"

	"github.com/akash-network/akash-api/go/manifest/v2beta2"
	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	types "github.com/akash-network/akash-api/go/node/types/v1beta3"

	"github.com/akash-network/akash-api/go/node/types/v1beta3"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
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
		Attributes: MapAttributes(pr.Attributes),
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
func MapAttributes(attrs v1beta3.Attributes) Attributes {
	mappedAttrs := make(Attributes, len(attrs))
	for i, attr := range attrs {
		mappedAttrs[i] = Attribute{
			Key:   attr.Key,
			Value: attr.Value,
		}
	}
	return mappedAttrs
}

// Helper function to map Attributes from repo 1 to repo 2
func MapAttributesToV1beta3(attrs Attributes) v1beta3.Attributes {
	mappedAttrs := make(v1beta3.Attributes, len(attrs))
	for i, attr := range attrs {
		mappedAttrs[i] = v1beta3.Attribute{
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
		Attributes: MapAttributes(cpu.Attributes),
	}
}

// Helper function to map Memory from repo 1 to repo 2
func mapMemory(mem *v1beta3.Memory) *Memory {
	if mem == nil {
		return nil
	}
	return &Memory{
		Units:      mem.Quantity.Value(),
		Attributes: MapAttributes(mem.Attributes),
	}
}

// Helper function to map Volumes from repo 1 to repo 2
func mapVolumes(vols v1beta3.Volumes) Volumes {
	mappedVols := make(Volumes, len(vols))
	for i, vol := range vols {
		mappedVols[i] = MapStorage(vol)
	}
	return mappedVols
}

func MapStorage(storage v1beta3.Storage) Storage {
	return Storage{
		Name:       storage.Name,
		Units:      storage.Quantity.Value(),
		Attributes: MapAttributes(storage.Attributes),
	}
}

// Helper function to map GPU from repo 1 to repo 2
func mapGPU(gpu *v1beta3.GPU) *GPU {
	if gpu == nil {
		return nil
	}
	return &GPU{
		Units:      gpu.Units.Value(),
		Attributes: MapAttributes(gpu.Attributes),
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

//	type OrderID struct {
//		Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
//		DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
//		GSeq  uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
//		OSeq  uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
//	}
func TransformOrderIDtoDeploymentID(orderID mtypes.OrderID) DeploymentID {
	return DeploymentID{
		Owner: orderID.Owner,
		DSeq:  orderID.DSeq,
	}
}

func TransformDeploymentIDtoOrderID(deploymentID DeploymentID) mtypes.OrderID {
	return mtypes.OrderID{
		Owner: deploymentID.Owner,
		DSeq:  deploymentID.DSeq,
		GSeq:  1,
		OSeq:  1,
	}
}

// transformGroup converts a v2beta2.Group to a targetpackage.Group
func TransformGroup(src *v2beta2.Group) Group {
	var dest Group
	dest.Name = src.Name
	dest.Services = transformServices(src.Services)
	return dest
}

// transformServices converts v2beta2.Services to targetpackage.Services
func transformServices(src v2beta2.Services) Services {
	dest := make(Services, len(src))
	for i, svc := range src {
		dest[i] = transformService(svc)
	}
	return dest
}

// transformService converts a v2beta2.Service to a targetpackage.Service
func transformService(src v2beta2.Service) Service {
	return Service{
		Name:      src.Name,
		Image:     src.Image,
		Command:   src.Command,
		Args:      src.Args,
		Env:       src.Env,
		Resources: mapResources(src.Resources),
		Count:     src.Count,
		Expose:    transformServiceExposes(src.Expose),
		Params:    transformServiceParams(src.Params),
		// Credentials: transformServiceCredentials(src.Credentials), //TODO (spheron): add this
	}
}

// transformServiceExposes converts a slice of v2beta2.ServiceExpose to a slice of targetpackage.ServiceExpose
func transformServiceExposes(src v2beta2.ServiceExposes) ServiceExposes {
	dest := make(ServiceExposes, len(src))
	for i, expose := range src {
		dest[i] = ServiceExpose{
			Port:                   expose.Port,
			ExternalPort:           expose.ExternalPort,
			Proto:                  ServiceProtocol(expose.Proto), // Casting enum type, assuming similar types
			Service:                expose.Service,
			Global:                 expose.Global,
			Hosts:                  expose.Hosts,
			HTTPOptions:            transformHTTPOptions(expose.HTTPOptions),
			IP:                     expose.IP,
			EndpointSequenceNumber: expose.EndpointSequenceNumber,
		}
	}
	return dest
}

// transformHTTPOptions converts v2beta2.ServiceExposeHTTPOptions to targetpackage.ServiceExposeHTTPOptions
func transformHTTPOptions(src v2beta2.ServiceExposeHTTPOptions) ServiceExposeHTTPOptions {
	return ServiceExposeHTTPOptions{
		MaxBodySize: src.MaxBodySize,
		ReadTimeout: src.ReadTimeout,
		SendTimeout: src.SendTimeout,
		NextTries:   src.NextTries,
		NextTimeout: src.NextTimeout,
		NextCases:   src.NextCases,
	}
}

// transformServiceParams converts v2beta2.ServiceParams to targetpackage.ServiceParams
func transformServiceParams(src *v2beta2.ServiceParams) *ServiceParams {
	if src == nil {
		return nil
	}
	dest := &ServiceParams{
		Storage: make([]StorageParams, len(src.Storage)),
	}
	for i, storage := range src.Storage {
		dest.Storage[i] = StorageParams{
			Name:     storage.Name,
			Mount:    storage.Mount,
			ReadOnly: storage.ReadOnly,
		}
	}
	return dest
}

func TransformToResourceValue(value uint64) types.ResourceValue {
	bigIntValue := new(big.Int).SetUint64(value)
	resourceInt := cosmostypes.NewIntFromBigInt(bigIntValue)
	return types.ResourceValue{
		Val: resourceInt,
	}
}
