package lambda

import (
	"fmt"

	"github.com/chtison/libgo/cli"
)

var (
	cmdZip     = cli.NewCommand("zip")
	flagZipOut = cli.NewFlagString('o', "out", "lambda.zip")
)

func init() {
	cmdZip.Usage.Arguments = "PATH | -"
	cmdZip.Usage.Synopsys = "Zip a lambda function"
	cmdZip.Flags.Add(flagZipOut)
	cmdZip.Function = cmdZipFunction

	flagZipOut.Usage().Synopsys = "Set the name of the outputed zip file"
}

func cmdZipFunction(cmd *cli.Command, args ...string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
