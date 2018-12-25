package main

import "fmt"

func forTest() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d ", i)
		for j := 0; j < 10; j++ {
			if j == 0 {
				break
			}
		}
	}
}

func main() {
	forTest()
}
