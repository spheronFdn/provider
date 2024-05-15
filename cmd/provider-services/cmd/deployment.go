package cmd

import (
	"context"
	"fmt"

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
