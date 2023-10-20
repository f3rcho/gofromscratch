package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplyTable() string {
	var num int
	var textTable string
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
		textTable += fmt.Sprintf("%d x %d = %d\n", num, i, num*i)
	}
	return textTable
}
