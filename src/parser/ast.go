package parser

import "github.com/rabraghib/darijascript/src/lexer"

type Program struct {
	Statements []Statement
}

type Statement interface {
	statementNode()
}

type Expression interface {
	expressionNode()
}

type statementMarker struct{}

func (s *statementMarker) statementNode() {}

type expressionMarker struct{}

func (e *expressionMarker) expressionNode() {}

type IfStatement struct {
	statementMarker
	Condition           Expression
	IsConditionReversed bool
	Consequence         *BlockStatement
	Alternative         *BlockStatement
}

type WhileStatement struct {
	statementMarker
	Condition   Expression
	Consequence *BlockStatement
}

type BlockStatement struct {
	statementMarker
	Statements []Statement
}

type ExpressionStatement struct {
	statementMarker
	Expression Expression
}

type FunctionDeclaration struct {
	statementMarker
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

type LetStatement struct {
	statementMarker
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	statementMarker
	ReturnValue Expression
}

type ThrowStatement struct {
	statementMarker
	ReturnValue Expression
}

type Identifier struct {
	expressionMarker
	Value string
	Token *lexer.Token
}

type NumberLiteral struct {
	expressionMarker
	Value float64
	Token *lexer.Token
}

type StringLiteral struct {
	expressionMarker
	Value string
	Token *lexer.Token
}

type BooleanLiteral struct {
	expressionMarker
	Value bool
	Token *lexer.Token
}

type PrefixExpression struct {
	expressionMarker
	Operator string
	Right    Expression
	Token    *lexer.Token
}

type AssignmentExpression struct {
	expressionMarker
	Variable Identifier
	Value    Expression
	Token    *lexer.Token
}

type InfixExpression struct {
	expressionMarker
	Left     Expression
	Operator string
	Right    Expression
	Token    *lexer.Token
}

type CallExpression struct {
	expressionMarker
	Function  *Identifier
	Arguments []Expression
	Token     *lexer.Token
}

type ArrayLiteral struct {
	expressionMarker
	Elements []Expression
	Token    *lexer.Token
}

type IndexExpression struct {
	expressionMarker
	Left  Expression
	Index Expression
	Token *lexer.Token
}

type HashLiteral struct {
	expressionMarker
	Pairs map[Expression]Expression
	Token *lexer.Token
}

type Error struct {
	expressionMarker
	Message string
	Pos     lexer.Position
}
