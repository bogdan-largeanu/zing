package main

import (
	"fmt"
	"github.com/bogdan-largeanu/zing/src/cmd"
	"os/exec"
)

func Sum(x int, y int) int {
	return x + y
}

func runBash(command string) {
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
	cmd.Execute()
}

func Main() {
	main()
}
