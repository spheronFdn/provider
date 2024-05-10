package entities

import "strconv"

type Deployment struct {
	ID        DeploymentID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Deployment_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta3.Group_State" json:"state" yaml:"state"`
	Spec      DeploymentSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
	CreatedAt int64            `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

// GetName method returns group name
func (d Deployment) GetName() string {
	return d.Spec.Name
}

// GetResourceUnits method returns resources list in group
func (d Deployment) GetResourceUnits() ResourceUnits {
	return d.Spec.Resources
}

type DeploymentID struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
}

// Equals method compares specific order with provided order
func (id DeploymentID) Equals(other DeploymentID) bool {
	return id.Owner == other.Owner && id.DSeq == other.DSeq
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

// GetResourceUnits method returns resources list in group
func (g DeploymentSpec) GetResourceUnits() ResourceUnits {
	resources := make(ResourceUnits, 0, len(g.Resources))

	resources = append(resources, g.Resources...)

	return resources
}

// GetName method returns group name
func (g DeploymentSpec) GetName() string {
	return g.Name
}

// Price method returns price of group
func (g DeploymentSpec) Price() uint64 {
	var price uint64
	for idx, resource := range g.Resources {
		if idx == 0 {
			price = resource.FullPrice()
			continue
		}
		price = price + resource.FullPrice()
	}
	return price
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

func (attr Attributes) Find(glob string) AttributeValue {
	// todo wildcard

	var val attributeValue

	for i := range attr {
		if glob == attr[i].Key {
			val.value = attr[i].Value
			break
		}
	}

	return val
}

type attributeValue struct {
	value string
}

func (val attributeValue) AsBool() (bool, bool) {
	if val.value == "" {
		return false, false
	}

	res, err := strconv.ParseBool(val.value)
	if err != nil {
		return false, false
	}

	return res, true
}

func (val attributeValue) AsString() (string, bool) {
	if val.value == "" {
		return "", false
	}

	return val.value, true
}

type AttributeValue interface {
	AsBool() (bool, bool)
	AsString() (string, bool)
}
