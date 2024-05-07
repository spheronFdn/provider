package spheron

import (
	"bytes"
	"encoding/json"
	"os"
)

func PrintJSON(v interface{}) error {
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
func PrintString(str string) error {
	return PrintBytes([]byte(str))
}

// PrintBytes prints the raw bytes to ctx.Output if it's defined, otherwise to os.Stdout.
// NOTE: for printing a complex state object, you should use ctx.PrintOutput
func PrintBytes(o []byte) error {
	writer := os.Stdout

	_, err := writer.Write(o)
	return err
}
