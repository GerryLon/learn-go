package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	time.Sleep(3 * time.Second)
	cancel()

}

// 每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
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
	time.Sleep(time.Second)
	logger.Printf("end")
}
