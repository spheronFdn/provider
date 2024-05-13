package entities

import "github.com/akash-network/akash-api/go/node/escrow/v1beta3"

// QueryDeploymentResponse is response type for the Query/Deployment RPC method
type QueryDeploymentResponse struct {
	Deployments   []Deployment    `protobuf:"bytes,2,rep,name=groups,proto3" json:"deployments" yaml:"deployments"`
	EscrowAccount v1beta3.Account `protobuf:"bytes,3,opt,name=escrow_account,json=escrowAccount,proto3" json:"escrow_account"`
}
