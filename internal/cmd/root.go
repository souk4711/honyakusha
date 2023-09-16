package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal"
)

func Execute(gitCommit string, builtTime string) {
	setDebugMode()
	setHonyakushaVersion(gitCommit, builtTime)

	var command = newHonyakushaCommand()
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func setDebugMode() {
	if os.Getenv("DEBUG") != "" {
		internal.EnableDebugMode()
	}
}

func setHonyakushaVersion(gitCommit string, builtTime string) {
	var version = internal.GetVersion()
	version.GitCommit = gitCommit

	if builtTime, err := strconv.ParseInt(builtTime, 10, 64); err == nil {
		version.BuiltTime = builtTime
	}
}

func newHonyakushaCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:   "honyakusha",
		Short: "Translate text using a variety of translation services",
	}

	rootCommand.SilenceErrors = true
	rootCommand.CompletionOptions.DisableDefaultCmd = true
	rootCommand.AddCommand(newInitCommand())
	rootCommand.AddCommand(newTransCommand())
	rootCommand.AddCommand(newVersionCommand())

	return rootCommand
}
