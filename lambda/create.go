package lambda

//go:generate templates/generate.sh

import (
	"io"
	"os"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"

	"github.com/chtison/libaws/lambda/templates"
)

var (
	cmdCreate         = cli.NewCommand("create")
	flagCreateRuntime = cli.NewFlagString('r', "runtime", "python3.6")
	runtimeMap        = map[string]func(string) error{
		"python3.6": runtimePython36,
	}
)

func init() {
	cmdCreate.Usage.Arguments = "PATH"
	cmdCreate.Usage.Synopsys = "Create a new lambda function"
	cmdCreate.Flags.Add(flagCreateRuntime)
	cmdCreate.Function = cmdCreateFunction

	flagCreateRuntime.Usage().Synopsys = "Set the runtime: python3.6"
}

func cmdCreateFunction(cmd *cli.Command, args ...string) error {
	// Check if PATH arg is present
	if len(args) != 1 {
		return cli.Usage(cmd, args...)
	}
	// Check runtime flag
	if _, ok := runtimeMap[flagCreateRuntime.Value]; !ok {
		return fmt.Errorf(`unknown runtime: "%s"`, flagCreateRuntime.Value)
	}
	// Create PATH
	os.Mkdir(args[0], 0700)
	// Change working directory to PATH
	if err := os.Chdir(args[0]); err != nil {
		return err
	}
	// Call runtime function
	if err := runtimeMap[flagCreateRuntime.Value](args[0]); err != nil {
		return err
	}
	return nil
}

func runtimePython36(path string) error {
	// Create file
	if w, err := os.OpenFile("lambda.py", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600); err == nil {
		defer w.Close()
		// Write template to file
		if _, err = io.WriteString(w, templates.Python36+"\n"); err != nil {
			return err
		}
	} else {
		return err
	}
	// Warn the user about created file
	fmt.Printfln(`%s/%s successfully created`, path, "lambda.py")
	return nil
}
