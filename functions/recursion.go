package functions

import "fmt"

func Exponential(val int) {
	if val > 100000000 {
		return
	}
	fmt.Println(val)
	Exponential(val * 2)
}
