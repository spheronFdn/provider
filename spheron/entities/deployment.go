package entities

type Deployment struct {
	ID        DeploymentID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Deployment_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta3.Group_State" json:"state" yaml:"state"`
	Spec      DeploymentSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
	CreatedAt int64            `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

type DeploymentID struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
}

type Deployment_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	DeploymentStateInvalid Deployment_State = 0
	// DeploymentOpen denotes state for group open
	DeploymentOpen Deployment_State = 1
	// DeploymentOrdered denotes state for group ordered
	DeploymentPaused Deployment_State = 2
	// DeploymentInsufficientFunds denotes state for group insufficient_funds
	DeploymentInsufficientFunds Deployment_State = 3
	// DeploymentClosed denotes state for group closed
	DeploymentClosed Deployment_State = 4
)

type DeploymentSpec struct {
	Name         string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Requirements PlacementRequirements `protobuf:"bytes,2,opt,name=requirements,proto3" json:"requirements" yaml:"requirements"`
	Resources    ResourceUnits         `protobuf:"bytes,3,rep,name=resources,proto3,castrepeated=ResourceUnits" json:"resources" yaml:"resources"`
}

type PlacementRequirements struct {
	// SignedBy list of keys that tenants expect to have signatures from
	SignedBy SignedBy `protobuf:"bytes,1,opt,name=signed_by,json=signedBy,proto3" json:"signed_by" yaml:"signed_by"`
	// Attribute list of attributes tenant expects from the provider
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes" yaml:"attributes"`
}

type SignedBy struct {
	// all_of all keys in this list must have signed attributes
	AllOf []string `protobuf:"bytes,1,rep,name=all_of,json=allOf,proto3" json:"all_of" yaml:"allOf"`
	// any_of at least of of the keys from the list must have signed attributes
	AnyOf []string `protobuf:"bytes,2,rep,name=any_of,json=anyOf,proto3" json:"any_of" yaml:"anyOf"`
}

type Attributes []Attribute

type Attribute struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
}

type ResourceUnits []ResourceUnit

type ResourceUnit struct {
	Resources `protobuf:"bytes,1,opt,name=resource,proto3,embedded=resource" json:"resource" yaml:"resource"`
	Count     uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price     uint64 `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
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

type Memory struct {
	Units      uint64     `protobuf:"bytes,1,opt,name=units,proto3" json:"units" yaml:"units"`
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Volumes []Storage

type Storage struct {
	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Units      uint64     `protobuf:"bytes,2,opt,name=units,proto3" json:"units" yaml:"units"`
	Attributes Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type GPU struct {
	Units      uint64     `protobuf:"bytes,1,opt,name=units,proto3" json:"units"`
	Attributes Attributes `protobuf:"bytes,2,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Endpoints []Endpoint

type Endpoint struct {
	Kind           int32  `protobuf:"varint,1,opt,name=kind,proto3,enum=akash.base.v1beta3.Endpoint_Kind" json:"kind,omitempty"`
	SequenceNumber uint32 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number" yaml:"sequence_number"`
}
