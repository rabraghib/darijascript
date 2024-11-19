package interpreter

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"

	"github.com/rabraghib/darijascript/src/parser"
)

func (eval *Evaluator) evaluateBuiltinFunctionCall(callExpression *parser.CallExpression) (interface{}, error) {
	switch callExpression.Function.Value {
	case "n3ess":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("n3ess() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		num, err := convertToInt64(arg)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(num) * time.Millisecond)
		return nil, nil
	case "abs":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("abs() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		num, err := convertToInt64(arg)
		if err != nil {
			return nil, err
		}
		if num < 0 {
			return -num, nil
		}
		return num, nil
	case "golih":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("golih() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		fmt.Println(arg)
		return nil, nil
	case "dakhel":
		if runtime.GOOS == "js" {
			return nil, fmt.Errorf("dakhel() is not supported in the browser")
		}
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
			return nil, fmt.Errorf(
				"%s() takes exactly 1 argument, %d given",
				callExpression.Function.Value,
				len(callExpression.Arguments),
			)
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
	case "c", "r", "l":
		list := make([]interface{}, len(callExpression.Arguments))
		for i, argument := range callExpression.Arguments {
			value, err := eval.evaluateExpression(argument)
			if err != nil {
				return nil, err
			}
			list[i] = value
		}
		return list, nil
	case "infini":
		// if len(callExpression.Arguments) != 0 {
		// 	return nil, fmt.Errorf("infini() takes exactly 0 arguments, %d given", len(callExpression.Arguments))
		// }
		return math.Inf(1), nil
	case "len":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("len() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		switch v := arg.(type) {
		case []interface{}:
			return len(v), nil
		case string:
			return len(v), nil
		case map[string]interface{}:
			return len(v), nil
		default:
			return nil, fmt.Errorf("unsupported type for len(): %T", arg)
		}
	case "dirRow", "dirCol":
		if len(callExpression.Arguments) != 2 {
			return nil, fmt.Errorf("%s() takes exactly 2 arguments, %d given", callExpression.Function.Value, len(callExpression.Arguments))
		}
		n, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		value, err := eval.evaluateExpression(callExpression.Arguments[1])
		if err != nil {
			return nil, err
		}
		num, err := convertToInt64(n)
		if err != nil {
			return nil, err
		}
		row := make([]interface{}, int(num))
		for i := range row {
			row[i] = value
		}
		return row, nil
	case "toLetter":
		if len(callExpression.Arguments) != 1 {
			return nil, fmt.Errorf("toLetter() takes exactly 1 argument, %d given", len(callExpression.Arguments))
		}
		n, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		num, err := convertToInt64(n)
		if err != nil {
			return nil, err
		}
		if num < 0 || num > 25 {
			return nil, fmt.Errorf("toLetter() argument out of range: %d", int(num))
		}
		return string('A' + int(num)), nil
	case "ara":
		if len(callExpression.Arguments) < 2 {
			return nil, fmt.Errorf("ara() takes at least 2 arguments, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		current := arg
		for i := 1; i < len(callExpression.Arguments); i++ {
			indexArg, err := eval.evaluateExpression(callExpression.Arguments[i])
			if err != nil {
				return nil, err
			}
			idx, err := convertToInt64(indexArg)
			if err != nil {
				return nil, err
			}
			index := int(idx)
			switch v := current.(type) {
			case []interface{}:
				if index < 0 || int(index) >= len(v) {
					return nil, fmt.Errorf("index %d out of range at level %d", index, i)
				}
				current = v[index]

			default:
				return nil, fmt.Errorf("unexpected type at level %d; expected []interface{}", i)
			}
		}

		// `current` now contains the final accessed element
		return current, nil
	case "atih":
		// atih(array, index, value)
		// atih(matrix, row, col, value)
		if len(callExpression.Arguments) < 3 {
			return nil, fmt.Errorf("atih() takes at least 3 arguments, %d given", len(callExpression.Arguments))
		}
		arg, err := eval.evaluateExpression(callExpression.Arguments[0])
		if err != nil {
			return nil, err
		}
		switch v := arg.(type) {
		case []interface{}:
			index, err := eval.evaluateExpression(callExpression.Arguments[1])
			if err != nil {
				return nil, err
			}
			i, err := convertToInt64(index)
			if err != nil {
				return nil, err
			}
			if i < 0 || int(i) >= len(v) {
				return nil, fmt.Errorf("index out of range: %d", int(i))
			}
			value, err := eval.evaluateExpression(callExpression.Arguments[2])
			if err != nil {
				return nil, err
			}
			v[int(i)] = value
			return nil, nil
		// case [][]interface{}:
		// 	if len(callExpression.Arguments) != 4 {
		// 		return nil, fmt.Errorf("atih() takes exactly 4 arguments, %d given", len(callExpression.Arguments))
		// 	}
		// 	row, err := eval.evaluateExpression(callExpression.Arguments[1])
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	col, err := eval.evaluateExpression(callExpression.Arguments[2])
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	i, err := convertToInt64(row)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	j, err := convertToInt64(col)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	if i < 0 || int(i) >= len(v) {
		// 		return nil, fmt.Errorf("row index out of range: %d", int(i))
		// 	}
		// 	if j < 0 || int(j) >= len(v[int(i)]) {
		// 		return nil, fmt.Errorf("column index out of range: %d", int(j))
		// 	}
		// 	value, err := eval.evaluateExpression(callExpression.Arguments[3])
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	v[int(i)][int(j)] = value
		// 	return nil, nil

		default:
			return nil, fmt.Errorf("unsupported type for atih(): %T", arg)
		}
	default:
		return nil, fmt.Errorf("function not found: %s", callExpression.Function.Value)
	}
}

func convertToInt64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("unsupported type for integer conversion: %T", value)
	}
}

func convertToBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case float64:
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
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case string:
		return v, nil
	default:
		return fmt.Sprint(value), nil
		// return "", fmt.Errorf("unsupported type for string conversion: %T", value)
	}
}
