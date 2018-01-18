package template

import "github.com/chtison/libgo/cli"

// CmdTemplate ...
var CmdTemplate = cli.NewCommand("template")

func init() {
	CmdTemplate.Usage.Synopsys = "Manage cloudformation template"
	CmdTemplate.Children.Add(cmdCreate, cmdRun)
}
