package cmd

//go:generate templates/generate.sh

import (
	"io/ioutil"
	"os"

	"github.com/chtison/libaws/pkg/templates"
	"github.com/chtison/libgo/yaml"
	"github.com/spf13/cobra"
)

var (
	CmdRun = &cobra.Command{
		Use:   "run [FLAGS] TEMPLATES",
		Short: "Run template engine",
		Args:  cobra.MinimumNArgs(1),
		RunE:  cmdRunFunction,
		DisableFlagsInUseLine: true,
	}
	CmdRunFlagDatas             []string
	CmdRunFlagTemplateToExecute string
)

func init() {
	CmdRun.PersistentFlags().StringArrayVarP(&CmdRunFlagDatas, "datas", "d", nil, "Paths to yaml data files")
}

func cmdRunFunction(cmd *cobra.Command, args []string) error {
	// Read data
	var data interface{}
	for _, f := range CmdRunFlagDatas {
		if err := yaml.ReadFile(f, &data); err != nil {
			return err
		}
	}
	t := templates.New()
	// Read templates
	for _, f := range args {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		if _, err := t.New(f).Parse(string(b)); err != nil {
			return err
		}
	}
	// Execute template
	if err := t.ExecuteTemplate(os.Stdout, args[0], data); err != nil {
		return err
	}
	return nil
}
