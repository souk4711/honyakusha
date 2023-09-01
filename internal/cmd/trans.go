package cmd

import (
	"github.com/spf13/cobra"
)

func newTransCommand() *cobra.Command {
	var transCommand = &cobra.Command{
		Use:   "trans TEXT",
		Short: "Translate text",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return transCommand
}
