package main

import (
	"github.com/f3rcho/gofromscratch/files"
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
	files.ReadArchive()

}
