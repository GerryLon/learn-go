package main

import "fmt"

func reverse(x int) int {

	arr := make([]int, 0)
	sign := 1

	if x < 0 {
		sign = 0
		x = -x
	}

	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	y := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		y = (arr[i] + y) * 10
	}

	if sign == 0 {
		y = -1 * y
	}

	return y / 10
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}
