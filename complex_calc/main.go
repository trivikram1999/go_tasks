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

	var ans float32
	// fmt.Scanln(&exp)

	res := strings.Fields(exp)

	ops := []float32{}
	opt := []string{}

	for _, v := range res {
		s, ok := strconv.Atoi(v)
		if ok == nil {
			ops = append(ops, float32(s))
		} else {
			opt = append(opt, v)
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
		ops = append([]float32{ans}, ops[2:]...)
	}
	fmt.Println("Result is", ans)
}
