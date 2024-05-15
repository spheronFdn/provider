package events

import (
	"encoding/hex"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"

	"github.com/akash-network/akash-api/go/sdkutil"
	"github.com/cosmos/cosmos-sdk/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func MapOrderCreated(event *OrderCreated) mtypes.EventOrderCreated {
	return mtypes.EventOrderCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "order-created"},
		ID: mtypes.OrderID{
			Owner: "owner", // TODO(spheron): Check if we need this field
			DSeq:  event.ID,
			GSeq:  1,
			OSeq:  1,
		}}
}

func MapOrderMatched(event *OrderMatched) mtypes.EventLeaseCreated {
	return mtypes.EventLeaseCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-created"}, ID: mtypes.LeaseID{
		Owner:    "owner", // TODO(spheron): check if this is needed
		DSeq:     event.ID,
		GSeq:     1,
		OSeq:     1,
		Provider: event.Provider,
	}, Price: cosmostypes.DecCoin{
		Denom:  "uakt",
		Amount: cosmostypes.OneDec(),
	}}
}

func MapOrderUpdateRequested(event *OrderUpdateRequest) dtypes.EventDeploymentUpdated {
	v, _ := hex.DecodeString("1")

	return dtypes.EventDeploymentUpdated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-updated"}, ID: dtypes.DeploymentID{
		Owner: "owner",
		DSeq:  15,
	}, Version: v,
	}
}

func MapOrderClosed(event *OrderClosed) (dtypes.EventDeploymentClosed, mtypes.EventLeaseClosed) {
	return dtypes.EventDeploymentClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-close"}, ID: dtypes.DeploymentID{
			Owner: "owner",
			DSeq:  event.ID,
		}}, mtypes.EventLeaseClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-close"}, ID: mtypes.LeaseID{
			Owner:    "owner",
			DSeq:     event.ID,
			GSeq:     1,
			OSeq:     1,
			Provider: "provider",
		}, Price: cosmostypes.DecCoin{
			Denom:  "uakt",
			Amount: types.OneDec(),
		}}
}
