package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	var name string

	// get name
	for strings.TrimSpace(name) == "" {
		fmt.Print("Please input your name:")

		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("err:%v", err)
			os.Exit(1)
		}

		name = input[:len(input)-1]
		if strings.TrimSpace(name) == "" {
			fmt.Println("name can't be empty!")
			continue
		}
		fmt.Printf("Hello %s, What Can I do for you?\n", name)
		break
	}

	// do help
	for {
		action, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("err:%v", err)
			continue
		}

		action = action[:len(action)-1]
		action = strings.ToLower(action)

		switch action {
		case "":
			continue
		case "bye", "nothing":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I can't help you.")
			continue
		}
	}
}
