package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/sdl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
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

	_, err = spCl.BcClient.CreateOrder(context.TODO(), order)
	if err != nil {
		return fmt.Errorf("Error while creating Deployment transaction")
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
