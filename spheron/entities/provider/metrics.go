package v1

import (
	"github.com/akash-network/provider/spheron/entities"
	"k8s.io/apimachinery/pkg/api/resource"
)

func NewResourcesMetric() ResourcesMetric {
	return ResourcesMetric{
		CPU:              resource.NewMilliQuantity(0, resource.DecimalSI),
		Memory:           resource.NewQuantity(0, resource.DecimalSI),
		GPU:              resource.NewQuantity(0, resource.DecimalSI),
		EphemeralStorage: resource.NewQuantity(0, resource.DecimalSI),
		Storage:          make(Storage),
	}
}
func (inv *ResourcesMetric) AddResources(res entities.Resources) {
	if res.CPU != nil {
		qcpu := *resource.NewMilliQuantity(int64(res.CPU.Units), resource.DecimalSI)
		inv.CPU.Add(qcpu)
	}

	if res.GPU != nil {
		qcpu := *resource.NewQuantity(int64(res.GPU.Units), resource.DecimalSI)
		inv.CPU.Add(qcpu)
	}

	if res.Memory != nil {
		qcpu := *resource.NewQuantity(int64(res.Memory.Units), resource.DecimalSI)
		inv.CPU.Add(qcpu)
	}

	for _, storage := range res.Storage {
		val := *resource.NewQuantity(int64(storage.Units), resource.DecimalSI)
		if storageClass, found := storage.Attributes.Find("class").AsString(); !found {
			inv.EphemeralStorage.Add(val)
		} else {
			inv.Storage[storageClass].Add(val)
		}
	}
}

func (inv *ResourcesMetric) AddResourceUnit(res entities.ResourceUnit) {
	if res.CPU != nil {
		val := int64(res.CPU.Units)
		val = val * int64(res.Count)

		qcpu := *resource.NewMilliQuantity(val, resource.DecimalSI)
		inv.CPU.Add(qcpu)
	}

	if res.GPU != nil {
		val := int64(res.GPU.Units)
		val = val * int64(res.Count)

		qgpu := *resource.NewQuantity(val, resource.DecimalSI)
		inv.GPU.Add(qgpu)
	}

	if res.Memory != nil {
		val := int64(res.Memory.Units)
		val = val * int64(res.Count)

		qmem := *resource.NewQuantity(val, resource.DecimalSI)
		inv.Memory.Add(qmem)
	}

	for _, storage := range res.Storage {
		val := int64(storage.Units)
		val = val * int64(res.Count)

		qstorage := *resource.NewQuantity(val, resource.DecimalSI)

		if storageClass, found := storage.Attributes.Find("class").AsString(); !found {
			inv.EphemeralStorage.Add(qstorage)
		} else {
			if _, exists := inv.Storage[storageClass]; !exists {
				inv.Storage[storageClass] = resource.NewQuantity(0, resource.DecimalSI)
			}

			inv.Storage[storageClass].Add(qstorage)
		}
	}
}

func (inv *ResourcesMetric) AddResourceUnits(res entities.ResourceUnits) {
	for _, unit := range res {
		inv.AddResourceUnit(unit)
	}
}
