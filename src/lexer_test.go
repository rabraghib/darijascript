package src

import (
	"testing"

	"github.com/rabraghib/darijascript/src/lexer"
)

func TestLexer(t *testing.T) {
	// Test case 1
	// Test case input
	input := `
		9ayed x = 5;
		9ayed y = 10;
		9ayed foobar = s7i7;
		golih("Somme: " + (x + y));
	`

	// Expected output
	expected := []*lexer.Token{
		{Type: lexer.TT_LET, Literal: "9ayed"},
		{Type: lexer.TT_IDENTIFIER, Literal: "x"},
		{Type: lexer.TT_ASSIGN, Literal: "="},
		{Type: lexer.TT_NUMBER, Literal: "5"},
		{Type: lexer.TT_SEMICOLON, Literal: ";"},
		{Type: lexer.TT_LET, Literal: "9ayed"},
		{Type: lexer.TT_IDENTIFIER, Literal: "y"},
		{Type: lexer.TT_ASSIGN, Literal: "="},
		{Type: lexer.TT_NUMBER, Literal: "10"},
		{Type: lexer.TT_SEMICOLON, Literal: ";"},
		{Type: lexer.TT_LET, Literal: "9ayed"},
		{Type: lexer.TT_IDENTIFIER, Literal: "foobar"},
		{Type: lexer.TT_ASSIGN, Literal: "="},
		{Type: lexer.TT_TRUE, Literal: "s7i7"},
		{Type: lexer.TT_SEMICOLON, Literal: ";"},
		{Type: lexer.TT_IDENTIFIER, Literal: "golih"},
		{Type: lexer.TT_LPAREN, Literal: "("},
		{Type: lexer.TT_STRING, Literal: "Somme: "},
		{Type: lexer.TT_PLUS, Literal: "+"},
		{Type: lexer.TT_LPAREN, Literal: "("},
		{Type: lexer.TT_IDENTIFIER, Literal: "x"},
		{Type: lexer.TT_PLUS, Literal: "+"},
		{Type: lexer.TT_IDENTIFIER, Literal: "y"},
		{Type: lexer.TT_RPAREN, Literal: ")"},
		{Type: lexer.TT_RPAREN, Literal: ")"},
		{Type: lexer.TT_SEMICOLON, Literal: ";"},
	}

	// Run the test
	runLexerTest(t, input, expected)
}

func runLexerTest(t *testing.T, input string, expected []*lexer.Token) {
	// Create a new lexer
	l := lexer.NewLexer(input)

	// Iterate over the expected tokens
	for _, expectedToken := range expected {
		// Get the next token
		token, err := l.NextToken()

		// Check for errors
		if err != nil {
			t.Fatalf("Error lexing source code: %v", err)
		}

		if token.Type != expectedToken.Type {
			t.Fatalf("Expected token type %q, got %q", expectedToken.Type, token.Type)
		}

		if token.Literal != expectedToken.Literal {
			t.Fatalf("Expected token literal %q, got %q", expectedToken.Literal, token.Literal)
		}
	}

	// Check for EOF
	if !l.IsEOL() {
		t.Fatalf("Expected end of file, got more tokens")
	}
}
