package ientities

import (
	"github.com/akash-network/provider/spheron/entities"
)

type ResourcesMetric struct {
	CPU              uint64 `json:"cpu"`
	GPU              uint64 `json:"gpu"`
	Memory           uint64 `json:"memory"`
	StorageEphemeral uint64 `json:"storage_ephemeral"`
}

type NodeMetrics struct {
	Name        string          `json:"name"`
	Allocatable ResourcesMetric `json:"allocatable"`
	Available   ResourcesMetric `json:"available"`
}

type Metrics struct {
	Nodes            []NodeMetrics `json:"nodes"`
	TotalAllocatable MetricTotal   `json:"total_allocatable"`
	TotalAvailable   MetricTotal   `json:"total_available"`
}

type MetricTotal struct {
	CPU              uint64           `json:"cpu"`
	GPU              uint64           `json:"gpu"`
	Memory           uint64           `json:"memory"`
	StorageEphemeral uint64           `json:"storage_ephemeral"`
	Storage          map[string]int64 `json:"storage,omitempty"`
}

type StorageStatus struct {
	Class string `json:"class"`
	Size  int64  `json:"size"`
}

// InventoryMetrics stores active, pending and available units
type InventoryMetrics struct {
	Active    []MetricTotal `json:"active,omitempty"`
	Pending   []MetricTotal `json:"pending,omitempty"`
	Available struct {
		Nodes   []NodeMetrics   `json:"nodes,omitempty"`
		Storage []StorageStatus `json:"storage,omitempty"`
	} `json:"available,omitempty"`
	Error error `json:"error,omitempty"`
}

func (inv *MetricTotal) AddResources(res entities.ResourceUnit) {
	cpu := inv.CPU
	gpu := inv.GPU
	mem := inv.Memory
	ephemeralStorage := inv.StorageEphemeral

	if res.CPU != nil {
		cpu = cpu + (res.CPU.Units * uint64(res.Count))
	}

	if res.GPU != nil {
		gpu = gpu + (res.GPU.Units * uint64(res.Count))
	}

	if res.Memory != nil {
		mem = mem + (res.Memory.Units * uint64(res.Count))
	}

	for _, storage := range res.Storage {
		if storageClass, found := storage.Attributes.Find("class").AsString(); !found {
			ephemeralStorage = ephemeralStorage + (storage.Units * uint64(res.Count))
		} else {
			val := uint64(inv.Storage[storageClass])
			val = val + (storage.Units * uint64(res.Count))
			inv.Storage[storageClass] = int64(val)
		}
	}

	inv.CPU = cpu
	inv.GPU = gpu
	inv.Memory = mem
	inv.StorageEphemeral = ephemeralStorage
}
