package lexer

import "fmt"

type TokenType int

const (
	TT_NUMBER TokenType = iota
	TT_STRING
	TT_IDENTIFIER
	TT_IF
	TT_IF_NOT
	TT_ELSE
	TT_COMMENT
	TT_WHILE
	TT_FUNCTION
	TT_TRUE
	TT_FALSE
	TT_ASSIGN
	TT_LET
	TT_RETURN
	TT_THROW

	TT_PLUS
	TT_MINUS
	TT_MULTIPLY
	TT_DIVIDE
	TT_MODULO
	TT_EQUAL
	TT_NOT_EQUAL
	TT_LESS_THAN
	TT_LESS_EQUAL
	TT_GREATER_THAN
	TT_GREATER_EQUAL
	TT_AND
	TT_OR
	TT_NOT

	TT_LPAREN   // (
	TT_RPAREN   // )
	TT_LBRACE   // {
	TT_RBRACE   // }
	TT_LBRACKET // [
	TT_RBRACKET // ]
	TT_COMMA
	TT_COLON
	TT_SEMICOLON
	TT_EOF
)

type Position struct {
	Source string
	Line   int
	Column int
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}

func (t *Token) String() string {
	return fmt.Sprintf(
		"<type:%v, l:c:%v:%v, str:%v>",
		t.Type, t.Pos.Line, t.Pos.Column, t.Literal,
	)
}
