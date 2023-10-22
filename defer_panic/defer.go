package defer_panic

import (
	"fmt"
	"log"
)

func WatchDefer() {
	fmt.Println("start")
	defer fmt.Println("Last")
	fmt.Println("End")
}

func ShowPanic() {
	a := 1

	if a == 1 {
		panic("The value of  a is equal to 1")
	}
}

func ShowRecovery() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("A Panic error occured: \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("The value of a is equal to 1")
	}
}
