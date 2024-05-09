package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

var errEmptyHostnames = errors.New("hostnames cannot be empty")

func migrateHostnames(cmd *cobra.Command, args []string) error {
	hostnames := args
	if len(hostnames) == 0 {
		return errEmptyHostnames
	}
	cctx, err := spheron.GetClientTxContext(cmd)
	cl := spheron.NewClientWithContext(cctx)

	prov, err := providerFromFlags(cmd.Flags())
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

	dseq, err := cmd.Flags().GetUint64("dseq")
	if err != nil {
		return err
	}

	gseq, err := cmd.Flags().GetUint32("gseq")
	if err != nil {
		return err
	}

	err = gclient.MigrateHostnames(cmd.Context(), hostnames, dseq, gseq)
	if err != nil {
		return showErrorToUser(err)
	}

	return nil
}

func MigrateHostnamesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "migrate-hostnames",
		Short:        "migrate hostnames between deployments on the same provider",
		SilenceUsage: true,
		RunE:         migrateHostnames,
	}

	addCmdFlags(cmd)
	cmd.Flags().Uint32(FlagGSeq, 1, "group sequence")

	return cmd
}
