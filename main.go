package main

import (
	"github.com/f3rcho/gofromscratch/variables"
)

func main() {
	variables.ShowIntegers()
	variables.RestVariables()
	status, text := variables.TextConverter(155)

	println(status)
	println(text)
}
