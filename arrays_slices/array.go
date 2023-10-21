package arrays_slices

import "fmt"

var table [10]int = [10]int{1, 2, 3, 4, 5}

func ShowArray() {
	table[7] = 33
	table[9] = 100

	table2 := [10]string{"Paul", "Dan"}
	fmt.Println(table)
	fmt.Println(table2)
}
