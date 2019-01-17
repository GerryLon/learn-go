package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	if n <= 0 {
		return nil
	}
	ret := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}

	return ret
}

func main() {
	fmt.Printf("%v", fizzBuzz(3))
}
