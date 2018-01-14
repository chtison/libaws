package template

import (
	"github.com/chtison/libaws/libaws"
	"github.com/chtison/libaws/template/templates"
	"github.com/chtison/libgo/cli"
)

var (
	cmdCreate = cli.NewCommand("create")
)

func init() {
	cmdCreate.Usage.Synopsys = "Create a new libaws template"
	cmdCreate.Function = cmdCreateFunction
}

func cmdCreateFunction(cmd *cli.Command, args ...string) error {
	if len(args) != 0 {
		return cli.Usage(cmd, args...)
	}
	if err := libaws.WriteFile("cloudformation.tmpl.yaml", templates.CloudFormationTmplYaml); err != nil {
		return err
	}
	if err := libaws.WriteFile("cloudformation.data.yaml", templates.CloudFormationDataYaml); err != nil {
		return err
	}
	if err := libaws.WriteFile("libaws.yaml", templates.LibawsYaml); err != nil {
		return err
	}
	return nil
}
