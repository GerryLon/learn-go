package main

import (
	"sync"
	"time"
)

func unlockEmpty() {
	var mutex sync.Mutex
	mutex.Lock()
	time.Sleep(time.Second)
	mutex.Unlock()
	time.Sleep(time.Second)
	mutex.Unlock()
}

func main() {
	unlockEmpty()
}
