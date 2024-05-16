package entities

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	
	ptypes "github.com/akash-network/akash-api/go/node/provider/v1beta3"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	"github.com/akash-network/akash-api/go/node/market/v1beta4"
	types "github.com/akash-network/akash-api/go/node/types/v1beta3"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

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

func MapOrderMatchingOrderToOrder(initialOrder OrderMatching.) (Order, error) {
	// Unmarshal the Resources JSON string
	var resources ServiceResources
	err := json.Unmarshal([]byte(initialOrder.Specs.Resources), &resources)
	if err != nil {
		return Order{}, fmt.Errorf("failed to unmarshal resources: %w", err)
	}

	// Map the order state
	var state OrderState
	switch initialOrder.State {
	case 1:
		state = OrderOpen
	case 2:
		state = OrderActive
	case 3:
		state = OrderClosed
	default:
		state = OrderInvalid
	}

	// Create the DeploymentSpec
	deploymentSpec := DeploymentSpec{
		PlacementsRequirement: PlacementRequirements{
			ProviderWallets: mapCommonAddressesToStrings(initialOrder.Specs.PlacementRequirement.ProviderWallets),
			Attributes:      AttributesFromStringSlice(initialOrder.Specs.PlacementRequirement.Attributes),
		},
		Resources: resources,
	}

	// Map the initial order to the new order
	order := Order{
		ID:         initialOrder.Id,
		Region:     initialOrder.Region,
		Uptime:     initialOrder.Uptime,
		Reputation: initialOrder.Reputation,
		Slashes:    initialOrder.Slashes,
		MaxPrice:   initialOrder.MaxPrice.Uint64(),
		Token:      initialOrder.Token,
		Creator:    initialOrder.Creator.Hex(),
		State:      state,
		Specs:      deploymentSpec,
	}

	return order, nil
}

func mapCommonAddressesToStrings(addresses []common.Address) []string {
	strAddresses := make([]string, len(addresses))
	for i, addr := range addresses {
		strAddresses[i] = addr.Hex()
	}
	return strAddresses
}

func MapOrderToV1Order(order *Order) v1beta4.Order {
	// Map OrderState to Group_State
	var state v1beta4.Order_State
	switch order.State {
	case OrderOpen:
		state = v1beta4.OrderOpen
	case OrderActive:
		state = v1beta4.OrderActive
	case OrderClosed:
		state = v1beta4.OrderClosed
	default:
		state = v1beta4.OrderStateInvalid
	}

	// Map DeploymentSpec to GroupSpec
	groupSpec := dtypes.GroupSpec{
		Name: "default", // Adjust as needed
		Requirements: v1beta3.PlacementRequirements{
			SignedBy: v1beta3.SignedBy{
				AllOf: order.Specs.PlacementsRequirement.ProviderWallets, // Assuming these map to AllOf
				AnyOf: []string{},                                        // Populate as needed
			},
			Attributes: MapAttributes(order.Specs.PlacementsRequirement.Attributes),
		},
		Resources: mapServiceResourcesToResourceUnits(order.Specs.Resources),
	}

	// Create the GroupID
	orderID := v1beta4.OrderID{
		Owner: order.Creator,
		DSeq:  order.ID,
		GSeq:  1, // Adjust as needed
		OSeq:  1,
	}

	// Create the Group
	group := v1beta4.Order{
		OrderID:   orderID,
		State:     state,
		Spec:      groupSpec,
		CreatedAt: 0, // Adjust as needed
	}

	return group
}

func mapServiceResourcesToResourceUnits(resources ServiceResources) dtypes.ResourceUnits {
	resourceUnits := make(dtypes.ResourceUnits, len(resources))
	for i, r := range resources {
		resourceUnits[i] = dtypes.ResourceUnit{
			Resources: v1beta3.Resources{
				ID:        r.Resources.ID,
				CPU:       mapGroupCPUToOrderCPU(r.Resources.CPU),
				Memory:    mapGroupMemoryToOrderMemory(r.Resources.Memory),
				Storage:   mapGroupVolumesToOrderVolumes(r.Resources.Storage),
				GPU:       mapGroupGPUToOrderGPU(r.Resources.GPU),
				Endpoints: mapGroupEndpointsToOrderEndpoints(r.Resources.Endpoints),
			},
			Count: r.ReplicaCount,
			Price: cosmostypes.DecCoin{
				Denom:  "usd", // Adjust as needed
				Amount: cosmostypes.OneDec(),
			},
		}
	}
	return resourceUnits
}

