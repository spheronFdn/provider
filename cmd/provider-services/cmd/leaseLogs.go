package cmd

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
)

func leaseLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "lease-logs",
		Short:        "get lease logs",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLeaseLogs(cmd)
		},
	}

	addServiceFlags(cmd)

	cmd.Flags().BoolP(flagFollow, "f", false, "Specify if the logs should be streamed. Defaults to false")
	cmd.Flags().Int64P(flagTail, "t", -1, "The number of lines from the end of the logs to show. Defaults to -1")
	cmd.Flags().StringP(flagOutput, "o", outputText, "Output format text|json. Defaults to text")

	return cmd
}

func doLeaseLogs(cmd *cobra.Command) error {

	ctx := cmd.Context()

	cl := spheron.NewClient()

	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	leases, err := leasesForDeployment(cmd.Context(), *cl, cmd.Flags(), dtypes.DeploymentID{
		//TODO(spheron) get this value from context or env
		Owner: "owner",
		DSeq:  dseq,
	})
	if err != nil {
		return markRPCServerError(err)
	}

	svcs, err := cmd.Flags().GetString(FlagService)
	if err != nil {
		return err
	}

	outputFormat, err := cmd.Flags().GetString(flagOutput)
	if err != nil {
		return err
	}

	if outputFormat != outputText && outputFormat != outputJSON {
		return errors.Errorf("invalid output format %s. expected text|json", outputFormat)
	}

	follow, err := cmd.Flags().GetBool(flagFollow)
	if err != nil {
		return err
	}

	tailLines, err := cmd.Flags().GetInt64(flagTail)
	if err != nil {
		return err
	}

	if tailLines < -1 {
		return errors.Errorf("tail flag supplied with invalid value. must be >= -1")
	}

	type result struct {
		lid    mtypes.LeaseID
		error  error
		stream *gwrest.ServiceLogs
	}

	streams := make([]result, 0, len(leases))
	for _, lid := range leases {
		stream := result{lid: lid}
		authToken, err := spheron.CreateAuthorizationToken(ctx)
		if err != nil {
			return err
		}
		gclient, err := gwrest.NewClient(*cl, lid.Provider, authToken)
		if err == nil {
			stream.stream, stream.error = gclient.LeaseLogs(ctx, lid, svcs, follow, tailLines)
		} else {
			stream.error = err
		}

		streams = append(streams, stream)
	}

	var wgStreams sync.WaitGroup

	type logEntry struct {
		gwrest.ServiceLogMessage `json:",inline"`
		Lid                      mtypes.LeaseID `json:"lease_id"`
	}

	outch := make(chan logEntry)

	printFn := func(evt logEntry) {
		fmt.Printf("[%s][%s] %s\n", evt.Lid, evt.Name, evt.Message)
	}

	if outputFormat == "json" {
		printFn = func(evt logEntry) {
			_ = spheron.PrintJSON(evt)
		}
	}

	go func() {
		for evt := range outch {
			printFn(evt)
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
					ServiceLogMessage: res,
					Lid:               stream.lid,
				}
			}
		}(stream)
	}

	wgStreams.Wait()
	close(outch)

	return nil
}
