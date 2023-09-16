package cmd

import (
	"errors"
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
	var specifiedTranslators []string

	var transCommand = &cobra.Command{
		Use:   "trans TEXT",
		Short: "Translate text via translation services",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(strings.Join(args, " ")) == 0 {
				return errors.New("missing args: requires TEXT")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			c := conf.Load()
			res := trans.Translate(strings.Join(args, " "), Source, Target, specifiedTranslators, c)
			fmt.Print(trans.Format(res, Format))
		},
	}

	transCommand.Flags().StringVarP(&Source, "source", "", "", "language of the text to be translated")
	transCommand.Flags().StringVarP(&Target, "target", "", "", "the language into which the text should be translated")
	transCommand.Flags().StringVarP(&Format, "format", "f", "plain", "choose an output formatter")
	transCommand.Flags().StringSliceVarP(&specifiedTranslators, "translator", "", []string{}, "use one or more specified translation services")
	return transCommand
}
