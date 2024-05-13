package cmd

import (
	"context"

	"github.com/spf13/cobra"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func leaseStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "lease-status",
		Short:        "get lease status",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLeaseStatus(cmd)
		},
	}

	addLeaseFlags(cmd)

	return cmd
}

func doLeaseStatus(cmd *cobra.Command) error {

	cctx, err := spheron.GetClientTxContext(cmd)

	cl := spheron.NewClientWithContext(cctx)

	prov, err := providerFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	bid, err := spheron.BidIDFromFlags(cmd.Flags(), spheron.WithOwner(cctx.Key.Address.Hex()))
	if err != nil {
		return err
	}

	if err != nil {
		return markRPCServerError(err)
	}

	authToken, err := spheron.CreateAuthorizationToken(context.TODO(), &cctx)
	if err != nil {
		return err
	}

	gclient, err := gwrest.NewClient(*cl, prov, authToken)
	if err != nil {
		return err
	}

	result, err := gclient.LeaseStatus(cmd.Context(), bid.LeaseID())
	if err != nil {
		return showErrorToUser(err)
	}

	return spheron.PrintJSON(result)
}
