package priority_queue

import "container/heap"

type Item struct {
	Value    string // 值
	Priority int    // 优先级
	Index    int    // 元素在堆中的索引
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 注意这里是优先级大的在前
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	n := len(*pq)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]

	// 为了安全起见? 应该是防止直接用这个值插入或者别的,
	// 将index设置为一个越界值, 可以防止误操作
	x.Index = -1
	*pq = old[0 : n-1]
	return x
}

// 更新queue中一个Item的Priority和value, 将将Item调整到queue中适当的位置(满足堆的特征)
// update modifies the Priority and Value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority

	// Fix操作比 先Remove再Push要快
	heap.Fix(pq, item.Index)
}
