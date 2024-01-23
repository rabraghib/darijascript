/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rabraghib/darija-script/src/interpreter"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [file.ds]",
	Short: "Run a DarijaScript program",
	Long: `Run a DarijaScript program.
Example:
darija-script run my_program.ds`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		sourceCode, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading source file: ", err)
			os.Exit(1)
		}
		eval := interpreter.NewEvaluator()
		runCode(string(sourceCode), eval)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
