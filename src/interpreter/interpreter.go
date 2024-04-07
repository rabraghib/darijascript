package interpreter

import (
	"fmt"
	"os"
	"sort"

	"github.com/rabraghib/darijascript/src/parser"
)

type Evaluator struct {
	env *Environment
}

func NewEvaluator() *Evaluator {
	return &Evaluator{env: NewEnvironment()}
}

func (eval *Evaluator) EvaluateProgram(program *parser.Program) (interface{}, error) {
	return eval.EvaluateStatements(program.Statements)
}

func (eval *Evaluator) EvaluateStatements(statements []parser.Statement) (interface{}, error) {
	sort.SliceStable(statements, func(i, j int) bool {
		switch statements[i].(type) {
		case *parser.FunctionDeclaration:
			_, isFnDeclJ := statements[j].(*parser.FunctionDeclaration)
			return !isFnDeclJ
		default:
			return false
		}
	})
	for _, statement := range statements {
		val, err := eval.evaluateStatement(statement)
		if err != nil {
			return nil, err
		}
		if val != nil {
			return val, nil
		}
	}
	return nil, nil
}

func (eval *Evaluator) evaluateStatement(statement parser.Statement) (interface{}, error) {
	if statement == nil {
		return nil, nil
	}
	switch stmt := statement.(type) {
	case *parser.LetStatement:
		return nil, eval.evaluateLetStatement(stmt)
	case *parser.ExpressionStatement:
		_, err := eval.evaluateExpression(stmt.Expression)
		return nil, err
	case *parser.FunctionDeclaration:
		eval.env.SetFunction(stmt.Name.Value, stmt)
		return nil, nil
	case *parser.IfStatement:
		condition, err := eval.evaluateBoolExpression(stmt.Condition)
		if err != nil {
			return nil, err
		}
		if *condition != stmt.IsConditionReversed {
			return eval.evaluateBlockStatement(stmt.Consequence)
		} else if stmt.Alternative != nil {
			return eval.evaluateBlockStatement(stmt.Alternative)
		}
	case *parser.WhileStatement:
		for {
			condition, err := eval.evaluateBoolExpression(stmt.Condition)
			if err != nil {
				return nil, err
			}
			if !*condition {
				break
			}
			val, err := eval.evaluateBlockStatement(stmt.Consequence)
			if err != nil || val != nil {
				return val, err
			}
		}
	case *parser.BlockStatement:
		return eval.evaluateBlockStatement(stmt)
	case *parser.ReturnStatement:
		return eval.evaluateExpression(stmt.ReturnValue)
	case *parser.ThrowStatement:
		val, err := eval.evaluateExpression(stmt.ReturnValue)
		if val != nil {
			fmt.Println("WAAAAAAA: ", val)
			os.Exit(1)
		}
		return nil, err
	default:
		return nil, fmt.Errorf("unsupported statement type: %T", statement)
	}
	return nil, nil
}

func (eval *Evaluator) evaluateBoolExpression(expression parser.Expression) (*bool, error) {
	rawExprValue, err := eval.evaluateExpression(expression)
	expressionValue, ok := rawExprValue.(bool)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("expression must evaluate to a boolean")
	}
	return &expressionValue, nil
}

func (eval *Evaluator) evaluateLetStatement(stmt *parser.LetStatement) error {
	val, err := eval.evaluateExpression(stmt.Value)
	if err != nil {
		return err
	}
	return eval.env.Set(stmt.Name.Value, val)
}

func (eval *Evaluator) evaluateBlockStatement(block *parser.BlockStatement) (interface{}, error) {
	return eval.evaluateBlockStatementExtended(block, nil)
}

func (eval *Evaluator) evaluateBlockStatementExtended(block *parser.BlockStatement, extraParams map[string]interface{}) (interface{}, error) {
	blockEnv := NewEnclosedEnvironment(eval.env)
	eval.env = blockEnv
	for k, v := range extraParams {
		err := eval.env.Set(k, v)
		if err != nil {
			return nil, err
		}
	}
	val, err := eval.EvaluateStatements(block.Statements)
	eval.env = eval.env.outer
	return val, err
}
