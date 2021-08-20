package calc

import (
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"strings"

	//"errors"
	operate "complex_calc/operations"
)

//Calculate is a function to calculate
func Calculate(exp string) (string, error) {
	// fmt.Println("Enter The space seperated expression")
	// consoleReader := bufio.NewReader(os.Stdin)
	// exp, _ := consoleReader.ReadString('\n')

	var ans float64
	var err error
	// fmt.Scan(&exp)
	// fmt.Println(exp)

	res := strings.Fields(exp) // 5 + 9 -9..... res =[ 5 +  9 - 9]
	// fmt.Println(res)
	ops := []float64{}
	opt := []string{}

	for _, v := range res {
		s, err1 := strconv.Atoi(v)
		if err1 == nil {
			ops = append(ops, float64(s)) //ops=[5 9 9]
		} else {
			opt = append(opt, v) //opt= [+ -]
		}

	}
	for _, v := range opt {
		op1 := ops[0]
		op2 := ops[1]

		ans, err = operate.PerformOperation(op1, op2, v)
		//fmt.Println(ans)
		if ops[2:] == nil {
			break
		}
		ops = append([]float64{ans}, ops[2:]...)
	}
	result := fmt.Sprintf("%v", ans)
	//fmt.Println(result, err)
	return result, err
}
