package parser

import (
	"strconv"

	"github.com/rabraghib/darijascript/src/lexer"
)

const HIGHEST_PRECEDENCE = 5

func (p *Parser) isExpressionEnded() bool {
	if p.isAtEnd() {
		return true
	}
	var currentToken = p.currentToken()
	return currentToken.Type == lexer.TT_SEMICOLON || currentToken.Type == lexer.TT_RBRACE || currentToken.Type == lexer.TT_EOF || currentToken.Type == lexer.TT_LBRACE
}

func (p *Parser) parseExpression() Expression {
	tokens := []*lexer.Token{}
	parentheses := 0
	for !p.isExpressionEnded() {
		token := p.currentToken()
		if token.Type == lexer.TT_LPAREN {
			parentheses++
		} else if token.Type == lexer.TT_RPAREN {
			parentheses--
			if parentheses < 0 {
				break
			}
		}
		tokens = append(tokens, p.consumeToken())
	}
	return p.parseWithPrecedence(p.buildTokensTree(tokens), HIGHEST_PRECEDENCE)
}

func (p *Parser) parseWithPrecedence(tree []*TokenTreeItem, precedence int) Expression {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		if tree[0].Token.Type == lexer.TT_LPAREN {
			return p.parseWithPrecedence(tree[0].children, HIGHEST_PRECEDENCE)
		}
		return p.parseSingleExpression(tree[0].Token)
	}
	if len(tree) == 2 {
		if tree[0].Token.Type == lexer.TT_PLUS || tree[0].Token.Type == lexer.TT_MINUS || tree[0].Token.Type == lexer.TT_NOT {
			return &PrefixExpression{
				Operator: tree[0].Token.Literal,
				Right:    p.parseWithPrecedence(tree[1:], precedence),
				Token:    tree[0].Token,
			}
		}
		if tree[0].Token.Type == lexer.TT_IDENTIFIER && tree[1].Token.Type == lexer.TT_LPAREN {
			return p.parseCallExpression(tree[0].Token, tree[1].children)
		}
	}
	if (precedence < 1) || (precedence > HIGHEST_PRECEDENCE) {
		p.errors = append(p.errors, Error{
			Pos:     tree[0].Token.Pos,
			Message: "Invalid expression",
		})
		return nil
	}
	if tree[1].Token.Type == lexer.TT_ASSIGN {
		if len(tree) < 3 || tree[0].Token.Type != lexer.TT_IDENTIFIER {
			p.errors = append(p.errors, Error{
				Pos:     tree[0].Token.Pos,
				Message: "Invalid assignment",
			})
			return nil
		}
		return &AssignmentExpression{
			Variable: Identifier{Value: tree[0].Token.Literal},
			Value:    p.parseWithPrecedence(tree[2:], precedence),
			Token:    tree[1].Token,
		}
	}
	left := []*TokenTreeItem{}
	for i := 0; i < len(tree); i++ {
		curPrecedence := p.getPrecedence(tree[i].Token.Type)
		if curPrecedence == precedence {
			return &InfixExpression{
				Left:     p.parseWithPrecedence(left, precedence-1),
				Operator: tree[i].Token.Literal,
				Right:    p.parseWithPrecedence(tree[i+1:], precedence),
				Token:    tree[i].Token,
			}
		} else {
			left = append(left, tree[i])
		}
	}
	return p.parseWithPrecedence(left, precedence-1)
}

func (p *Parser) getPrecedence(tokenType lexer.TokenType) int {
	switch tokenType {
	case lexer.TT_DIVIDE, lexer.TT_MODULO, lexer.TT_MULTIPLY:
		return 1
	case lexer.TT_PLUS, lexer.TT_MINUS:
		return 2
	case lexer.TT_EQUAL, lexer.TT_NOT_EQUAL, lexer.TT_LESS_THAN, lexer.TT_LESS_EQUAL, lexer.TT_GREATER_THAN, lexer.TT_GREATER_EQUAL:
		return 3
	case lexer.TT_AND:
		return 4
	case lexer.TT_OR:
		return 5
	default:
		return 6
	}
}

func (p *Parser) parseCallExpression(token *lexer.Token, children []*TokenTreeItem) Expression {
	callExpression := &CallExpression{}
	callExpression.Function = &Identifier{Value: token.Literal}
	callExpression.Arguments = []Expression{}
	args := []*TokenTreeItem{}
	start := 0
	for index, child := range children {
		if child.Token.Type == lexer.TT_COMMA {
			args = append(args, &TokenTreeItem{children: children[start:index]})
			start = index + 1
		}
	}
	args = append(args, &TokenTreeItem{children: children[start:]})
	for _, arg := range args {
		children := arg.children
		if len(children) > 0 {
			if children[0].Token.Type == lexer.TT_COMMA {
				children = children[1:]
			}
			if children[len(children)-1].Token.Type == lexer.TT_COMMA {
				children = children[:len(children)-1]
			}
		}
		callExpression.Arguments = append(callExpression.Arguments, p.parseWithPrecedence(children, HIGHEST_PRECEDENCE))
	}
	return callExpression
}

func (p *Parser) parseSingleExpression(token *lexer.Token) Expression {
	if token.Type == lexer.TT_IDENTIFIER {
		return &Identifier{Value: token.Literal}
	}
	if token.Type == lexer.TT_NUMBER {
		num, err := strconv.ParseFloat(token.Literal, 64)
		if err != nil {
			p.errors = append(p.errors, Error{
				Pos:     token.Pos,
				Message: "Invalid number",
			})
			return nil
		}
		return &NumberLiteral{Value: num}
	}
	if token.Type == lexer.TT_STRING {
		return &StringLiteral{Value: token.Literal}
	}
	if token.Type == lexer.TT_TRUE {
		return &BooleanLiteral{Value: true}
	}
	if token.Type == lexer.TT_FALSE {
		return &BooleanLiteral{Value: false}
	}
	return nil
}

type TokenTreeItem struct {
	Token    *lexer.Token
	children []*TokenTreeItem
}

func (p *Parser) buildTokensTree(tokens []*lexer.Token) []*TokenTreeItem {
	tree, _ := p.buildTokensTreeInt(tokens)
	return tree
}

func (p *Parser) buildTokensTreeInt(tokens []*lexer.Token) ([]*TokenTreeItem, int) {
	var tree []*TokenTreeItem
	var index int
	for index = 0; index < len(tokens); index++ {
		token := tokens[index]
		if token.Type == lexer.TT_LPAREN {
			children, intIndex := p.buildTokensTreeInt(tokens[index+1:])
			tree = append(tree, &TokenTreeItem{Token: token, children: children})
			index += intIndex
			continue
		} else if token.Type == lexer.TT_RPAREN {
			index++
			break
		}
		tree = append(tree, &TokenTreeItem{Token: token})
	}
	return tree, index
}
