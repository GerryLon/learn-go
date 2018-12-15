package main

import (
	"fmt"
	"time"
)

func main() {
	sendingInterval := time.Second       // 发送间隔
	receptionInterval := time.Second * 2 // 接收间隔

	intChan := make(chan int, 5)

	t1 := time.Now().Unix()

	// sending
	go func() {
		var ts0, ts1 int64
		for i := 0; i < 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()

			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent: %d, interval: %ds\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()

	var ts0, ts1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()

			if ts0 == 0 {
				fmt.Println("Recv:", v)
			} else {
				fmt.Printf("Recv: %d, interval: %ds\n", v, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(receptionInterval)
		}
	}

	// 10s
	fmt.Printf("End, used:%ds\n ", time.Now().Unix()-t1)
}
