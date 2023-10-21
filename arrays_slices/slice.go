package arrays_slices

import "fmt"

var tableS []int = []int{2, 3, 4}
var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func ShowSlice() {
	fmt.Println(tableS)

	portion := array[3:]
	portion2 := array[:5]
	portion3 := array[4:6]

	fmt.Println(portion)
	fmt.Println(portion2)
	fmt.Println(portion3)
}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Large %d, Capacity %d", len(elements), cap(elements))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Large %d, Capacity %d", len(nums), cap(nums))
}
