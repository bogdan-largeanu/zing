package main

import (
	"bytes"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return c, buf.String(), err
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func emptyRun(*cobra.Command, []string) {}

func TestSingleCommand(t *testing.T) {
	var rootCmdArgs []string
	rootCmd := &cobra.Command{
		Use:  "root",
		Args: cobra.ExactArgs(2),
		Run:  func(_ *cobra.Command, args []string) { rootCmdArgs = args },
	}
	aCmd := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: emptyRun}
	bCmd := &cobra.Command{Use: "b", Args: cobra.NoArgs, Run: emptyRun}
	rootCmd.AddCommand(aCmd, bCmd)
	output, err := executeCommand(rootCmd, "one", "two")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	got := strings.Join(rootCmdArgs, " ")
	expected := "one two"
	if got != expected {
		t.Errorf("rootCmdArgs expected: %q, got: %q", expected, got)
	}
}
