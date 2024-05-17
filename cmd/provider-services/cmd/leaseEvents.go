package cmd

import (
	"sync"

	"github.com/spf13/cobra"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"

	cltypes "github.com/akash-network/provider/cluster/types/v1beta3"
	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func leaseEventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "deployment-events",
		Short:        "get deployment events",
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
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	ctx := cmd.Context()

	cl := spheron.NewClientWithContext(cctx)

	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	leases, err := leasesForDeployment(cmd.Context(), *cl, cmd.Flags(), dtypes.DeploymentID{
		Owner: cl.Context.Key.Address.Hex(),
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

	authToken, err := spheron.CreateAuthorizationToken(ctx, &cctx)
	if err != nil {
		return err
	}
	for _, lid := range leases {
		stream := result{lid: lid}
		gclient, err := gwrest.NewClient(*cl, lid.Provider, authToken)
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
			_ = spheron.PrintJSON(evt)
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
