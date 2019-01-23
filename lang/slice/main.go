package main

import "fmt"

func appendValue(slice []int, value int) {
	slice = append(slice, value)
}

func main() {
	slice := []int{1, 2, 3}
	appendValue(slice, 4)
	fmt.Println(slice) // [1 2 3]
}
