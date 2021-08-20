package operations

import (
	"fmt"
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
			return -999, fmt.Errorf("division by zero cannot be done")
		}
		return op1 / op2, nil
	default:
		return 0, fmt.Errorf("wrong operator used")
	}
}
