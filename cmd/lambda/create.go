package lambda

import (
	"io"
	"os"

	"github.com/chtison/libaws/pkg/runtimes"
	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"
)

var (
	cmdCreate         = cli.NewCommand("create")
	flagCreateRuntime = cli.NewFlagString('r', "runtime", "python3.6")
	flagCreateName    = cli.NewFlagString('n', "name", "lambda")
)

func init() {
	cmdCreate.Usage.Arguments = "PATH"
	cmdCreate.Usage.Synopsys = "Create a new lambda function"
	cmdCreate.Flags.Add(flagCreateRuntime, flagCreateName)
	cmdCreate.Function = cmdCreateFunction

	availableRuntimes := make([]string, 0, len(runtimes.Runtimes))
	for key := range runtimes.Runtimes {
		availableRuntimes = append(availableRuntimes, key)
	}
	flagCreateRuntime.Usage().Synopsys = fmt.Sprintf("Available runtimes: %v Default: %s", availableRuntimes, flagCreateRuntime.Value)
	flagCreateName.Usage().Synopsys = fmt.Sprintf("Name of the outputed file. Default: %s", flagCreateName.Value)
}

func cmdCreateFunction(cmd *cli.Command, args ...string) error {
	// Check if PATH arg is present
	if len(args) != 1 {
		return cli.Usage(cmd, args...)
	}
	// Check runtime flag
	t, ok := runtimes.Runtimes[flagCreateRuntime.Value]
	if !ok {
		return fmt.Errorf(`unknown runtime: "%s"`, flagCreateRuntime.Value)
	}
	// Create PATH
	os.Mkdir(args[0], 0700)
	// Change working directory to PATH
	if err := os.Chdir(args[0]); err != nil {
		return err
	}
	// Create file
	if err := createRuntime(args[0], t); err != nil {
		return err
	}
	return nil
}

func createRuntime(path string, t *runtimes.Runtime) error {
	fileName := flagCreateName.Value + "." + t.Extension
	// Create lambda function
	if w, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600); err == nil {
		defer w.Close()
		// Write template to file
		if _, err = io.WriteString(w, t.Template); err != nil {
			return err
		}
	} else {
		return err
	}
	fmt.Printfln(`%s/%s successfully created`, path, fileName)
	return nil
}
