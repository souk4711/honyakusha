package cmd

import (
	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal/conf"
)

func newTransCommand() *cobra.Command {
	var transCommand = &cobra.Command{
		Use:   "trans TEXT",
		Short: "Translate text",
		Run: func(cmd *cobra.Command, args []string) {
			conf.Load()
		},
	}

	return transCommand
}
