package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplyTable() {
	var num int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		// fmt.Println(num, "x", i, "=", num*i)
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
