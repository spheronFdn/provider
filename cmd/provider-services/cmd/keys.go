package cmd

import (
	"fmt"
	"os"

	"github.com/akash-network/provider/spheron"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
)

func KeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Manage keys",
	}

	cmd.AddCommand(
		AddKeyCommand(),
	)

	return cmd
}

func AddKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add new keys",
		Args:  cobra.ExactArgs(1),
		RunE:  runAddCmd,
	}
	cmd.Flags().String(FlagHome, "", "location of home folder")
	cmd.Flags().String(FlagKeySecret, "", "wallet passphrase")

	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	cctx, err := spheron.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	homeDir := cctx.HomeDir
	secret, _ := cmd.Flags().GetString(FlagKeySecret)

	name := args[0]

	walletPath := homeDir + "/" + name + ".json"
	// check if wallet already exists
	_, err = os.Stat(walletPath)
	if err == nil {
		return fmt.Errorf("account already exists")
	}

	// create account if not available
	ks := keystore.NewKeyStore(homeDir, keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.NewAccount(secret)
	if err != nil {
		return err
	}

	err = os.Rename(acc.URL.Path, walletPath)
	if err != nil {
		return err
	}

	cmd.Printf("account created on this path: %s, with address %s", walletPath, acc.Address.Hex())
	return nil
}
