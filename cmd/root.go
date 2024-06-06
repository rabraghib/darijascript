package cmd

import (
	"fmt"
	"os"

	"github.com/rabraghib/darijascript/src/interpreter"
	"github.com/rabraghib/darijascript/src/lexer"
	"github.com/rabraghib/darijascript/src/parser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "darijascript",
	Short: "DarijaScript Interpreter CLI",
	Long: `DarijaScript is a programming language that is based on the Moroccan Arabic language.
This application is a CLI that can be used to run DarijaScript programs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func RunCode(sourceCode string, eval *interpreter.Evaluator) {
	l := lexer.NewLexer(string(sourceCode))
	tokens := []*lexer.Token{}
	for !l.IsEOL() {
		token, err := l.NextToken()
		if err != nil {
			fmt.Println("Error lexing source code:", err)
			os.Exit(1)
		}
		tokens = append(tokens, token)
	}
	p := parser.NewParser(tokens)
	program, err := p.ParseProgram()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	_, err = eval.EvaluateProgram(program)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
