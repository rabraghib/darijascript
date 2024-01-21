package parser

import (
	"strconv"

	"github.com/rabraghib/darija-script/src/lexer"
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
	if tree[0].Token.Type == lexer.TT_LPAREN {
		return p.parseTreeExpression(tree[0].children)
	}
	if len(tree) == 1 {
		return p.parseSingleExpression(tree[0].Token)
	}
	// TODO: This has a bug, it should be fixed
	if tree[0].Token.Type == lexer.TT_NOT {
		return &PrefixExpression{
			Operator: tree[0].Token.Literal,
			Right:    p.parseTreeExpression(tree[1:]),
		}
	}
	if tree[0].Token.Type == lexer.TT_MINUS {
		return &PrefixExpression{
			Operator: tree[0].Token.Literal,
			Right:    p.parseTreeExpression(tree[1:]),
		}
	}
	if tree[0].Token.Type == lexer.TT_PLUS {
		return &PrefixExpression{
			Operator: tree[0].Token.Literal,
			Right:    p.parseTreeExpression(tree[1:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_PLUS {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_MINUS {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_MULTIPLY {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_DIVIDE {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_MODULO {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_EQUAL {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_NOT_EQUAL {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_LESS_THAN {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_LESS_EQUAL {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_GREATER_THAN {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[1].Token.Type == lexer.TT_GREATER_EQUAL {
		return &InfixExpression{
			Left:     p.parseTreeExpression(tree[:1]),
			Operator: tree[1].Token.Literal,
			Right:    p.parseTreeExpression(tree[2:]),
		}
	}

	if tree[0].Token.Type == lexer.TT_IDENTIFIER {
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

	if tree[0].Token.Type == lexer.TT_LBRACKET {
		arrayLiteral := &ArrayLiteral{}
		arrayLiteral.Elements = []Expression{}
		for _, child := range tree[0].children {
			arrayLiteral.Elements = append(arrayLiteral.Elements, p.parseTreeExpression([]*TokenTreeItem{child}))
		}
		return arrayLiteral
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
	var tree []*TokenTreeItem
	for index := 0; index < len(tokens); index++ {
		token := tokens[index]
		if token.Type == lexer.TT_LPAREN || token.Type == lexer.TT_LBRACKET {
			tree = append(tree, &TokenTreeItem{Token: token, children: p.buildTokensTree(tokens[index+1:])})
			index++
			continue
		}
		if token.Type == lexer.TT_RPAREN || token.Type == lexer.TT_LBRACKET {
			break
		}
		tree = append(tree, &TokenTreeItem{Token: token})
	}
	return tree
}

type TokenTreeItem struct {
	Token    *lexer.Token
	children []*TokenTreeItem
}
