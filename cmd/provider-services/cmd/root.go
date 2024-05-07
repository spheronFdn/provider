package cmd

import (
	"github.com/spf13/cobra"

	// init default cosmos-sdk config
	_ "github.com/akash-network/akash-api/go/sdkutil"

	"github.com/akash-network/provider/operator"
	"github.com/akash-network/provider/operator/hostname"
	"github.com/akash-network/provider/operator/ip"
	"github.com/akash-network/provider/version"
)

func NewRootCmd() *cobra.Command {
	// encodingConfig := app.MakeEncodingConfig()

	cmd := &cobra.Command{
		Use:          "provider-services",
		Short:        "Provider services commands",
		SilenceUsage: true,
		// PersistentPreRunE: acmd.GetPersistentPreRunE(encodingConfig, []string{"AP", "AKASH"}),
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

	return cmd
}
