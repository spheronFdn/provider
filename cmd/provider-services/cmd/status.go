package cmd

import (
	"context"

	"github.com/spf13/cobra"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func statusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "status [address]",
		Short:        "get provider status",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Spheron fix: take address as a string
			addr := args[0]
			return doStatus(cmd, addr)
		},
	}

	return cmd
}

func doStatus(cmd *cobra.Command, addr string) error {
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	cl := spheron.NewClientWithContext(cctx)

	authToken, err := spheron.CreateAuthorizationToken(context.TODO(), &cctx)
	if err != nil {
		return err
	}

	gclient, err := gwrest.NewClient(*cl, addr, authToken)
	if err != nil {
		return err
	}

	result, err := gclient.Status(cmd.Context())
	if err != nil {
		return showErrorToUser(err)
	}

	return spheron.PrintJSON(result)
}
