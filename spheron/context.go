package spheron

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Context implements a typical context created in SDK modules for transaction
// handling and queries.
type Context struct {
	// Deprecated: Codec codec will be changed to Codec: codec.Codec
	// JSONCodec         codec.JSONCodec
	// Codec             codec.Codec
	// InterfaceRegistry codectypes.InterfaceRegistry
	// Input             io.Reader
	// Keyring        keyring.Keyring
	// KeyringOptions []keyring.Option
	// Output            io.Writer
	// OutputFormat      string
	// Height            int64
	HomeDir string
	Key     *keystore.Key
	// KeyringDir string
	// From              string
	// BroadcastMode     string
	// FromName          string
	// SignModeStr       string
	// UseLedger         bool
	// Simulate          bool
	// GenerateOnly      bool
	// Offline           bool
	// SkipConfirm       bool
	// TxConfig          TxConfig
	// AccountRetriever  AccountRetriever
	// NodeURI           string
	// FeeGranter        sdk.AccAddress
	// Viper             *viper.Viper

	// TODO: Deprecated (remove).
	// LegacyAmino *codec.LegacyAmino
}

// WithHomeDir returns a copy of the Context with HomeDir set.
func (ctx Context) WithHomeDir(dir string) Context {
	if dir != "" {
		ctx.HomeDir = dir
	}
	return ctx
}

func (ctx Context) WithKeyFromFile(walletPath string, password string) Context {
	ctx.Key = ReadKey(walletPath, password)
	return ctx
}

func (ctx Context) PrintJSON(v interface{}) error {
	marshaled, err := json.Marshal(v)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = json.Indent(buf, marshaled, "", "  ")
	if err != nil {
		return err
	}

	// Add a newline, for printing in the terminal
	_, err = buf.WriteRune('\n')
	if err != nil {
		return err
	}

	return PrintString(buf.String())
}

// PrintString prints the raw string to ctx.Output if it's defined, otherwise to os.Stdout
func (ctx Context) PrintString(str string) error {
	return PrintBytes([]byte(str))
}

// PrintBytes prints the raw bytes to ctx.Output if it's defined, otherwise to os.Stdout.
// NOTE: for printing a complex state object, you should use ctx.PrintOutput
func (ctx Context) PrintBytes(o []byte) error {
	writer := os.Stdout

	_, err := writer.Write(o)
	return err
}
