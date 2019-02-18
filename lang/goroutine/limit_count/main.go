package main

import (
	"fmt"
	"sync"
	"time"
)

// main1
// func main() {
// 	userCount := 10
// 	ch := make(chan bool, 2)
//
// 	for i := 0; i < userCount; i++ {
// 		ch <- true
// 		go Read(ch, i)
// 	}
// }

// main2
// func main() {
// 	var wg sync.WaitGroup
// 	userCount := 10
// 	for i := 0; i < userCount; i++ {
// 		wg.Add(1)
// 		go Read2(&wg, i)
// 	}
// 	wg.Wait()
// }

// main3
// 可以控制, 两个两个goroutine同时执行
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan bool, 2)
	userCount := 10

	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read3(&wg, ch, i)
	}
	wg.Wait()
}

func Read1(ch chan bool, i int) {
	fmt.Printf("%d go func\n", i)
	<-ch
}

func Read2(wg *sync.WaitGroup, i int) {
	fmt.Printf("%d go func, time %d\n", i, time.Now().Unix())
	wg.Done()
}

func Read3(wg *sync.WaitGroup, ch chan bool, i int) {
	defer wg.Done()
	ch <- true
	fmt.Printf("%d go func, time %d\n", i, time.Now().Unix())
	time.Sleep(time.Second * 1)
	<-ch
}
