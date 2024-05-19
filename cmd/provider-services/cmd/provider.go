package cmd

import (
	"context"
	"math/big"

	"github.com/akash-network/provider/spheron"
	"github.com/spf13/cobra"
)

const (
	flagRegion       = "region"
	flagPaymentToken = "payment-token"
	flagValidatorId  = "validator-id"
)

func ProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "Provider registration",
	}

	cmd.AddCommand(
		AddNodeProviderCommand(),
	)

	return cmd
}

func AddNodeProviderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add --region=<region> --from=<wallet_path> --key-secret=<secret> --payment-token=<payment-token>",
		Short: "Command to add a node provider to spheron",
		RunE:  doAddNodeProvider,
	}
	cmd.Flags().StringP(flagRegion, "r", "", "region for node provider")
	cmd.Flags().StringP(FlagFrom, "f", "", "wallet path")
	cmd.Flags().StringP(FlagKeySecret, "k", "", "key secret of wallet")
	cmd.Flags().StringSliceP(flagPaymentToken, "p", []string{}, "payment token for node provider")
	cmd.MarkFlagRequired(flagRegion)
	cmd.MarkFlagRequired(FlagFrom)
	cmd.MarkFlagRequired(FlagKeySecret)
	cmd.MarkFlagRequired(flagPaymentToken)

	return cmd
}

func doAddNodeProvider(cmd *cobra.Command, args []string) error {

	region, err := cmd.Flags().GetString(flagRegion)
	if err != nil {
		return err
	}
	paymentToken, err := cmd.Flags().GetStringSlice(flagPaymentToken)
	if err != nil {
		return err
	}
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	cmd.Printf("region: %s, wallet address: %s, token: %v \n", region, cctx.Key.Address, paymentToken)

	spheronClient := spheron.NewClientWithContext(cctx)
	tx, err := spheronClient.BcClient.AddNodeProvider(context.Background(), region, paymentToken)
	if err != nil {
		return err
	}
	cmd.Printf("Node added with following transaction: %s\n", tx)

	return nil
}

func RemoveNodeProviderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove --validator-id=<validatorId>",
		Short: "Command to remove a node provider to spheron",
		RunE:  doRemoveNodeProvider,
	}
	cmd.Flags().Int64P(flagValidatorId, "v", 0, "validator id of node provider")
	cmd.MarkFlagRequired(flagValidatorId)

	return cmd
}

func doRemoveNodeProvider(cmd *cobra.Command, args []string) error {

	validatorId, err := cmd.Flags().GetInt64(flagValidatorId)
	if err != nil {
		return err
	}

	cmd.Printf("removing validator: %d \n", validatorId)

	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	spheronClient := spheron.NewClientWithContext(cctx)
	tx, err := spheronClient.BcClient.RemoveNodeProvider(context.Background(), big.NewInt(validatorId))
	if err != nil {
		return err
	}
	cmd.Printf("Node removed with following transaction: %s\n", tx)

	return nil
}
