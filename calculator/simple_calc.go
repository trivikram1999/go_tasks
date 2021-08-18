package main

import (
	"fmt"
)

func main() {
	var n_1, n2 int
	var op string
	fmt.Println("Enter the first digit: ")
	fmt.Scanln(&n_1)
	fmt.Println("Enter the second digit: ")
	fmt.Scanln(&n2)
	fmt.Println("Enter the operator(+,-,*,/): ")
	fmt.Scanln(&op)
	result := 0
	switch op {
	case "+":
		result = n_1 + n2
		break
	case "-":
		result = n_1 - n2
		break
	case "*":
		result = n_1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("Cannot be divided with zero")
		} else {
			result = n_1 / n2
		}
	default:
		fmt.Println("Invalid operator")
	}
	fmt.Println("The result is: ")
	fmt.Printf("%d %s %d = %d", n_1, op, n2, result)

}
