package main

import "fmt"

func generate(numRows int) [][]int {
	ret := make([][]int, 0)
	tmp := make([]int, 0)

	if numRows < 1 {
		return ret
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				l := len(ret[i-1])
				t := 0
				if j == l {
					t = 0
				} else {
					t = ret[i-1][j]
				}
				tmp = append(tmp, ret[i-1][j-1]+t)
			}
		}
		ret = append(ret, tmp)
		tmp = make([]int, 0)
	}

	return ret
}

func main() {
	ret := generate(5)
	fmt.Printf("%v", ret)
}
