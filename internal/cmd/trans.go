package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/trans"
)

func newTransCommand() *cobra.Command {
	var Source string
	var Target string
	var Format string

	var transCommand = &cobra.Command{
		Use:   "trans TEXT",
		Short: "Translate text",
		Run: func(cmd *cobra.Command, args []string) {
			c := conf.Load()
			res := trans.Translate(strings.Join(args, " "), Source, Target, c)
			fmt.Print(trans.Format(res, Format))
		},
	}

	transCommand.Flags().StringVarP(&Source, "source", "", "", "Language of the text to be translated")
	transCommand.Flags().StringVarP(&Target, "target", "", "", "The language into which the text should be translated")
	transCommand.Flags().StringVarP(&Format, "format", "f", "plain", "Choose an output formatter")
	return transCommand
}
