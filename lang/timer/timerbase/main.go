package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	fmt.Printf("currentTime: %v\n", time.Now())
	expirationTime := <-t.C
	fmt.Printf("expirationTime:%v\n", expirationTime)
	fmt.Printf("stop: %v\n", t.Stop()) // false, 表示timer已经过期
}
