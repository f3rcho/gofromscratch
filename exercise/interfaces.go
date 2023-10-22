package exercise

import (
	"fmt"

	"github.com/f3rcho/gofromscratch/interfaces"
)

func HumanBreath(h interfaces.Human) {
	h.Breath()
	fmt.Printf("I'm a %s, and I'm breathing\n", h.Sex())
}
