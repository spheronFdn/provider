package blockchain

import (
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/ethereum/go-ethereum/common"
)

func stringSliceToAddressSlice(stringSlice []string) []common.Address {
	addressSlice := make([]common.Address, len(stringSlice))
	for i, str := range stringSlice {
		addressSlice[i] = common.HexToAddress(str) // Assuming common.Address has a constructor that accepts a string
	}
	return addressSlice
}

func getMatchingResourceAttribute(res entities.Resources) OrderMatching.OrderMatchingResourceAttributes {
	return OrderMatching.OrderMatchingResourceAttributes{
		CpuUnits:      int64(res.CPU.Units),
		CpuAttributes: nil,
		RamUnits:      int64(res.Memory.Units),
		RamAttributes: nil,
		Volume: OrderMatching.OrderMatchingVolume{
			Name:       res.Storage[0].Name,
			Units:      int64(res.Storage[0].Units),
			Attributes: nil,
		},
		GpuUnits:                int64(res.GPU.Units),
		GpuAttributes:           nil,
		EndpointsKind:           int32(res.Endpoints[0].Kind),
		EndpointsSequenceNumber: int32(res.Endpoints[0].SequenceNumber),
	}
}
