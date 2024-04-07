package parser

import (
	"fmt"

	"github.com/rabraghib/darijascript/src/lexer"
)

type Parser struct {
	errors []Error
	tokens []*lexer.Token
	curPos int
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{tokens: tokens, curPos: 0}
}

func (p *Parser) errorsOrNil() error {
	if len(p.errors) == 0 {
		return nil
	}
	errorMsg := "Parser errors: "
	for _, err := range p.errors {
		errorMsg += fmt.Sprintf("\n\t%s at %d:%d", err.Message, err.Pos.Line, err.Pos.Column)
	}
	return fmt.Errorf(errorMsg)
}

func (p *Parser) ParseProgram() (*Program, error) {
	program := &Program{}
	program.Statements = []Statement{}

	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program, p.errorsOrNil()
}

func (p *Parser) isAtEnd() bool {
	return p.curPos >= len(p.tokens) || p.currentToken().Type == lexer.TT_EOF
}

func (p *Parser) currentToken() *lexer.Token {
	return p.tokens[p.curPos]
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

func (p *Parser) parseIfStatement(isNegative bool) *IfStatement {
	p.consumeToken() // consume TT_IF or TT_IF_NOT token
	ifStatement := &IfStatement{}
	ifStatement.IsConditionReversed = isNegative
	ifStatement.Condition = p.parseExpression()
	ifStatement.Consequence = p.parseBlockStatement()
	if p.matchToken(lexer.TT_ELSE) {
		if p.checkToken(lexer.TT_IF) {
			ifStatement.Alternative = &BlockStatement{
				Statements: []Statement{p.parseIfStatement(false)},
			}
		} else if p.checkToken(lexer.TT_IF_NOT) {
			ifStatement.Alternative = &BlockStatement{
				Statements: []Statement{p.parseIfStatement(true)},
			}
		} else {
			ifStatement.Alternative = p.parseBlockStatement()
		}
	}
	return ifStatement
}

func (p *Parser) parseWhileStatement() *WhileStatement {
	p.consumeToken() // consume TT_WHILE token
	whileStatement := &WhileStatement{}
	whileStatement.Condition = p.parseExpression()
	whileStatement.Consequence = p.parseBlockStatement()
	return whileStatement
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	p.consumeTokenOfType(lexer.TT_LBRACE)
	blockStatement := &BlockStatement{}
	blockStatement.Statements = []Statement{}
	for !p.checkToken(lexer.TT_RBRACE) && !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			blockStatement.Statements = append(blockStatement.Statements, stmt)
		}
	}
	p.consumeTokenOfType(lexer.TT_RBRACE)
	return blockStatement
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	expressionStatement := &ExpressionStatement{
		Expression: p.parseExpression(),
	}
	p.consumeTokenOfType(lexer.TT_SEMICOLON)
	return expressionStatement
}

func (p *Parser) parseFunctionDeclaration() *FunctionDeclaration {
	functionDeclaration := &FunctionDeclaration{}
	p.consumeToken() // consume TT_FUNCTION token
	functionDeclaration.Name = p.parseIdentifier()
	p.consumeTokenOfType(lexer.TT_LPAREN)
	functionDeclaration.Parameters = p.parseFunctionParameters()
	functionDeclaration.Body = p.parseBlockStatement()
	return functionDeclaration
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	p.consumeToken() // consume TT_RETURN token
	returnStatement := &ReturnStatement{}
	if !p.checkToken(lexer.TT_SEMICOLON) {
		returnStatement.ReturnValue = p.parseExpression()
	}
	p.consumeTokenOfType(lexer.TT_SEMICOLON)
	return returnStatement
}

func (p *Parser) parseThrowStatement() Statement {
	p.consumeToken() // consume TT_THROW token
	throwStatement := &ThrowStatement{}
	throwStatement.ReturnValue = p.parseExpression()
	p.consumeTokenOfType(lexer.TT_SEMICOLON)
	return throwStatement
}

func (p *Parser) parseLetStatement() Statement {
	p.consumeToken() // consume TT_LET token
	name := p.parseIdentifier()
	p.consumeTokenOfType(lexer.TT_ASSIGN)
	value := p.parseExpression()
	p.consumeTokenOfType(lexer.TT_SEMICOLON)
	return &LetStatement{
		Name:  name,
		Value: value,
	}
}

func (p *Parser) parseIdentifier() *Identifier {
	token := p.consumeTokenOfType(lexer.TT_IDENTIFIER)
	return &Identifier{
		Value: token.Literal,
		Token: token,
	}
}

func (p *Parser) parseFunctionParameters() []*Identifier {
	parameters := []*Identifier{}
	if p.checkToken(lexer.TT_IDENTIFIER) {
		parameters = append(parameters, p.parseIdentifier())
		for p.matchToken(lexer.TT_COMMA) {
			parameters = append(parameters, p.parseIdentifier())
		}
	}
	p.consumeTokenOfType(lexer.TT_RPAREN)
	return parameters
}

// Utility methods
func (p *Parser) consumeToken() *lexer.Token {
	token := p.currentToken()
	p.curPos++
	return token
}

func (p *Parser) consumeTokenOfType(tokenType lexer.TokenType) *lexer.Token {
	token := p.consumeToken()
	if token.Type != tokenType {
		p.errors = append(p.errors, Error{
			Message: fmt.Sprintf("Expected token of type %s, got %s", tokenType, token.Type),
			Pos:     token.Pos,
		})
		return nil
	}
	return token
}

func (p *Parser) matchToken(tokenType lexer.TokenType) bool {
	if p.checkToken(tokenType) {
		p.consumeToken()
		return true
	}
	return false
}

func (p *Parser) checkToken(tokenType lexer.TokenType) bool {
	return !p.isAtEnd() && p.currentToken().Type == tokenType
}
