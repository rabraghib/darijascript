package parser

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
	Name       string
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
}

type IntegerLiteral struct {
	expressionMarker
	Value int64
}

type StringLiteral struct {
	expressionMarker
	Value string
}

type BooleanLiteral struct {
	expressionMarker
	Value bool
}

type PrefixExpression struct {
	expressionMarker
	Operator string
	Right    Expression
}

type AssignmentExpression struct {
	expressionMarker
	Variable Identifier
	Value    Expression
}

type InfixExpression struct {
	expressionMarker
	Left     Expression
	Operator string
	Right    Expression
}

type CallExpression struct {
	expressionMarker
	Function  Identifier
	Arguments []Expression
}

type ArrayLiteral struct {
	expressionMarker
	Elements []Expression
}

type IndexExpression struct {
	expressionMarker
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	expressionMarker
	Pairs map[Expression]Expression
}

type Error struct {
	expressionMarker
	Message string
}
