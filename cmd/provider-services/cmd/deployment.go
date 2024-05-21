package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	gwrest "github.com/akash-network/provider/gateway/rest"
	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/sdl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	orderMatchedEvent *OrderMatching.OrderMatchingOrderMatched
	waitForOrderMatch sync.WaitGroup
)

func DeploymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deployment",
		Short: "Create deployment request",
	}

	cmd.AddCommand(
		AddDeploymentCommand(),
		CloseDeploymentCommand(),
		MatchDeploymentCommand(),
	)
	return cmd
}

func AddDeploymentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [sdl-file]",
		Short: "Deployment",
		Args:  cobra.ExactArgs(1),
		RunE:  runDeploymentCmd,
	}
	addCmdFlags(cmd)

	return cmd
}

func runDeploymentCmd(cmd *cobra.Command, args []string) error {
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	if cctx.Key == nil {
		return fmt.Errorf("Transaction can not be created. Wallet needs to be injected")
	}

	spCl := spheron.NewClientWithContext(cctx)

	sdlManifest, err := sdl.ReadFile(args[0])
	if err != nil {
		return err
	}

	// TODO(spheron): make it so SDL can have only 1 group
	groups, err := sdlManifest.DeploymentGroups()
	if err != nil {
		return err
	}

	order := entities.TransformGroupToOrder(groups[0])

	fmt.Println("Waiting for ptovider bids..")

	_, err = spCl.BcClient.CreateOrder(context.TODO(), order)
	if err != nil {
		return fmt.Errorf("Error while creating Deployment transaction")
	}

	// Subscribe to order matched event and wait
	waitForOrderMatch.Add(1)
	go waitForOrderMatchedEvent(cctx, spCl, sdlManifest)
	waitForOrderMatch.Wait()

	return nil
}

func waitForOrderMatchedEvent(cctx spheron.Context, spCl *spheron.Client, sdl sdl.SDL) {
	defer waitForOrderMatch.Done()

	// Assume this function subscribes and waits for the OrderMatchingOrderMatched event
	orderMatchedEvent := <-spCl.BcClient.SubscribeToOrderMatched()

	fmt.Println("Bid found.")

	if orderMatchedEvent != nil {
		fmt.Println("Sending manifest.")

		sendManifest(cctx, spCl, sdl, orderMatchedEvent)
	}
}

func sendManifest(cctx spheron.Context, spClient *spheron.Client, sdl sdl.SDL, orderMatchedEvent *OrderMatching.OrderMatchingOrderMatched) error {
	ctx := context.TODO()

	mani, err := sdl.Manifest()
	if err != nil {
		return err
	}

	submitFailed := false

	lease, err := spClient.BcClient.GetLeaseById(ctx, orderMatchedEvent.OrderId)
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
	err = gclient.SubmitManifest(ctx, orderMatchedEvent.OrderId, mani)

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

func CloseDeploymentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close",
		Short: "Close",
		Args:  cobra.ExactArgs(0),
		RunE:  runCloseDeploymentCmd,
	}
	addCmdFlags(cmd)
	cmd.Flags().Uint64(FlagDSeq, 0, "deployment sequence")

	if err := cmd.MarkFlagRequired(FlagDSeq); err != nil {
		panic(err.Error())
	}
	return cmd
}

func runCloseDeploymentCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return fmt.Errorf("Invalid dseq: %v", err)
	}

	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	if cctx.Key == nil {
		return fmt.Errorf("Transaction can not be created. Wallet needs to be injected")
	}

	spCl := spheron.NewClientWithContext(cctx)

	_, err = spCl.BcClient.CloseOrder(ctx, dseq)

	if err != nil {
		return fmt.Errorf("Error while closing Deployment")
	}
	return nil
}

func MatchDeploymentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "match [provider] [agreedPrice]",
		Short: "Match deployment with provider",
		Args:  cobra.ExactArgs(2),
		RunE:  runMatchDeploymentCmd,
	}
	addCmdFlags(cmd)
	cmd.Flags().Uint64(FlagDSeq, 0, "deployment sequence")

	if err := cmd.MarkFlagRequired(FlagDSeq); err != nil {
		panic(err.Error())
	}

	return cmd
}

func runMatchDeploymentCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return fmt.Errorf("Invalid dseq: %v", err)
	}
	provider := args[0]
	ap, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("Invalid agreedPrice: %v", err)
	}
	agreedPrice := big.NewInt(ap)

	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	if cctx.Key == nil {
		return fmt.Errorf("Transaction can not be created. Wallet needs to be injected")
	}

	spCl := spheron.NewClientWithContext(cctx)

	_, err = spCl.BcClient.MatchOrder(ctx, dseq, common.HexToAddress(provider), agreedPrice)

	if err != nil {
		return fmt.Errorf("Error while matching Deployment %+v", err)
	}
	return nil
}
