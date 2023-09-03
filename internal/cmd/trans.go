package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/trans"
)

func newTransCommand() *cobra.Command {
	var transCommand = &cobra.Command{
		Use:   "trans TEXT",
		Short: "Translate text",
		Run: func(cmd *cobra.Command, args []string) {
			c := conf.Load()
			res := trans.Translate(strings.Join(args, " "), c)
			fmt.Print(trans.Format(res, "json"))
		},
	}

	return transCommand
}
