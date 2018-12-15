package channelval

import (
	"github.com/GerryLon/learn-go/util/channel"
	"testing"
)

// fatal error: all goroutines are asleep - deadlock!
func TestIsChannelClosed(t *testing.T) {
	c1 := make(chan interface{})

	c1 <- 1
	if IsChannelClosed(c1) {
		t.Error("IsChannelClosed(c1) should be false, but got true")
	}
	close(c1)

	if !IsChannelClosed(c1) {
		t.Error("IsChannelClosed(c1) should be true, but got false")
	}
}

func TestIsChannelClosed2(t *testing.T) {
	c1 := make(chan interface{})

	if channel.IsClosed(c1) {
		t.Error("channel.IsClosed(c1):should be false, but got true")
	}
	channel.SafeClose(c1)

	if !channel.IsClosed(c1) {
		t.Error("channel.IsClosed(c1):should be true, but got false")
	}

	c2 := make(chan interface{}, 1)
	c2 <- 1
	if channel.IsClosed(c2) {
		t.Error("channel.IsClosed(c2):should be false, but got true")
	}
	channel.SafeClose(c2)
	if !channel.IsClosed(c2) {
		t.Error("channel.IsClosed(c2):should be true, but got false")
	}

	var c3 chan interface{}
	if !channel.IsClosed(c3) {
		t.Error("channel.IsClosed(c3):should be true, but got false")
	}
}

func TestSome(t *testing.T) {
	// var c chan int
	//
	// if c == nil {
	// 	fmt.Println("c is nil")
	// 	c = make(chan int)
	// }
	//
	// if v, ok := <-c; ok {
	// 	fmt.Println(v, ok)
	// } else {
	// 	fmt.Println("false")
	// }

	c1 := make(chan<- int)
	close(c1)

	c2 := make(<-chan int)
	close(c2)
}
