package entities

type Order struct {
	ID         uint64
	Region     string
	Uptime     uint64
	Reputation uint64
	Slashes    uint64
	MaxPrice   uint64
	Token      string
	Creator    string
	State      OrderState
	Specs      DeploymentSpec
}

type DeploymentID struct {
	Owner string
	DSeq  uint64
}

type OrderState int32

const (
	OrderInvalid OrderState = iota + 1 // OrderInvalid = 1
	OrderOpen
	OrderActive
	OrderClosed
)

type DeploymentSpec struct {
	PlacementsRequirement PlacementRequirements
	Resources             ServiceResources
}

type PlacementRequirements struct {
	// SignedBy list of wallets that tenants wants to take bids
	ProviderWallets []string
	// Attribute list of attributes tenant expects from the provider
	Attributes Attributes
}

type ServiceResources []ServiceResource

type ServiceResource struct {
	Resources    Resources
	ReplicaCount uint32
}

type Resources struct {
	ID        uint32
	CPU       *CPU
	Memory    *Memory
	Storage   Volumes
	GPU       *GPU
	Endpoints Endpoints
}

type CPU struct {
	Units      uint64
	Attributes Attributes
}

type Memory struct {
	Units      uint64
	Attributes Attributes
}

type Volumes []Volume

type Volume struct {
	Name       string
	Units      uint64
	Attributes Attributes
}

type GPU struct {
	Units      uint64
	Attributes Attributes
}

type Endpoints []Endpoint

type Endpoint struct {
	Kind           EndpointKind
	SequenceNumber uint32
}

type EndpointKind int32

const (
	// Describes an endpoint that becomes a Kubernetes Ingress
	EndpointSharedHttp EndpointKind = 0
	// Describes an endpoint that becomes a Kubernetes NodePort
	EndpointRandomPort EndpointKind = 1
	// Describes an endpoint that becomes a leased IP
	EndpointLeasedIP EndpointKind = 2
)
