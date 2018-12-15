package channel

func SafeClose(c chan interface{}) {
	if c == nil || IsClosed(c) {
		return
	}
	close(c)
}

// nil channel -> true: means channel can not be closed again
// c := make(chan interface{}) -> false: means channel can be closed
// c := make(chan interface{}, 1} -> false: means channel can be closed
// close(c); -> true: means channel can not be closed again
// c := make(chan interface{}), 刚创建的channel，里边没数据，阻塞的情况会到default中去
func IsClosed(c chan interface{}) bool {
	if c == nil {
		return true
	}
	select {
	case _, ok := <-c:
		if ok {
			return false
		}
		return true
	default:
		return false
	}
}
