package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func timeAfter(t time.Duration) {
	for {
		select {
		case <-time.After(t):
			fmt.Println("end")
			return
		default:
			fmt.Println("wait")
			time.Sleep(time.Second)
		}
	}
}

func signalTest() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for s := range c {
		fmt.Printf("catch signal: %v\n", s)
		break
	}
	return
}

func main() {
	timeAfter(time.Second * 3)
}
