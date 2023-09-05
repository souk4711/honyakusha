package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/souk4711/honyakusha/internal/conf"
)

func newInitCommand() *cobra.Command {
	var initCommand = &cobra.Command{
		Use:   "init",
		Short: "Create an configuration in working directory",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal("CreateFailed: " + err.Error())
			}

			path := filepath.Join(wd, "honyakusha.toml")
			if _, err := os.Stat(path); err == nil {
				log.Fatal("CreateFailed: File already existed in " + path)
			}

			if err := os.WriteFile(path, conf.Sample(), 0644); err != nil {
				log.Fatal("CreateFailed: " + err.Error())
			} else {
				fmt.Print("Create a config file in " + path + "\n")
			}
		},
	}

	return initCommand
}
