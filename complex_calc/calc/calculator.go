package calc

import (
	operate "complex_calc/operations"
	"fmt"
	"strconv"
	"strings"
)

//Calculate performs calculator operations
func Calculate(exp string) (string, error) {

	var ans float64
	var err error

	res := strings.Fields(exp)
	ops := []float64{}
	opt := []string{}

	for _, v := range res {
		s, err1 := strconv.Atoi(v)
		if err1 == nil {
			ops = append(ops, float64(s))
		} else {
			opt = append(opt, v)
		}
	}
	for _, v := range opt {
		op1 := ops[0]
		op2 := ops[1]

		ans, err = operate.PerformOperation(op1, op2, v)
		if err != nil {
			result := fmt.Sprintf("%v", ans)
			return result, fmt.Errorf("warning: %s", err)
		}
		ops = append([]float64{ans}, ops[2:]...)
	}
	result := fmt.Sprintf("%v", ans)
	return result, err
}
