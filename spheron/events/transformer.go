package events

import (
	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"

	"github.com/akash-network/akash-api/go/sdkutil"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func MapOrderCreated(event *OrderCreated) mtypes.EventOrderCreated {
	return mtypes.EventOrderCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "order-created"},
		ID: mtypes.OrderID{
			Owner: event.Creator,
			DSeq:  event.ID,
			GSeq:  1,
			OSeq:  1,
		}}
}

func MapOrderMatched(event *OrderMatched) mtypes.EventLeaseCreated {
	return mtypes.EventLeaseCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-created"}, ID: mtypes.LeaseID{
		Owner:    event.Creator,
		DSeq:     event.ID,
		GSeq:     1,
		OSeq:     1,
		Provider: event.Provider,
	}, Price: cosmostypes.DecCoin{},
	}
}

func MapOrderClosed(event *OrderClosed) (dtypes.EventDeploymentClosed, mtypes.EventLeaseClosed) {
	return dtypes.EventDeploymentClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-close"}, ID: dtypes.DeploymentID{
			Owner: event.Creator,
			DSeq:  event.ID,
		}}, mtypes.EventLeaseClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-close"}, ID: mtypes.LeaseID{
			Owner:    event.Creator,
			DSeq:     event.ID,
			GSeq:     1,
			OSeq:     1,
			Provider: event.Provider,
		}, Price: cosmostypes.DecCoin{},
		}
}
