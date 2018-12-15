package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许2个并发同时执行
	taskNum := 10

	for i := 0; i < taskNum; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // 获取信号
			defer func() { <-sem }() // 释放信号

			// do something for task
			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}
