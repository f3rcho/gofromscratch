package keyword

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngressNumbers() {
	var num1 int
	var num2 int
	var leyend string
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1 : ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2 : ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyend = scanner.Text()
	}

	fmt.Println(leyend, num1*num2)
}
