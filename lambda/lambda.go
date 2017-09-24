package lambda

import "github.com/chtison/libgo/cli"

// ...
var (
	CmdLambda = cli.NewCommand("lambda")
)

func init() {
	CmdLambda.Usage.Synopsys = "Manage lambda function"
	CmdLambda.Children.Add(cmdCreate, cmdZip)
}