package cmd

import (
	"github.com/akash-network/provider/operator"
	"github.com/akash-network/provider/operator/hostname"
	"github.com/akash-network/provider/operator/ip"
	"github.com/akash-network/provider/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:          "provider-services",
		Short:        "Provider services commands",
		SilenceUsage: true,
	}

	cmd.AddCommand(SendManifestCmd())
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

	cmd.AddCommand(operator.OperatorsCmd())
	cmd.AddCommand(operator.ToolsCmd())

	cmd.AddCommand(version.NewVersionCommand())
	cmd.AddCommand(KeysCmd("~/.spheron")) // TODO(spheron): Read this from a file

	return cmd
}
