package main

import "fmt"

func channelConv1() {
	var ok bool
	ch := make(chan int, 1)

	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int, ", ok)

	s := "BrainWu"
	// if v, ok := s.(string); ok { // Invalid type assertion: s.(string) (non-interface type string on left)
	if v, ok := interface{}(s).(string); ok {
		fmt.Println(v)
	}
}

func main() {
	channelConv1()
}
