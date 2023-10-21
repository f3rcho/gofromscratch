package functions

import "fmt"

func Calculus(num1 int, num2 int) {
	sum := func(num1 int, num2 int) int {
		return num1 + num2
	}

	subtracts := func(num1 int, num2 int) int {
		return num1 - num2
	}

	multiply := func(num1 int, num2 int) int {
		return num1 * num2
	}

	divide := func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Println(sum(num1, num2))
	fmt.Println(subtracts(num1, num2))
	fmt.Println(multiply(num1, num2))
	fmt.Println(divide(num1, num2))
}
