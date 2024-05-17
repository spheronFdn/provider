package cmd

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/akash-network/node/sdl"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

var (
	errSubmitManifestFailed = errors.New("submit manifest to some providers has been failed")
)

// SendManifestCmd looks up the Providers blockchain information,
// and POSTs the SDL file to the Gateway address.
func SendManifestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "send-manifest <sdl-path>",
		Args:         cobra.ExactArgs(1),
		Short:        "Submit manifest to provider(s)",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doSendManifest(cmd, args[0])
		},
	}

	addCmdFlags(cmd)
	cmd.Flags().StringP(flagOutput, "o", outputText, "output format text|json|yaml. default text")
	cmd.Flags().Uint64(FlagDSeq, 0, "deployment sequence")

	if err := cmd.MarkFlagRequired(FlagDSeq); err != nil {
		panic(err.Error())
	}
	return cmd
}

func doSendManifest(cmd *cobra.Command, sdlpath string) error {
	ctx := context.TODO()
	sdl, err := sdl.ReadFile(sdlpath)
	if err != nil {
		return err
	}

	mani, err := sdl.Manifest()
	if err != nil {
		return err
	}

	// TODO(spheron): cert, err := cutils.LoadAndQueryCertificateForAccount(cmd.Context(), cctx, nil)
	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	cctx, err := spheron.GetClientTxContext(cmd)
	spClient := spheron.NewClientWithContext(cctx)

	submitFailed := false

	lease, err := spClient.BcClient.GetLeaseById(ctx, dseq)
	if err != nil {
		return err
	}
	fmt.Printf("Lease %v+", lease)
	// if lease.State != entities.OrderActive {
	// 	return errors.New("Lease is not active")
	// }

	authToken, err := spheron.CreateAuthorizationToken(ctx, &cctx)
	if err != nil {
		return err
	}
	gclient, err := gwrest.NewClient(*spClient, lease.Provider, authToken)
	if err != nil {
		return err
	}
	err = gclient.SubmitManifest(cmd.Context(), dseq, mani)

	if err != nil {
		return err
	}

	_, err = fmt.Println("Manifest sent")

	if err != nil {
		return err
	}

	if submitFailed {
		return errSubmitManifestFailed
	}

	return nil
}
