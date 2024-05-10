package cmd

import (
	"github.com/akash-network/provider/operator"
	"github.com/akash-network/provider/operator/hostname"
	"github.com/akash-network/provider/operator/ip"
	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:               "provider-services",
		Short:             "Provider services commands",
		SilenceUsage:      true,
		PersistentPreRunE: GetInitContextPreRunE(),
	}

	// main commands
	cmd.AddCommand(KeysCmd())
	cmd.AddCommand(DeploymentCmd())
	cmd.AddCommand(SendManifestCmd())

	// not main commands
	cmd.AddCommand(statusCmd())
	cmd.AddCommand(leaseStatusCmd())
	cmd.AddCommand(leaseEventsCmd())
	cmd.AddCommand(leaseLogsCmd())
	cmd.AddCommand(serviceStatusCmd())
	cmd.AddCommand(RunCmd())
	cmd.AddCommand(LeaseShellCmd())
	cmd.AddCommand(hostname.Cmd())
	cmd.AddCommand(ip.Cmd())
	cmd.AddCommand(clusterNSCmd())
	cmd.AddCommand(SDL2ManifestCmd())
	cmd.AddCommand(MigrateHostnamesCmd())
	cmd.AddCommand(MigrateEndpointsCmd())
	cmd.AddCommand(ProviderCmd())

	cmd.AddCommand(operator.OperatorsCmd())
	cmd.AddCommand(operator.ToolsCmd())

	cmd.AddCommand(version.NewVersionCommand())
	return cmd
}

func GetInitContextPreRunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		initClientCtx, err :=
			spheron.ReadCommandFlags(spheron.Context{}, cmd.Flags())
		if err != nil {
			return err
		}
		spheron.SetCmdClientContext(cmd, initClientCtx)
		return nil
	}
}
