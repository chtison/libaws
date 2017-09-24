package main

import (
	"os"
	"path/filepath"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"

	"github.com/chtison/libaws/lambda"
	"github.com/chtison/libaws/template"
)

func main() {
	cmd := cli.NewCommand(filepath.Base(os.Args[0]))
	cmd.Usage.Synopsys = "A tool for generating AWS cloudformation files from templates"
	cmd.Function = entrypoint
	cmd.Children.Add(lambda.CmdLambda, template.CmdTemplate)
	if err := cmd.Execute(os.Args[1:]...); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func entrypoint(cmd *cli.Command, args ...string) error {
	return cli.Usage(cmd, args...)
}
