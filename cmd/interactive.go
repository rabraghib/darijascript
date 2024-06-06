/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rabraghib/darijascript/src/interpreter"
	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Run the DarijaScript interpreter in interactive mode",
	Long: `Run the DarijaScript interpreter in interactive mode.
You can type DarijaScript code and it will be evaluated immediately.`,
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to the DarijaScript interpreter!\n\n")
		fmt.Printf("Type 'exit' to exit the interpreter.\n\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		sourceCode := ""
		eval := interpreter.NewEvaluator()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(">>> ")
		for scanner.Scan() {
			sourceCode = scanner.Text()
			if sourceCode == "exit" {
				break
			}
			if sourceCode != "" {
				RunCode(sourceCode, eval)
			}
			fmt.Print(">>> ")
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nGoodbye!")
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
