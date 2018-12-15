package main

import (
	"fmt"
	"time"
)

func main() {
	channelTimeout3()
}

func channelTimeout1() {
	ch := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	t := time.Now().Unix()

	// select会卡在这里,等其中一个case可以执行
	select {
	case v := <-ch:
		fmt.Println("Received: ", v)
	// case <-time.NewTimer(time.Second).C:
	case <-time.After(time.Second): // 一秒后超时
		fmt.Println("Timeout")
		fmt.Printf("%ds\n", time.Now().Unix()-t)
	}
}

func channelTimeout2() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			ch <- i
		}
		close(ch)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer

	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}

		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("End")
				return
			}
			fmt.Println("Received:", v)
		case <-timer.C:
			fmt.Println("timeout")
		}
	}
}

func channelTimeout3() {
	time.AfterFunc(time.Second, func() {
		fmt.Println("hello: time.AfterFunc")
	})
}
