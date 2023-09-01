package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal"
)

func newVersionCommand() *cobra.Command {
	var versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Run: func(cmd *cobra.Command, args []string) {
			var version = internal.GetVersion()
			fmt.Print(version.Info())
		},
	}

	return versionCommand
}
