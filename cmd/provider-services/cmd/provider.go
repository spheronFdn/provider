package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/spf13/cobra"
)

const (
	flagRegion       = "region"
	flagPaymentToken = "payment-token"
	flagValidatorId  = "validator-id"
	flagDomain       = "domain"
	flagAttributes   = "attributes"
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
		Use:   "add --region 'us-east-1' --payment-token 'token1,token2' --domain 'example.com' --attributes 'key1=value1,key2=value2'",
		Short: "Command to add a node provider to spheron",
		RunE:  doAddNodeProvider,
	}
	cmd.Flags().StringP(flagRegion, "r", "", "region for node provider")
	cmd.Flags().StringP(FlagFrom, "f", "", "wallet path")
	cmd.Flags().StringP(FlagKeySecret, "k", "", "key secret of wallet")
	cmd.Flags().StringP(flagDomain, "d", "", "domain that is managed")
	cmd.Flags().StringP(flagAttributes, "a", "", "provider attributes")
	cmd.Flags().StringSliceP(flagPaymentToken, "p", []string{}, "payment token for node provider")

	cmd.MarkFlagRequired(flagRegion)
	cmd.MarkFlagRequired(FlagFrom)
	cmd.MarkFlagRequired(FlagKeySecret)
	cmd.MarkFlagRequired(flagPaymentToken)
	cmd.MarkFlagRequired(flagDomain)

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

	domain, err := cmd.Flags().GetString(flagDomain)
	if err != nil {
		return err
	}

	attributes, err := cmd.Flags().GetString(flagAttributes)
	if err != nil {
		return err
	}

	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	cmd.Printf("region: %s, wallet address: %s, token: %v \n", region, cctx.Key.Address, paymentToken)

	spheronClient := spheron.NewClientWithContext(cctx)

	// Parse attributes into a slice of Attribute structs
	var parsedAttributes []entities.Attribute
	attributePairs := strings.Split(attributes, ",")
	for _, attr := range attributePairs {
		parts := strings.SplitN(attr, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid attribute format: %s", attr)
		}
		parsedAttributes = append(parsedAttributes, entities.Attribute{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	provider := entities.Provider{
		Region:     region,
		Tokens:     paymentToken,
		Attributes: parsedAttributes,
		Domain:     domain,
	}

	tx, err := spheronClient.BcClient.AddNodeProvider(context.Background(), &provider)
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
