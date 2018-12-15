package channelval

import (
	"fmt"
	"github.com/GerryLon/learn-go/util/channel"
	"sync"
	"time"
)

func main() {
	// changeVal2()

	SafeCloseDemo()
}

// 对值类型的值来说，接收方对值的修改不会影响发送方
// 对于引用类型的值来说，接收方可以修改发送方发的值
// countMap: map[count:1].
// countMap: map[count:2].
// countMap: map[count:3].
// countMap: map[count:4].
// countMap: map[count:5].
// stopped: [receiver]
func changeVal1() {
	var mapChan = make(chan map[string]int, 1)
	syncChan := make(chan struct{}, 2)

	// receiver
	go func() {
		for {
			if m, ok := <-mapChan; ok {
				m["count"]++
			} else {
				break
			}
		}
		fmt.Println("stopped: [receiver]")
		syncChan <- struct{}{}
	}()

	// sender
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("countMap: %v.\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

type Counter struct {
	count int
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d", c.count)
}

// also ok
// func (c Counter) String() string {
// 	return fmt.Sprintf("%d", c.count)
// }

func changeVal2() {
	// var mapChan = make(chan map[string]Counter, 1)
	var mapChan = make(chan map[string]*Counter, 1)
	syncChan := make(chan bool, 2)

	go func() {
		for {
			if m, ok := <-mapChan; ok {
				counter := m["count"]
				counter.count++
				// m["count"].count++
			} else {
				break
			}
		}
		fmt.Println("stopped: [receiver]")
		syncChan <- true
	}()

	go func() {
		// countMap := map[string]Counter{
		// 	"count": Counter{},
		// }
		countMap := map[string]*Counter{
			"count": &Counter{},
		}

		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("countMap: %v.\n", countMap)
		}
		close(mapChan)
		syncChan <- true
	}()

	<-syncChan
	<-syncChan
}

// BUG!
func IsChannelClosed(c chan interface{}) bool {
	// c = chan interface{}(c) // 如果c是“send only"的，要转换
	var ok bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, ok = <-c
		wg.Done()
	}()
	wg.Wait()
	return ok
}

func SafeCloseDemo() {
	// panic: close of nil channel
	var c chan interface{}
	// close(c)
	channel.SafeClose(c)
}
