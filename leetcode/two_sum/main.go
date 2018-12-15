package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l := len(nums)
	var ret []int

	for i := 0; i < l; i++ {
		t1 := nums[i]

		for j := i + 1; j < l; j++ {
			t2 := nums[j]
			if t1+t2 == target {
				ret = append(ret, i, j)
				return ret
			}
		}
	}
	return ret
}

func main() {
	s := []int{3, 2, 4}
	t := 6
	fmt.Println(twoSum(s, t))
}
