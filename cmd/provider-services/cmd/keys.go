package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
)

func KeysCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys2",
		Short: "Manage keys for spheron-provider",
	}

	cmd.AddCommand(
		AddKeyCommand(),
	)

	cmd.PersistentFlags().String(FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

func AddKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add new keys",
		Args:  cobra.ExactArgs(1),
		RunE:  runAddCmd,
	}
	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	homeDir := cmd.Flag(FlagHome).Value.String()
	name := args[0]

	ks := keystore.NewKeyStore(homeDir, keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.NewAccount(name)
	if err != nil {
		return err
	}

	cmd.Printf("account created on this path: %s, with address %s", homeDir, acc.Address.Hex())
	return nil
}
