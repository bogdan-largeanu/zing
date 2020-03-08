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

//type yamlStructure []struct {
//	Run struct {
//		Key                  string `yaml:"key"`
//		Description          string `yaml:"description"`
//		Path                 string `yaml:"path"`
//		LiteralBlockBashFile string `yaml:"literal_block_bash_file"`
//	} `yaml:"run"`
//}
//
//func (c *yamlStructure) ReadYml() *yamlStructure {
//
//	yamlFile, err := ioutil.ReadFile("conf.yaml")
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//	err = yaml.Unmarshal(yamlFile, &c)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//
//	return c
//}

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

	c, _, err := modules.ReadYml()

	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(buildCommands(c[0].Run.Key, c[0].Run.Description, c[0].Run.LiteralBlockBashFile))
	rootCmd.AddCommand(buildCommands(c[1].Run.Key, c[1].Run.Description, c[1].Run.LiteralBlockBashFile))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
