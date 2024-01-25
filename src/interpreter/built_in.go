package interpreter

import (
	"fmt"
	"strconv"

	"github.com/rabraghib/darijascript/src/parser"
)

func (eval *Evaluator) evaluateBuiltinFunctionCall(callExpression *parser.CallExpression) (interface{}, error) {
	switch callExpression.Function.Value {
	case "golih":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("golih() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		fmt.Println(arg)
		return fmt.Sprintf("%d", arg), nil
	case "dakhel":
		if len(callExpression.Arguments) > 1 {
			return nil, fmt.Errorf("dakhel() takes at most 1 argument, %d given", len(callExpression.Arguments))
		}
		if len(callExpression.Arguments) == 1 {
			message, err := eval.evaluateExpression(callExpression.Arguments[0])
			if err != nil {
				return nil, err
			}
			fmt.Print(message)
		}
		var input string
		fmt.Scanln(&input)
		return input, nil
	case "rdo3adad", "rdoBooleen", "rdoString":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("rdo3adad() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		if callExpression.Function.Value == "rdo3adad" {
			return convertToInt64(arg)
		}
		if callExpression.Function.Value == "rdoBooleen" {
			return convertToBool(arg)
		}
		return convertToString(arg)
	default:
		return nil, fmt.Errorf("function not found: %s", callExpression.Function.Value)
	}
}

func convertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("unsupported type for integer conversion: %T", value)
	}
}

func convertToBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case int64:
		return v != 0, nil
	case string:
		if v == "s7i7" || v == "S7I7" {
			return true, nil
		}
		if v == "ghalt" || v == "GHALT" {
			return false, nil
		}
		return strconv.ParseBool(v)
	default:
		return false, fmt.Errorf("unsupported type for boolean conversion: %T", value)
	}
}

func convertToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case bool:
		if v {
			return "s7i7", nil
		}
		return "ghalt", nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("unsupported type for string conversion: %T", value)
	}
}
