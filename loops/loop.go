package loops

import "fmt"

func Looping() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done")
}

func Looping2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			continue
			//break
		}
	}
	fmt.Println("Done")
}
