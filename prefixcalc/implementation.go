package prefixcalc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func EvaluatePrefixExpression(input string) (float64, error) {
	if input == "" {
		return 0, errors.New("input is empty")
	}

	tokens := strings.Fields(input)
	stack := []float64{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression: insufficient operands for %s", token)
			}
			a, b := stack[0], stack[1]
			stack = stack[2:]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			case "^":
				result = math.Pow(a, b)
			}
			stack = append([]float64{result}, stack...)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack = append([]float64{num}, stack...)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression: remaining stack %v", stack)
	}

	return stack[0], nil
}
