package main

import (
	"github.com/f3rcho/gofromscratch/webserver"
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

	// channel1 := make(chan bool)
	// go d.MySlowName("Fernando Cordero", channel1)

	// fmt.Println("Please enter a char")
	// // this is like await til gorouting ends
	// var s string
	// fmt.Scanf(s)

	// defer func() {
	// 	<-channel1
	// 	// <-channel2
	// 	// ...
	// 	// <-channelN
	// }()

	webserver.MyWebServer()
}
