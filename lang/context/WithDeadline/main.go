package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func someHandler() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	go doStuff(ctx)

	// 10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

// 每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			logger.Printf("deadline set")
			if time.Now().After(deadline) {
				log.Printf(ctx.Err().Error())
				return
			}
		}

		select {
		case <-ctx.Done():
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	logger.Printf("start")
	someHandler()
	logger.Printf("end")
}
