package entities

type ResourceUnits []ResourceUnit

type ResourceGroup interface {
	GetName() string
	GetResourceUnits() ResourceUnits
}

type ResourceUnit struct {
	Resources `protobuf:"bytes,1,opt,name=resource,proto3,embedded=resource" json:"resource" yaml:"resource"`
	Count     uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price     uint64 `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (r *ResourceUnit) FullPrice() uint64 {
	return r.Price * uint64(r.Count) //TODO (spheron): check on this
}

func (m *Resources) GetCPU() *CPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *Resources) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Resources) GetStorage() Volumes {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Resources) GetGPU() *GPU {
	if m != nil {
		return m.GPU
	}
	return nil
}

func (m *Resources) GetEndpoints() Endpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Resources struct {
	ID        uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id" yaml:"id"`
	CPU       *CPU      `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory    *Memory   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty" yaml:"memory,omitempty"`
	Storage   Volumes   `protobuf:"bytes,4,rep,name=storage,proto3,castrepeated=Volumes" json:"storage,omitempty" yaml:"storage,omitempty"`
	GPU       *GPU      `protobuf:"bytes,5,opt,name=gpu,proto3" json:"gpu,omitempty" yaml:"gpu,omitempty"`
	Endpoints Endpoints `protobuf:"bytes,6,rep,name=endpoints,proto3,castrepeated=Endpoints" json:"endpoints" yaml:"endpoints"`
}

type CPU struct {
	Units      uint64     `protobuf:"bytes,1,opt,name=units,proto3" json:"units"`
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *CPU) GetUnits() uint64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *CPU) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Memory struct {
	Units      uint64     `protobuf:"bytes,1,opt,name=units,proto3" json:"units" yaml:"units"`
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *Memory) GetUnits() uint64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Memory) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Volumes []Storage

type Storage struct {
	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Units      uint64     `protobuf:"bytes,2,opt,name=units,proto3" json:"units" yaml:"units"`
	Attributes Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *Storage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Storage) GetUnits() uint64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Storage) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GPU struct {
	Units      uint64     `protobuf:"bytes,1,opt,name=units,proto3" json:"units"`
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *GPU) GetUnits() uint64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *GPU) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Endpoints []Endpoint

type Endpoint struct {
	Kind           int32  `protobuf:"varint,1,opt,name=kind,proto3,enum=akash.base.v1beta3.Endpoint_Kind" json:"kind,omitempty"`
	SequenceNumber uint32 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number" yaml:"sequence_number"`
}

// This describes how the endpoint is implemented when the lease is deployed
type Endpoint_Kind int32

const (
	// Describes an endpoint that becomes a Kubernetes Ingress
	Endpoint_SHARED_HTTP Endpoint_Kind = 0
	// Describes an endpoint that becomes a Kubernetes NodePort
	Endpoint_RANDOM_PORT Endpoint_Kind = 1
	// Describes an endpoint that becomes a leased IP
	Endpoint_LEASED_IP Endpoint_Kind = 2
)

func (m Resources) Dup() Resources {
	res := Resources{
		ID:        m.ID,
		CPU:       m.CPU.Dup(),
		GPU:       m.GPU.Dup(),
		Memory:    m.Memory.Dup(),
		Storage:   m.Storage.Dup(),
		Endpoints: m.Endpoints.Dup(),
	}

	return res
}

func (u Endpoints) Dup() Endpoints {
	res := make(Endpoints, len(u))

	copy(res, u)

	return res
}

func (m CPU) Dup() *CPU {
	return &CPU{
		Units:      m.Units,
		Attributes: m.Attributes.Dup(),
	}
}

func (m Memory) Dup() *Memory {
	return &Memory{
		Units:      m.Units,
		Attributes: m.Attributes.Dup(),
	}
}

func (m Storage) Dup() *Storage {
	return &Storage{
		Name:       m.Name,
		Units:      m.Units,
		Attributes: m.Attributes.Dup(),
	}
}

func (m GPU) Dup() *GPU {
	return &GPU{
		Units:      m.Units,
		Attributes: m.Attributes.Dup(),
	}
}

func (attr Attributes) Dup() Attributes {
	if attr == nil {
		return nil
	}

	res := make(Attributes, 0, len(attr))

	for _, pair := range attr {
		res = append(res, Attribute{
			Key:   pair.Key,
			Value: pair.Value,
		})
	}

	return res
}

func (m Volumes) Dup() Volumes {
	res := make(Volumes, 0, len(m))

	for _, storage := range m {
		res = append(res, *storage.Dup())
	}

	return res
}

func (m *CPU) EqualUnits(that *CPU) bool {
	if that == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.Units != that.Units {
		return false
	}

	return true
}

func (m *GPU) EqualUnits(that *GPU) bool {
	if that == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.Units != that.Units {
		return false
	}

	return true
}

func (m *Memory) EqualUnits(that *Memory) bool {
	if that == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.Units != that.Units {
		return false
	}

	return true
}

func (m Volumes) EqualUnits(that Volumes) bool {
	if len(m) != len(that) {
		return false
	}

	for idx, vol := range m {
		if vol.Name != that[idx].Name {
			return false
		}

		if vol.Units != that[idx].Units {
			return false
		}

	}

	return true
}
