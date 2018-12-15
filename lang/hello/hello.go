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

func main() {
	fmt.Println(gcd(6, 18))
}
