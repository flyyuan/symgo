package symgo

import (
	"fmt"
	"symgo/algebra"
)

func Sympify(expr string) string {
	symbolicExpr, err := algebra.Sympify(expr)
	if err != nil {
		return "Error: invalid expression"
	}
	result, err := symbolicExpr.Evaluate()
	if err != nil {
		return "Error: invalid expression"
	}
	if result == float64(int(result)) {
		return fmt.Sprintf("%d", int(result))
	}
	return fmt.Sprintf("%f", result)
}
