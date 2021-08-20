package operations

import (
	"errors"
	"math"
)

//PerformOperation is a function that performs operations
func PerformOperation(op1 float64, op2 float64, opt string) (float64, error) {
	var ans float64
	var err error
	switch opt {
	case "+":

		ans = op1 + op2
		//fmt.Println(ans)
	case "-":
		ans = op1 - op2
	case "*":
		ans = op1 * op2
	case "/":
		if op2 == 0 {
			err = errors.New("division by zero cannot be done")
			ans = math.Inf(+1)
			// panic("Can not divide by Zero")
		} else {
			ans = op1 / op2
		}

	}
	//fmt.Println(ans, err)
	return ans, err

}
