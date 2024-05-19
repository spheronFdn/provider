package event

import (
	mani "github.com/akash-network/akash-api/go/manifest/v2beta2"
	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
)

// LeaseWon is the data structure that includes leaseID, group and price
type LeaseWon struct {
	LeaseID mtypes.LeaseID
	Group   *dtypes.Group
}

// ManifestReceived stores leaseID, manifest received, deployment and group details
// to be provisioned by the Provider.
type ManifestReceived struct {
	LeaseID  mtypes.LeaseID
	Manifest *mani.Manifest
	Group    *dtypes.Group
}

// ManifestGroup returns group if present in manifest or nil
func (ev ManifestReceived) ManifestGroup() *mani.Group {
	group := ev.Manifest.GetGroups()[0]
	return &group
}

// ClusterDeploymentStatus represents status of the cluster deployment
type ClusterDeploymentStatus string

const (
	// ClusterDeploymentUpdated is used whenever the deployment in the cluster is updated but may not be functional
	ClusterDeploymentUpdated ClusterDeploymentStatus = "updated"
	// ClusterDeploymentPending is used when cluster deployment status is pending
	ClusterDeploymentPending ClusterDeploymentStatus = "pending"
	// ClusterDeploymentDeployed is used when cluster deployment status is deployed
	ClusterDeploymentDeployed ClusterDeploymentStatus = "deployed"
)

// ClusterDeployment stores leaseID, group details and deployment status
type ClusterDeployment struct {
	LeaseID mtypes.LeaseID
	Group   *mani.Group
	Status  ClusterDeploymentStatus
}

type LeaseAddFundsMonitor struct {
	mtypes.LeaseID
	IsNewLease bool
}

type LeaseRemoveFundsMonitor struct {
	mtypes.LeaseID
}
