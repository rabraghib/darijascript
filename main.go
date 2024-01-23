package main

import (
	"fmt"
	"os"

	"github.com/rabraghib/darija-script/src/interpreter"
	"github.com/rabraghib/darija-script/src/lexer"
	"github.com/rabraghib/darija-script/src/parser"
)

func main() {
	// Read source code from a file or other sources
	sourceCode, err := os.ReadFile("./example.ds")
	if err != nil {
		fmt.Println("Error reading source file:", err)
		os.Exit(1)
	}

	// Initialize lexer
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

	// Initialize parser
	p := parser.NewParser(tokens)

	// Parse the source code and build the AST
	program := p.ParseProgram()

	// Initialize interpreter
	eval := interpreter.NewEvaluator()
	_, err = eval.EvaluateProgram(program)
	if err != nil {
		fmt.Println("Error evaluating program:", err)
		os.Exit(1)
	}
}
