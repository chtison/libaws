package template

import (
	"github.com/chtison/libaws/pkg/fileutils"
	"github.com/chtison/libaws/pkg/templates"
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
	if err := fileutils.WriteFile("cloudformation.tmpl.yaml", templates.CloudFormationTmplYaml); err != nil {
		return err
	}
	// if err := fileutils.WriteFile("cloudformation.data.yaml", templates.CloudFormationDataYaml); err != nil {
	// 	return err
	// }
	if err := fileutils.WriteFile("libaws.yaml", templates.LibawsYaml); err != nil {
		return err
	}
	return nil
}