func mapGroupCPUToOrderCPU(cpu *CPU) *v1beta3.CPU {
	if cpu == nil {
		return nil
	}
	return &v1beta3.CPU{
		Units:      TransformToResourceValue(cpu.Units),
		Attributes: MapAttributes(cpu.Attributes),
	}
}

func mapGroupMemoryToOrderMemory(mem *Memory) *v1beta3.Memory {
	if mem == nil {
		return nil
	}
	return &v1beta3.Memory{
		Quantity:   TransformToResourceValue(mem.Units),
		Attributes: MapAttributes(mem.Attributes),
	}
}

func mapGroupVolumesToOrderVolumes(vols Volumes) v1beta3.Volumes {
	mappedVols := make(v1beta3.Volumes, len(vols))
	for i, vol := range vols {
		mappedVols[i] = v1beta3.Storage{
			Name:       vol.Name,
			Quantity:   TransformToResourceValue(vol.Units),
			Attributes: MapAttributes(vol.Attributes),
		}
	}
	return mappedVols
}

func mapGroupGPUToOrderGPU(gpu *GPU) *v1beta3.GPU {
	if gpu == nil {
		return nil
	}
	return &v1beta3.GPU{
		Units:      TransformToResourceValue(gpu.Units),
		Attributes: MapAttributes(gpu.Attributes),
	}
}

func mapGroupEndpointsToOrderEndpoints(eps Endpoints) v1beta3.Endpoints {
	mappedEps := make(v1beta3.Endpoints, len(eps))
	for i, ep := range eps {
		mappedEps[i] = v1beta3.Endpoint{
			Kind:           v1beta3.Endpoint_Kind(ep.Kind),
			SequenceNumber: ep.SequenceNumber,
		}
	}
	return mappedEps
}

func MapOrderToGroup(order *Order) dtypes.Group {
	// Map OrderState to Group_State
	var state dtypes.Group_State
	switch order.State {
	case OrderOpen:
		state = dtypes.GroupOpen
	case OrderActive:
		state = dtypes.GroupPaused
	case OrderClosed:
		state = dtypes.GroupClosed
	default:
		state = dtypes.GroupStateInvalid
	}

	// Map DeploymentSpec to GroupSpec
	groupSpec := dtypes.GroupSpec{
		Name: "default", // Adjust as needed
		Requirements: v1beta3.PlacementRequirements{
			SignedBy: v1beta3.SignedBy{
				AllOf: order.Specs.PlacementsRequirement.ProviderWallets, // Assuming these map to AllOf
				AnyOf: []string{},                                        // Populate as needed
			},
			Attributes: MapAttributes(order.Specs.PlacementsRequirement.Attributes),
		},
		Resources: mapServiceResourcesToResourceUnits(order.Specs.Resources),
	}

	// Create the GroupID
	groupID := dtypes.GroupID{
		Owner: order.Creator,
		DSeq:  order.ID,
		GSeq:  1, // Adjust as needed
	}

	// Create the Group
	group := dtypes.Group{
		GroupID:   groupID,
		State:     state,
		GroupSpec: groupSpec,
		CreatedAt: 0, // Adjust as needed
	}

	return group
}


func MapProviderToV3Provider(provider *Provider) (*ptypes.Provider){
	// TODO(spheron) remove this mock data
	p := &ptypes.Provider{
		Owner:      "provider",
		HostURI:    provider.Domain,
		Attributes: types.Attributes{types.Attribute{Key: "region", Value: "us-west"}, types.Attribute{Key: "capabilities/storage/1/persistent", Value: "true"}},
	}
	return p
}