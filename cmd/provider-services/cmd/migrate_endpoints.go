package cmd

import (
	"crypto/tls"
	"errors"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cutils "github.com/akash-network/node/x/cert/utils"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

var errEmptyEndpoints = errors.New("endpoints cannot be empty")

func migrateEndpoints(cmd *cobra.Command, args []string) error {
	endpoints := args
	if len(endpoints) == 0 {
		return errEmptyEndpoints
	}

	cctx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	cl := spheron.NewClient()

	prov, err := providerFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	cert, err := cutils.LoadAndQueryCertificateForAccount(cmd.Context(), cctx, nil)
	if err != nil {
		return markRPCServerError(err)
	}

	gclient, err := gwrest.NewClient(*cl, prov, []tls.Certificate{cert})
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

	err = gclient.MigrateEndpoints(cmd.Context(), endpoints, dseq, gseq)
	if err != nil {
		return showErrorToUser(err)
	}

	return nil

}

func MigrateEndpointsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "migrate-endpoints",
		Short:        "migrate endpoints between deployments on the same provider",
		SilenceUsage: true,
		RunE:         migrateEndpoints,
	}

	addCmdFlags(cmd)
	cmd.Flags().Uint32(FlagGSeq, 1, "group sequence")

	return cmd
}
