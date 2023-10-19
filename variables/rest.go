package variables

import (
	"fmt"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func RestVariables() {
	Name = "Fernando"
	Status = true
	Salary = 1577.66
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}
