package parser

import "github.com/rabraghib/darija-script/src/lexer"

type Parser struct {
	errors []string
	tokens []*lexer.Token
	curPos int
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{tokens: tokens, curPos: 0}
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (p *Parser) isAtEnd() bool {
	return p.curPos >= len(p.tokens) || p.currentToken().Type == lexer.TT_EOF
}

func (p *Parser) currentToken() *lexer.Token {
	return p.tokens[p.curPos]
}

func (p *Parser) isExpressionEnded() bool {
	if p.isAtEnd() {
		return true
	}
	var currentToken = p.currentToken()
	return currentToken.Type == lexer.TT_SEMICOLON ||
		currentToken.Type == lexer.TT_RBRACE || currentToken.Type == lexer.TT_LBRACE
}

func (p *Parser) parseStatement() Statement {
	currentToken := p.currentToken()
	switch currentToken.Type {
	case lexer.TT_COMMENT:
		p.curPos++ // Skip comment token
		return nil
	case lexer.TT_IF:
		return p.parseIfStatement(false)
	case lexer.TT_IF_NOT:
		return p.parseIfStatement(true)
	case lexer.TT_WHILE:
		return p.parseWhileStatement()
	case lexer.TT_FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.TT_RETURN:
		return p.parseReturnStatement()
	case lexer.TT_THROW:
		return p.parseThrowStatement()
	case lexer.TT_LET:
		return p.parseLetStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseIfStatement(reversed bool) *IfStatement {
	p.curPos++ // Skip if token
	ifStatement := &IfStatement{}
	ifStatement.IsConditionReversed = reversed
	ifStatement.Condition = p.parseExpression()
	ifStatement.Consequence = p.parseBlockStatement()
	if !p.isAtEnd() && p.currentToken().Type == lexer.TT_ELSE {
		p.curPos++ // Skip else token
		ifStatement.Alternative = p.parseBlockStatement()
	}
	return ifStatement
}

func (p *Parser) parseWhileStatement() *WhileStatement {
	p.curPos++ // Skip while token
	whileStatement := &WhileStatement{}
	whileStatement.Condition = p.parseExpression()
	whileStatement.Consequence = p.parseBlockStatement()
	return whileStatement
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	if p.currentToken().Type != lexer.TT_LBRACE {
		p.errors = append(p.errors, "Expected { found "+p.currentToken().Literal+" instead")
		return nil
	}

	blockStatement := &BlockStatement{}
	blockStatement.Statements = []Statement{}

	p.curPos++ // Skip {
	for !p.isAtEnd() && p.currentToken().Type != lexer.TT_RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			blockStatement.Statements = append(blockStatement.Statements, stmt)
		}
	}
	p.curPos++ // Skip }
	return blockStatement
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	expressionStatement := &ExpressionStatement{}
	expressionStatement.Expression = p.parseExpression()
	if expressionStatement.Expression == nil {
		return nil
	}
	return expressionStatement
}

func (p *Parser) parseFunctionDeclaration() *FunctionDeclaration {
	functionDeclaration := &FunctionDeclaration{}
	functionDeclaration.Parameters = []*Identifier{}
	p.curPos++ // Skip function token
	functionDeclaration.Name = p.currentToken().Literal
	p.curPos++ // Skip function name
	if p.currentToken().Type != lexer.TT_LPAREN {
		p.errors = append(p.errors, "Expected ( after function name")
		return nil
	}
	p.curPos++ // Skip (
	for !p.isAtEnd() && p.currentToken().Type != lexer.TT_RPAREN {
		if p.currentToken().Type != lexer.TT_IDENTIFIER {
			p.errors = append(p.errors, "Expected identifier as function parameter")
			return nil
		}
		functionDeclaration.Parameters = append(functionDeclaration.Parameters, &Identifier{
			Value: p.currentToken().Literal,
		})
		p.curPos++ // Skip parameter name
		if p.currentToken().Type == lexer.TT_COMMA {
			p.curPos++ // Skip ,
		} else if p.currentToken().Type != lexer.TT_RPAREN {
			p.errors = append(p.errors, "Expected , or ) after function parameter")
			return nil
		}
	}
	p.curPos++ // Skip )
	functionDeclaration.Body = p.parseBlockStatement()
	return functionDeclaration
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	p.curPos++ // Skip return token
	returnStatement := &ReturnStatement{}
	returnStatement.ReturnValue = p.parseExpression()
	return returnStatement
}

func (p *Parser) parseThrowStatement() *ThrowStatement {
	p.curPos++ // Skip throw token
	throwStatement := &ThrowStatement{}
	throwStatement.ReturnValue = p.parseExpression()
	return throwStatement
}

func (p *Parser) parseLetStatement() *LetStatement {
	p.curPos++ // Skip let token
	if p.isAtEnd() || p.currentToken().Type != lexer.TT_IDENTIFIER {
		p.errors = append(p.errors, "Expected identifier after 9ayed keyword")
		return nil
	}
	letStatement := &LetStatement{}
	letStatement.Name = &Identifier{
		Value: p.currentToken().Literal,
	}
	p.curPos++ // Skip identifier
	if p.isAtEnd() || p.currentToken().Type != lexer.TT_ASSIGN {
		p.errors = append(p.errors, "Expected = after 9ayed statement")
		return nil
	}
	p.curPos++ // Skip =
	letStatement.Value = p.parseExpression()
	return letStatement
}
