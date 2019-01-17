package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n&1 != 0 { // len is odd, error
		return false
	}

	if n == 0 {
		return true
	}

	stack := make([]rune, 0)
	var tmpC rune

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c) // push
		} else if c == ')' || c == ']' || c == '}' {
			if len(stack) == 0 {
				return false
			}

			tmpC = stack[len(stack)-1:][0]
			if c == ')' && tmpC == '(' || c == ']' && tmpC == '[' || c == '}' && tmpC == '{' {
				stack = stack[:len(stack)-1]
			}

		} else {
			return false // not () [] {}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Printf("%v\n", isValid("()"))
}
