package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i int32 = 0
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			if atomic.CompareAndSwapInt32(&i, i, i+1) {
				fmt.Println("newValue: ", i)
			}
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	i = 2
	time.Sleep(time.Second * 2)
	cancel()
	i = 1
}
