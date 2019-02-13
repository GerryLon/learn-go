package main

import (
	"container/heap"
	"fmt"
	"github.com/GerryLon/learn-go/lang/ds/lru"
	"github.com/GerryLon/learn-go/lang/ds/month_heap"
	"github.com/GerryLon/learn-go/lang/ds/priority_queue"
)

func traverseCache(cache *lru.CacheDB) {
	cache.Traverse(func(k, v interface{}) bool {
		fmt.Printf("%s:%d\t", k, v)
		return true
	})
	fmt.Println()
}

func LRUTest() {
	cache := lru.New(5) // 缓存最大容量: 5

	// 依次存入 a:1, b:2, ..., e: 5
	for i := 0; i < 5; i++ {
		err := cache.Set(string(rune('a'+i)), i+1)
		if err != nil {
			panic(err)
		}
	}

	// 打印结果: e:5	d:4	c:3	b:2	a:1
	// 也就是最后存储的(最近访问的)在最前边
	traverseCache(cache)

	// 打印结果: f:6	e:5	d:4	c:3	b:2
	// 也就是缓存满后, 再放, 会将最后的元素删除
	cache.Set("f", 6)
	traverseCache(cache)

	fmt.Println("test Get()")

	// b:2	c:3	d:4	e:5	f:6
	for i := 0; i < 6; i++ {
		k := string(rune('a' + i))
		v, success, err := cache.Get(k)
		if err != nil {
			panic(err)
		}
		if success {
			fmt.Printf("%s:%d\t", k, v)
		}
	}
	fmt.Println()
}

func MonthHeapTest() {
	h := &month_heap.MonthHeap{"Jan", "Feb", "Mar"}
	heap.Init(h)
	heap.Push(h, "May")
	heap.Push(h, "Apr")
	// first month: Jan
	fmt.Println("first month:", (*h)[0])

	// 输出: Jan	Feb	Mar	Apr	May
	for h.Len() > 0 {
		fmt.Printf("%s\t", heap.Pop(h)) // 注意不是h.Pop()
	}
	fmt.Println()
}

func PriorityQueueTest() {
	items := map[string]int{
		"AA": 5,
		"BB": 8,
		"CC": 3,
	}
	pq := make(priority_queue.PriorityQueue, len(items))
	i := 0

	for value, priority := range items {
		pq[i] = &priority_queue.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := priority_queue.Item{
		Value:    "DD",
		Priority: 1,
	}
	heap.Push(&pq, &item)
	pq.Update(&item, "EE", 99)

	for pq.Len() > 0 {
		x := heap.Pop(&pq).(*priority_queue.Item)
		fmt.Printf("%s %d\n", x.Value, x.Priority)
	}
	fmt.Println()
	return
}

func main() {
	// LRUTest()
	// MonthHeapTest()
	PriorityQueueTest()
}
