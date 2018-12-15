package main

import "fmt"

func main() {
	for1()
	fmt.Println("-------------------")
	for2()
}

// 0 - 4
func for2() {
	const N = 5
	for i := 0; i < N; i++ {
		func() {
			fmt.Println(i)
		}()
	}
}

// 5 5 5 5 5
func for1() {
	var funcs []func()
	const N = 5
	for i := 0; i < N; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for i := 0; i < N; i++ {
		funcs[i]()
	}
}
