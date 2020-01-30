package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func Execute() {
	tasks := make([]string, 10)

	var cmdRunRequest = &cobra.Command{
		Use:   "add [string to add task]",
		Short: "add task to a listPrint anything to the screen",
		Long: `add will add tasks to a list.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Add: " + strings.Join(args, " "))
			//fmt.Println("task added")
			tasks = append(tasks, strings.Join(args, "-"))
			fmt.Println("Added to list: ", tasks)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdRunRequest)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
