package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string, channel1 chan bool) {
	chars := strings.Split(name, "")
	for _, char := range chars {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(char)
	}
	channel1 <- true
}
