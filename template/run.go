package template

//go:generate templates/generate.sh

import (
	"errors"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/chtison/libgo/cli"
	yaml "gopkg.in/yaml.v2"

	"github.com/chtison/libaws/template/templates"
)

var (
	cmdRun          = cli.NewCommand("run")
	flagRunOut      = cli.NewFlagString('o', "out", "-")
	flagRunYamlData = cli.NewFlagStringList('y', "yamldata", nil)
)

func init() {
	cmdRun.Usage.Arguments = "[OPTIONS] TEMPLATE [TEMPLATE ...]"
	cmdRun.Usage.Synopsys = "Run template engine"
	cmdRun.Function = cmdRunFunction
	cmdRun.Flags.Add(flagRunOut, flagRunYamlData)

	flagRunOut.Usage().Synopsys = "Specify a name for output file"
	flagRunYamlData.Usage().Synopsys = "Pass data for the template engine"
}

func cmdRunFunction(cmd *cli.Command, args ...string) error {
	if len(args) < 1 {
		return cli.Usage(cmd, args...)
	}
	t := template.Must(templates.Template.Clone())
	if _, err := t.ParseFiles(args...); err != nil {
		return err
	}
	var data map[interface{}]interface{}
	for _, fileName := range flagRunYamlData.Value {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(b, &data); err != nil {
			return errors.New(fileName + ": " + err.Error())
		}
	}
	out := os.Stdout
	if flagRunOut.Value != "-" && flagRunOut.Value != "" {
		fd, err := os.OpenFile(flagRunOut.Value, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0600)
		if err != nil {
			return err
		}
		out = fd
		defer out.Close()
	}
	if err := t.ExecuteTemplate(out, args[0], data); err != nil {
		return err
	}
	return nil
}
