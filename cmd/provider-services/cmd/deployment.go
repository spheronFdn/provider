package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/spheron/sdl"
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
		Use:   "close [id]",
		Short: "Close",
		Args:  cobra.ExactArgs(0),
		RunE:  runCloseDeploymentCmd,
	}

	addCmdFlags(cmd)
	return cmd
}

func runCloseDeploymentCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	dseq, err := strconv.ParseUint(args[0], 10, 64)
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
