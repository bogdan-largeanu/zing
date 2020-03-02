package main

import (
	"fmt"
	"github.com/bogdan-largeanu/zing/src/cmd"
	"github.com/bogdan-largeanu/zing/src/modules"
	"os/exec"
	"strings"
)

func Sum(x int, y int) int {
	return x + y
}

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

func main() {
	var c modules.YamlStructure
	c.ReadYml()
	//fmt.Printf("%+v\n", c)
	literalBlockBashFile := c[0].Run.LiteralBlockBashFile
	//fmt.Println("Wasabi: "+ literalBlockBashFile)
	_ = strings.Split(literalBlockBashFile, "\n")

	//for i, command := range arrayOfBashCommands {
	//	println(i, command)
	//}

	//modules.WriteYml()
	//runBash(arrayOfBashCommands[0])

	//	rootCmd.Execute()
	cmd.Execute()
}

func Main() {
	main()
}
