package parser

import (
	"strconv"

	"github.com/rabraghib/darijascript/src/lexer"
)

func (p *Parser) parseExpression() Expression {
	expressionTokens := []*lexer.Token{}
	for !p.isExpressionEnded() {
		expressionTokens = append(expressionTokens, p.currentToken())
		p.curPos++
	}
	if !p.isAtEnd() && p.currentToken().Type == lexer.TT_SEMICOLON {
		p.curPos++
	}
	tokensTree := p.buildTokensTree(expressionTokens)
	return p.parseTreeExpression(tokensTree)
}

func (p *Parser) parseTreeExpression(tree []*TokenTreeItem) Expression {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		if tree[0].Token.Type == lexer.TT_LPAREN {
			return p.parseTreeExpression(tree[0].children)
		}
		return p.parseSingleExpression(tree[0].Token)
	}

	switch tree[1].Token.Type {
	case lexer.TT_PLUS,
		lexer.TT_MINUS,
		lexer.TT_MULTIPLY,
		lexer.TT_DIVIDE,
		lexer.TT_MODULO,
		lexer.TT_EQUAL,
		lexer.TT_NOT_EQUAL,
		lexer.TT_LESS_THAN,
		lexer.TT_LESS_EQUAL,
		lexer.TT_GREATER_THAN,
		lexer.TT_GREATER_EQUAL:
		if len(tree) == 3 {
			return &InfixExpression{
				Left:     p.parseTreeExpression(tree[:1]),
				Operator: tree[1].Token.Literal,
				Right:    p.parseTreeExpression(tree[2:]),
			}
		}
		newTree := []*TokenTreeItem{
			{
				Token: &lexer.Token{
					Type:    lexer.TT_LPAREN,
					Literal: "(",
				},
				children: []*TokenTreeItem{
					tree[0],
					tree[1],
					tree[2],
				},
			},
		}
		newTree = append(newTree, tree[3:]...)
		return p.parseTreeExpression(newTree)
	}

	switch tree[0].Token.Type {
	case lexer.TT_PLUS, lexer.TT_MINUS, lexer.TT_NOT:
		if len(tree) == 2 {
			return &PrefixExpression{
				Operator: tree[0].Token.Literal,
				Right:    p.parseTreeExpression(tree[1:]),
			}
		}
		newTree := []*TokenTreeItem{
			{
				Token: &lexer.Token{
					Type:    lexer.TT_LPAREN,
					Literal: "(",
				},
				children: []*TokenTreeItem{
					tree[0],
					tree[1],
				},
			},
		}
		newTree = append(newTree, tree[2:]...)
		return p.parseTreeExpression(newTree)

	case lexer.TT_LBRACKET:
		arrayLiteral := &ArrayLiteral{}
		arrayLiteral.Elements = []Expression{}
		for _, child := range tree[0].children {
			arrayLiteral.Elements = append(arrayLiteral.Elements, p.parseTreeExpression([]*TokenTreeItem{child}))
		}
		return arrayLiteral

	case lexer.TT_IDENTIFIER:
		if tree[1].Token.Type == lexer.TT_LPAREN {
			callExpression := &CallExpression{}
			callExpression.Function = Identifier{Value: tree[0].Token.Literal}
			callExpression.Arguments = []Expression{}
			// group tree[1].children by comma
			args := []*TokenTreeItem{}
			start := 0
			for index, child := range tree[1].children {
				if child.Token.Type == lexer.TT_COMMA {
					args = append(args, &TokenTreeItem{children: tree[1].children[start:index]})
					start = index + 1
				}
			}
			args = append(args, &TokenTreeItem{children: tree[1].children[start:]})
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
				callExpression.Arguments = append(callExpression.Arguments, p.parseTreeExpression(children))
			}
			return callExpression
		}
		if tree[1].Token.Type == lexer.TT_ASSIGN {
			return &AssignmentExpression{
				Variable: Identifier{Value: tree[0].Token.Literal},
				Value:    p.parseTreeExpression(tree[2:]),
			}
		}
		if tree[1].Token.Type == lexer.TT_LBRACKET {
			indexExpression := &IndexExpression{}
			indexExpression.Left = p.parseTreeExpression(tree[:1])
			indexExpression.Index = p.parseTreeExpression(tree[2:])
			return indexExpression
		}
	}

	return nil
}

func (p *Parser) parseSingleExpression(token *lexer.Token) Expression {
	if token.Type == lexer.TT_IDENTIFIER {
		return &Identifier{Value: token.Literal}
	}
	if token.Type == lexer.TT_NUMBER {
		num, err := strconv.ParseInt(token.Literal, 10, 64)
		if err != nil {
			p.errors = append(p.errors, err.Error())
			return nil
		}
		return &IntegerLiteral{Value: num}
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

type TokenTreeItem struct {
	Token    *lexer.Token
	children []*TokenTreeItem
}
