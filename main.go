package main

import (
	"os"

	"github.com/chtison/libaws/pkg/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
