package functions

import "fmt"

func table(val int) func() int {
	num := val
	sequence := 0

	return func() int {
		sequence++
		return num * sequence
	}
}

func CallClosure() {
	tableof2 := 2
	MyTable := table(tableof2)
	for i := 1; i < 10; i++ {
		fmt.Println(MyTable())
	}
}
