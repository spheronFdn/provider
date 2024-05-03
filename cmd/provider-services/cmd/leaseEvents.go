package cmd

import (
	"crypto/tls"
	"sync"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	cmdcommon "github.com/akash-network/node/cmd/common"
	cutils "github.com/akash-network/node/x/cert/utils"

	cltypes "github.com/akash-network/provider/cluster/types/v1beta3"
	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func leaseEventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "lease-events",
		Short:        "get lease events",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLeaseEvents(cmd)
		},
	}

	addServiceFlags(cmd)
	cmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed. Defaults to false")
	cmd.Flags().Int64P("tail", "t", -1, "The number of lines from the end of the logs to show. Defaults to -1")

	return cmd
}

func doLeaseEvents(cmd *cobra.Command) error {
	cctx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	ctx := cmd.Context()

	cl := spheron.NewClient()

	cert, err := cutils.LoadAndQueryCertificateForAccount(cmd.Context(), cctx, nil)
	if err != nil {
		return markRPCServerError(err)
	}

	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	leases, err := leasesForDeployment(cmd.Context(), *cl, cmd.Flags(), dtypes.DeploymentID{
		Owner: cctx.GetFromAddress().String(),
		DSeq:  dseq,
	})
	if err != nil {
		return markRPCServerError(err)
	}

	svcs, err := cmd.Flags().GetString(FlagService)
	if err != nil {
		return err
	}

	follow, err := cmd.Flags().GetBool(flagFollow)
	if err != nil {
		return err
	}

	type result struct {
		lid    mtypes.LeaseID
		error  error
		stream *gwrest.LeaseKubeEvents
	}

	streams := make([]result, 0, len(leases))

	for _, lid := range leases {
		stream := result{lid: lid}
		gclient, err := gwrest.NewClient(*cl, lid.Provider, []tls.Certificate{cert})
		if err == nil {
			stream.stream, stream.error = gclient.LeaseEvents(ctx, lid, svcs, follow)
		} else {
			stream.error = err
		}

		streams = append(streams, stream)
	}

	var wgStreams sync.WaitGroup
	type logEntry struct {
		cltypes.LeaseEvent `json:",inline"`
		Lid                mtypes.LeaseID `json:"lease_id"`
	}

	outch := make(chan logEntry)

	go func() {
		for evt := range outch {
			_ = cmdcommon.PrintJSON(cctx, evt)
		}
	}()

	for _, stream := range streams {
		if stream.error != nil {
			continue
		}

		wgStreams.Add(1)
		go func(stream result) {
			defer wgStreams.Done()

			for res := range stream.stream.Stream {
				outch <- logEntry{
					LeaseEvent: res,
					Lid:        stream.lid,
				}
			}
		}(stream)
	}

	wgStreams.Wait()
	close(outch)

	return nil
}
