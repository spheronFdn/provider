package cmd

import (
	"context"

	"github.com/spf13/cobra"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func serviceStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "service-status",
		Short:        "get service status",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doServiceStatus(cmd)
		},
	}

	addServiceFlags(cmd)
	if err := cmd.MarkFlagRequired(FlagService); err != nil {
		panic(err.Error())
	}

	return cmd
}

func doServiceStatus(cmd *cobra.Command) error {
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	cl := spheron.NewClientWithContext(cctx)

	if err != nil {
		return err
	}

	svcName, err := cmd.Flags().GetString(FlagService)
	if err != nil {
		return err
	}

	prov, err := providerFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	bid, err := spheron.BidIDFromFlags(cmd.Flags(), spheron.WithOwner(cctx.Key.Address.Hex()))
	if err != nil {
		return err
	}

	authToken, err := spheron.CreateAuthorizationToken(context.TODO(), &cctx)
	if err != nil {
		return err
	}

	gclient, err := gwrest.NewClient(*cl, prov, authToken)
	if err != nil {
		return err
	}

	result, err := gclient.ServiceStatus(cmd.Context(), bid.LeaseID(), svcName)
	if err != nil {
		return showErrorToUser(err)
	}

	return spheron.PrintJSON(result)
}
