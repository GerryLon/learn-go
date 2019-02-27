package main

import (
	"fmt"
	"github.com/pkg/profile"
)

// go build main.go
// main.exe
// go tool pprof -pdf main.exe  C:\Users\gerrylon\AppData\Local\Temp\profile564631819\mem.pprof > D:\mem.pdf
func main() {
	// defer profile.Start().Stop() // 1
	defer profile.Start(profile.MemProfile).Stop() // 2
	s := make([]int, 0)
	slowJob(&s)
	fmt.Println(s)
}

func slowJob(slice *[]int) {
	for i := 0; i < 1<<20; i++ {
		*slice = append(*slice, i)
	}
}
