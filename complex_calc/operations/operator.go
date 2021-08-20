package operations

import (
	"fmt"
	"math"
)

//PerformOperation is a function that returns basic operations
func PerformOperation(op1 float64, op2 float64, opt string) (float64, error) {
	switch opt {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return math.Inf(+1), fmt.Errorf("division by zero cannot be done")
		} else {
			return op1 / op2, nil
		}
	default:
		return 0, fmt.Errorf("wrong operator used")
	}
}
