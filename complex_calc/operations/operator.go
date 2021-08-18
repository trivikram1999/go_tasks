package operations

//PerformOperation is a function that performs operations
func PerformOperation(op1 float32, op2 float32, opt string) float32 {
	var ans float32

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
			panic("Can not divide by Zero")
		}
		ans = op1 / op2

	}

	return ans

}
