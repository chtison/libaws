package template

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"

	"github.com/chtison/libaws/template/templates"
)

var (
	cmdCreate = cli.NewCommand("create")
)

func init() {
	cmdCreate.Usage.Arguments = "PATH | -"
	cmdCreate.Usage.Synopsys = "Create a new libaws template"
	cmdCreate.Function = cmdCreateFunction
}

func cmdCreateFunction(cmd *cli.Command, args ...string) error {
	// Check if PATH arg is present
	if len(args) != 1 {
		return cli.Usage(cmd, args...)
	}
	// Get PATH from STDIN if PATH is "-"
	if args[0] == "-" {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, os.Stdin); err != nil {
			return err
		}
		args[0] = strings.TrimSuffix(buf.String(), "\n")
	}
	// Create PATH
	os.Mkdir(args[0], 0700)
	// Change working directory to PATH
	if err := os.Chdir(args[0]); err != nil {
		return err
	}
	// Create file
	if w, err := os.OpenFile("libaws.yaml", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600); err == nil {
		defer w.Close()
		// Write template to file
		if _, err = io.WriteString(w, templates.LibAwsYaml+"\n"); err != nil {
			return err
		}
	} else {
		return err
	}
	// Warn the user about created file
	fmt.Printfln(`%s/%s successfully created`, args[0], "libaws.yaml")
	return nil
}
