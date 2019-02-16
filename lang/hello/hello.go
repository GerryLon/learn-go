package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b uint32) uint32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func demo1() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	input = input[:len(input)-1] // delete '\n'
	fmt.Println("Your name is:", input)
}

func demo2(a interface{}) {
	switch b := a.(type) {
	case string:
		fmt.Println("string", b)
	default:
		fmt.Println("default", b)
	}
}

func deferTest() {
	a := 2

	if a > 3 {
		return
	}
	defer func() {
		fmt.Println("test defer")
	}()
}

func main() {
	fmt.Println(gcd(6, 18))
	demo2("3")
	demo2(3)
	deferTest()
}
