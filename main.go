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
	cmd.Usage.Synopsys = "LibAws is a tool to generate files from templates"
	cmd.Children.Add(
		template.CmdTemplate,
		lambda.CmdLambda,
	)
	if err := cmd.Execute(os.Args[1:]...); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
