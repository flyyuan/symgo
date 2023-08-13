package algebra

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Expr interface {
	String() string
	Evaluate() (float64, error)
}

type Symbol string

func (s Symbol) String() string {
	if val, err := strconv.ParseFloat(string(s), 64); err == nil {
		return fmt.Sprintf("%.1f", val)
	}
	return string(s)
}

func (s Symbol) Evaluate() (float64, error) {
	return strconv.ParseFloat(string(s), 64)
}

type Add struct {
	Left  Expr
	Right Expr
}

func (a Add) String() string {
	return fmt.Sprintf("(%s + %s)", a.Left, a.Right)
}

func (a Add) Evaluate() (float64, error) {
	leftVal, err := a.Left.Evaluate()
	if err != nil {
		return 0, err
	}
	rightVal, err := a.Right.Evaluate()
	if err != nil {
		return 0, err
	}
	return leftVal + rightVal, nil
}

type Mul struct {
	Left  Expr
	Right Expr
}

func (m Mul) String() string {
	return fmt.Sprintf("(%s * %s)", m.Left, m.Right)
}

func (m Mul) Evaluate() (float64, error) {
	leftVal, err := m.Left.Evaluate()
	if err != nil {
		return 0, err
	}
	rightVal, err := m.Right.Evaluate()
	if err != nil {
		return 0, err
	}
	return leftVal * rightVal, nil
}

func isOperator(token string) bool {
	return token == "+" || token == "*"
}

func precedence(op string) int {
	switch op {
	case "+":
		return 1
	case "*":
		return 2
	default:
		return 0
	}
}

func infixToPostfix(expr string) ([]string, error) {
	tokens := strings.Fields(expr)
	var stack []string
	var postfix []string

	for _, token := range tokens {
		if unicode.IsDigit(rune(token[0])) {
			postfix = append(postfix, token)
		} else if isOperator(token) {
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil, errors.New("invalid token in expression")
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix, nil
}

func Sympify(expr string) (Expr, error) {
	postfix, err := infixToPostfix(expr)
	if err != nil {
		return nil, err
	}

	var stack []Expr
	for _, token := range postfix {
		if unicode.IsDigit(rune(token[0])) {
			stack = append(stack, Symbol(token))
		} else if token == "+" {
			if len(stack) < 2 {
				return nil, errors.New("invalid expression")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, Add{Left: left, Right: right})
		} else if token == "*" {
			if len(stack) < 2 {
				return nil, errors.New("invalid expression")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, Mul{Left: left, Right: right})
		}
	}

	if len(stack) != 1 {
		return nil, errors.New("invalid expression")
	}

	return stack[0], nil
}
