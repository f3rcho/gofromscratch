package arrays_slices

import "fmt"

var table [10]int

func ShowArray() {
	table[7] = 33
	table[9] = 100

	fmt.Println(table)
}
