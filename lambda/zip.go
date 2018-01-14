package lambda

import (
	"github.com/chtison/libaws/libaws"
	"github.com/chtison/libgo/cli"
)

var (
	cmdZip = cli.NewCommand("zip")
)

func init() {
	cmdZip.Usage.Synopsys = "Zip a lambda function"
	cmdZip.Function = cmdZipFunction
	cmdZip.Flags.Add(libaws.FlagConfig)
}

func cmdZipFunction(cmd *cli.Command, args ...string) error {
	laws, err := libaws.NewLibAwsFromFlagConfig()
	if err != nil {
		return err
	}

	return laws.ZipLambda("")
}
