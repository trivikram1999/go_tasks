package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	o "complex_calc/operations"
)

func main() {
	fmt.Println("Enter The expression")
	consoleReader := bufio.NewReader(os.Stdin)
	exp, _ := consoleReader.ReadString('\n')
	

	var ans float64
	// fmt.Scan(&exp)
	// fmt.Println(exp)

	res := strings.Fields(exp)// 5 + 9 -9..... res =[ 5 +  9 - 9]
	// fmt.Println(res)
	ops := []float64{}
	opt := []string{}

	for _, v := range res {
		s, ok := strconv.Atoi(v)
		if ok == nil {
			ops = append(ops, float64(s)) //ops=[5 9 9]
		} else {
			opt = append(opt, v) //opt= [+ -]
		}

	}
	for _, v := range opt {
		op1 := ops[0] 
		op2 := ops[1]

		ans = o.PerformOperation(op1, op2, v)
		//fmt.Println(ans)
		if ops[2:] == nil {
			break
		}
		ops = append([]float64{ans}, ops[2:]...)
	}
	fmt.Printf("Result is %v", ans)
}
