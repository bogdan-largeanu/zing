package cmd

import (
	"fmt"
	"github.com/bogdan-largeanu/zing/src/modules"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runBash(command string) {
	//app := "echo hello dork"
	shell := "bash"

	shCommand := exec.Command(shell, "-c", command)
	stdout, err := shCommand.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
func buildCommands(key string, description string, bashBlock string) *cobra.Command {

	bashCommands := strings.Split(bashBlock, "\n")
	return &cobra.Command{
		Use:                    key,
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  description,
		Long:                   "",
		Example:                "",
		ValidArgs:              nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Hidden:                 false,
		Annotations:            nil,
		Version:                "",
		PersistentPreRun:       nil,
		PersistentPreRunE:      nil,
		PreRun:                 nil,
		PreRunE:                nil,
		Run: func(cmd *cobra.Command, args []string) {

			for k, v := range bashCommands {
				if k == 0 {
					fmt.Println("$" + v)
					runBash(v)
				}
				if k == 1 {
					fmt.Println("$" + v)
					runBash(v)
				}

			}
		},
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
		TraverseChildren:           false,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	}

}

func Execute() {
	var rootCmd = &cobra.Command{Use: "default"}

	c, err := modules.ReadYml()

	if err != nil {
		log.Fatal(err)
	}

	for comandIndex := range c {
		rootCmd.AddCommand(
			buildCommands(
				c[comandIndex].Run.Key,
				c[comandIndex].Run.Description,
				c[comandIndex].Run.LiteralBlockBashFile))
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
