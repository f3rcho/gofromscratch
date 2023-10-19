package exercise

import (
	"fmt"
	"strconv"
)

func StringConverter(s string) (int, string) {
	integer, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	if integer > 100 {
		fmt.Println("Is greater than 100")
	} else {
		fmt.Println("Is lower than 100")
	}

	return integer, s
}
