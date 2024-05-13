package entities

type MsgCreateBid struct {
	Order          DeploymentID   `protobuf:"bytes,1,opt,name=order,proto3" json:"order" yaml:"order"`
	Provider       string         `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider" yaml:"provider"`
	Price          uint64         `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
	ResourcesOffer ResourcesOffer `protobuf:"bytes,5,rep,name=resources_offer,json=resourcesOffer,proto3,castrepeated=ResourcesOffer" json:"resources_offer" yaml:"resources_offer"`
}
