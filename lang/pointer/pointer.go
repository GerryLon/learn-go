package main

import "fmt"

func returnPointer() *int {
	var a int = 1
	var pa = &a
	return pa
}

func test() {
	address := returnPointer()
	fmt.Printf("%x\n", address)
	address = returnPointer()
	fmt.Printf("%x\n", address)
}

func main() {
	test()
}
