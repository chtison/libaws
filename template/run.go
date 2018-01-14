package template

//go:generate templates/generate.sh

import (
	"os"

	"github.com/chtison/libgo/cli"

	"github.com/chtison/libaws/libaws"
)

var (
	cmdRun = cli.NewCommand("run")
)

func init() {
	cmdRun.Usage.Synopsys = "Run template engine"
	cmdRun.Function = cmdRunFunction
	cmdRun.Flags.Add(libaws.FlagConfig)
}

func cmdRunFunction(cmd *cli.Command, args ...string) error {
	if len(args) != 0 {
		return cli.Usage(cmd, args...)
	}
	laws, err := libaws.NewLibAwsFromFlagConfig()
	if err != nil {
		return err
	}
	if err = laws.ReadDataFiles(); err != nil {
		return err
	}
	if err = laws.ReadTemplateFiles(); err != nil {
		return err
	}
	if err = laws.Template.ExecuteTemplate(os.Stdout, laws.Root.TemplateToExecute, laws.Data); err != nil {
		return err
	}
	return nil
}
