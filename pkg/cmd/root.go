package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Version:               "0.1.0",
		Use:                   filepath.Base(os.Args[0]),
		Short:                 "Libaws is a tool to generate files from templates",
		DisableSuggestions:    true,
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
	}
)

func init() {
	Cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	Cmd.AddCommand(CmdRun)
}
