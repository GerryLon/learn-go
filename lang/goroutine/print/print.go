package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	const N = 3
	var wg sync.WaitGroup
	wg.Add(2 * N)

	for i := 'a'; i < 'a'+N; i++ {
		go func(i rune) {
			fmt.Printf("%c ", i)
			wg.Done()
		}(i)
	}

	for i := 'A'; i < 'A'+N; i++ {
		go func(i rune) {
			fmt.Printf("%c ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
}
