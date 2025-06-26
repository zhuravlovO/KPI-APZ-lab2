package lab2

import (
	"fmt"
	"strings"
)

// PostfixToLisp перетворює постфіксний вираз у Lisp-подібний.
func PostfixToLisp(input string) (string, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", fmt.Errorf("input string is empty")
	}

	stack := []string{}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands for operator %s", token)
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var newExpr string
			if token == "^" {
				newExpr = fmt.Sprintf("(pow %s %s)", operand1, operand2)
			} else {
				newExpr = fmt.Sprintf("(%s %s %s)", token, operand1, operand2)
			}

			stack = append(stack, newExpr)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: incorrect number of operators or operands")
	}

	return stack[0], nil
}

// isOperator - допоміжна функція для перевірки, чи є токен оператором.
func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	default:
		return false
	}
}
