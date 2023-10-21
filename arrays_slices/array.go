package arrays_slices

import "fmt"

var table [10]int = [10]int{1, 2, 3, 4, 5}
var matrix [8][8]int

func ShowArray() {
	table[7] = 33
	table[9] = 100

	table2 := [10]string{"Paul", "Dan"}
	fmt.Println(table)
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i], "<<")
	}

	matrix[7][2] = 15
	fmt.Println(matrix)
}
