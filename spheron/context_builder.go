package spheron

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	FlagHome      = "home"
	FlagFrom      = "from"
	FlagKeySecret = "key-secret"
)
const ClientContextKey = "client.context"

func ReadCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {

	if clientCtx.HomeDir == "" || flagSet.Changed(FlagHome) {
		homeDir, _ := flagSet.GetString(FlagHome)
		clientCtx = clientCtx.WithHomeDir(homeDir)
	}

	if clientCtx.Key == nil || flagSet.Changed(FlagFrom) {
		keyPath, _ := flagSet.GetString(FlagFrom)
		secret, _ := flagSet.GetString(FlagKeySecret)
		clientCtx = clientCtx.WithKeyFromFile(keyPath, secret)
	}

	return clientCtx, nil
}

func readTxCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	clientCtx, err := ReadCommandFlags(clientCtx, flagSet)
	if err != nil {
		return clientCtx, err
	}

	return clientCtx, nil
}

func readQueryCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	// TODO (spheron) extend some query specific flags if needed
	// if clientCtx.Height == 0 || flagSet.Changed(flags.FlagHeight) {
	// 	height, _ := flagSet.GetInt64(flags.FlagHeight)
	// 	clientCtx = clientCtx.WithHeight(height)
	// }
	return ReadCommandFlags(clientCtx, flagSet)
}

func GetClientTxContext(cmd *cobra.Command) (Context, error) {
	ctx := GetClientContextFromCmd(cmd)
	return readTxCommandFlags(ctx, cmd.Flags())
}

func GetClientContextFromCmd(cmd *cobra.Command) Context {
	if v := cmd.Context().Value(ClientContextKey); v != nil {
		clientCtxPtr := v.(*Context)
		return *clientCtxPtr
	}

	return Context{}
}

func SetCmdClientContext(cmd *cobra.Command, clientCtx Context) {
	v := cmd.Context().Value(ClientContextKey)
	if v == nil {
		cmd.SetContext(context.WithValue(cmd.Context(), ClientContextKey, &clientCtx))
		return
	}
	clientCtxPtr := v.(*Context)
	*clientCtxPtr = clientCtx
}
