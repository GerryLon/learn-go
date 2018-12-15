package main

import (
	"fmt"
	"sync"
	"time"
)

func doing(wg *sync.WaitGroup, i int) {
	fmt.Printf("start goroutine id %d\n", i)
	time.Sleep(200 * time.Second)
	fmt.Printf("finish goroutine id %d\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	workerCount := 10

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go doing(&wg, i)
	}
	wg.Wait()

}
