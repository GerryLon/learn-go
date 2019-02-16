package main

import (
	"fmt"
)

var v4redeclare string = "1,2,3"

func redeclare() {
	v := []int{4, 5, 7}
	if v != nil {
		var v = 789
		fmt.Println("v=", v)
	}
}

func main() {
	redeclare()
}
