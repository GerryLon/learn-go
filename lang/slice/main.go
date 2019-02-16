package main

import (
	"fmt"
)

func appendValue(slice []int, value int) {
	slice = append(slice, value)
}

func changeValue(slice []int) {
	if len(slice) > 0 {
		slice[0] = slice[0] + 1
	}
}

func main() {
	slice := []int{1, 2, 3}
	appendValue(slice, 4)
	fmt.Println(slice) // [1 2 3]

	changeValue(slice)
	fmt.Println(slice) // [2 2 3]
}
