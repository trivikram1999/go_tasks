package operations
import("math")

//PerformOperation is a function that performs operations
func PerformOperation(op1 float64, op2 float64, opt string) float64 {
	var ans float64

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
			return(math.Inf(1))
			// panic("Can not divide by Zero")
			
		}
		ans = op1 / op2

	}

	return ans

}
