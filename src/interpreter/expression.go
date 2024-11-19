package interpreter

import (
	"fmt"
	"strconv"

	"github.com/rabraghib/darijascript/src/parser"
)

func (eval *Evaluator) evaluateExpression(expression parser.Expression) (interface{}, error) {
	if expression == nil {
		return nil, nil
	}
	switch expr := expression.(type) {
	case *parser.Identifier:
		val, ok := eval.env.Get(expr.Value)
		if !ok {
			return nil, fmt.Errorf("identifier not found: %s", expr.Value)
		}
		return val, nil
	case *parser.NumberLiteral:
		return expr.Value, nil
	case *parser.StringLiteral:
		return expr.Value, nil
	case *parser.BooleanLiteral:
		return expr.Value, nil
	case *parser.PrefixExpression:
		rightValue, err := eval.evaluateExpression(expr.Right)
		if err != nil {
			return nil, err
		}
		return eval.evaluatePrefixExpression(expr.Operator, rightValue)
	case *parser.AssignmentExpression:
		value, err := eval.evaluateExpression(expr.Value)
		if err != nil {
			return nil, err
		}
		eval.env.Update(expr.Variable.Value, value)
		return value, nil
	case *parser.InfixExpression:
		leftValue, err := eval.evaluateExpression(expr.Left)
		if err != nil {
			return nil, err
		}
		rightValue, err := eval.evaluateExpression(expr.Right)
		if err != nil {
			return nil, err
		}
		val, err2 := eval.evaluateInfixExpression(expr.Operator, leftValue, rightValue)
		if err2 != nil {
			fmt.Printf("Error in %d:%d\n", expr.Token.Pos.Line, expr.Token.Pos.Column)
		}
		return val, err2
	case *parser.CallExpression:
		return eval.evaluateCallExpression(expr)
	// case *parser.ArrayLiteral:
	// 	return eval.evaluateArrayLiteral(expr)
	// case *parser.IndexExpression:
	// 	return eval.evaluateIndexExpression(expr)
	// case *parser.HashLiteral:
	// 	return eval.evaluateHashLiteral(expr)
	// case *parser.Error:
	// 	return nil, fmt.Errorf(expr.Message)
	default:
		return nil, fmt.Errorf("unsupported expression type: %T", expression)
	}
}

func (eval *Evaluator) evaluatePrefixExpression(operator string, rightValue interface{}) (interface{}, error) {
	switch operator {
	case "!":
		value, err := toBool(rightValue)
		if err != nil {
			return nil, err
		}
		return !value, nil
	case "-":
		value, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return -value, nil
	default:
		return nil, fmt.Errorf("unsupported prefix operator: %s", operator)
	}
}

func (eval *Evaluator) evaluateInfixExpression(operator string, leftValue interface{}, rightValue interface{}) (interface{}, error) {
	switch operator {
	case "+":
		left, err1 := toNumber(leftValue)
		right, err2 := toNumber(rightValue)
		if err1 == nil && err2 == nil {
			return left + right, nil
		}
		leftStr, err1 := convertToString(leftValue)
		rightStr, err2 := convertToString(rightValue)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("unsupported infix operator: %s %s %s", leftValue, operator, rightValue)
		}
		return leftStr + rightStr, nil
	case "-":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left - right, nil
	case "*":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left * right, nil
	case "/":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left / right, nil
	case "%":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		leftInt := int64(left)
		rightInt := int64(right)
		return float64(leftInt % rightInt), nil
	case "<":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left < right, nil
	case "<=":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left <= right, nil
	case ">":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left > right, nil
	case ">=":
		left, err := toNumber(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toNumber(rightValue)
		if err != nil {
			return nil, err
		}
		return left >= right, nil
	case "==":
		return leftValue == rightValue, nil
	case "!=":
		return leftValue != rightValue, nil
	case "&&":
		left, err := toBool(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toBool(rightValue)
		if err != nil {
			return nil, err
		}
		return left && right, nil
	case "||":
		left, err := toBool(leftValue)
		if err != nil {
			return nil, err
		}
		right, err := toBool(rightValue)
		if err != nil {
			return nil, err
		}
		return left || right, nil
	default:
		return nil, fmt.Errorf("unsupported infix operator: %s", operator)
	}
}

func (eval *Evaluator) evaluateCallExpression(callExpression *parser.CallExpression) (interface{}, error) {
	function, ok, err := eval.env.GetFunction(callExpression.Function.Value)
	if err != nil {
		return nil, err
	}
	if !ok {
		return eval.evaluateBuiltinFunctionCall(callExpression)
	}
	if len(function.Parameters) != len(callExpression.Arguments) {
		return nil, fmt.Errorf("wrong number of arguments: expected %d, got %d", len(function.Parameters), len(callExpression.Arguments))
	}

	extraParams := map[string]interface{}{}
	for i, argument := range callExpression.Arguments {
		value, err := eval.evaluateExpression(argument)
		if err != nil {
			return nil, err
		}
		extraParams[function.Parameters[i].Value] = value
	}
	return eval.evaluateBlockStatementExtended(function.Body, extraParams)
}

func toBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("unsupported type for boolean conversion: %T", value)
	}
}

func toNumber(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("unsupported type for number conversion: %T", value)
	}
}
