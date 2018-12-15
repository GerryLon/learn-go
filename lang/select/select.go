package main

import (
	"fmt"
	"time"
)

func select1() {
	const N = 5
	ch := make(chan int, N)

	startTime := time.Now()
	for i := 0; i < N; i++ {
		select {
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		}
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", <-ch)
	}
	fmt.Println()
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}

func main() {
	select1()
}
