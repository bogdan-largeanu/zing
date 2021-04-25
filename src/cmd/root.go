package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
	"zing/src/modules"
)

func executeShellCommand(command string, path string, extraArgs []string) {
	shell := "bash"
	var shCommand *exec.Cmd

	if path == "" {
		shCommand = exec.Command(shell, "-ci", command)

	} else {
		shCommand = exec.Command(shell, "-ci", command)
		//TODO add the path from yaml to run all the lines of code from block
		//shCommand.Path = "~/Documents"
	}

	stdout, err := shCommand.Output()

	if err != nil {
		fmt.Println("ERR: " + err.Error())
	}

	fmt.Print(string(stdout))
}

func buildCommands(key string, path string, bashBlock string, description string) *cobra.Command {

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
		Run: func(cmd *cobra.Command, extraArguments []string) {
			runShellCommands(bashCommands, path, extraArguments)
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

func runShellCommands(bashCommands []string, path string, extraArguments []string) {
	for k, bashLine := range bashCommands {
		//first shell line accepts extra arguments , otherwise will execute only the code described
		if k == 0 {
			bashLine = bashLine + " " + strings.Join(extraArguments, " ")
			println("DEBUG: src$ " + bashLine)
			executeShellCommand(bashLine, path, extraArguments)
		} else {
			println("DEBUG: src$ " + bashLine)
			executeShellCommand(bashLine, path, extraArguments)
		}
	}
}

func Execute() {
	var rootCmd = &cobra.Command{}

	userCommands, err := modules.ReadYml()
	if err != nil {
		log.Fatal(err)
	}

	for commandIndex := range userCommands {
		rootCmd.AddCommand(
			buildCommands(userCommands[commandIndex].Run.Key,
				userCommands[commandIndex].Run.Path,
				userCommands[commandIndex].Run.LiteralBlockBashFile,
				userCommands[commandIndex].Run.Description))
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
