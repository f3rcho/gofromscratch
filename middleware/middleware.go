package middleware

import "fmt"

func sum(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Init")

	result := middleOperations(sum)(2, 3)
	fmt.Println(result)
	result = middleOperations(substract)(2, 3)
	fmt.Println(result)
	result = middleOperations(multiply)(2, 3)
	fmt.Println(result)
}

func middleOperations(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Executing function")
		return f(a, b)
	}
}
