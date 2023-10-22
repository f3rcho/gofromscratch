package main

import (
	"fmt"

	d "github.com/f3rcho/gofromscratch/goroutines"
)

func main() {

	/*
			os := runtime.GOOS

			if os == "linux" {
				println("Linux")
			} else {
				fmt.Println("Ahother SO")
			}

			if os := runtime.GOOS; os == "darwin" {
				println("darwin")
			} else {
				println("Ahother SO")
			}

			switch os := runtime.GOOS; os {
			case "linux":
				fmt.Println("This is linux")
			case "windows":
				fmt.Println("This is windows")
			default:
				fmt.Printf("%s \n", os)
			}
			num, s := exercise.StringConverter("101")

			fmt.Println(num)
			fmt.Println(s)

		keyword.IngressNumbers()
	*/

	// fmt.Println(exercise.MultiplyTable())
	// files.CreateArchive()
	// files.ReadArchive()

	// maps.ShowMaps()

	// users.HighUser()

	// Pedro := new(models.Man)
	// e.HumanBreath(Pedro)

	// Maria := new(models.Woman)
	// e.HumanBreath(Maria)

	// d.WatchDefer()
	// d.ShowPanic()
	// d.ShowRecovery()

	go d.MySlowName("Fernando Cordero")

	fmt.Println("Please enter a char")
	var x string
	fmt.Scanln(&x)
}
