package month_heap

import (
	"strings"
)

var monthMap map[string]int = make(map[string]int)

func init() {
	months := strings.Split("Jan, Feb, Mar, Apr, May, June, July, Aug, Sep, Oct, Nov, Dec", ",")
	for i, v := range months { // 将月份简写下序号对应
		monthMap[strings.TrimSpace(v)] = i + 1
	}
}

type MonthHeap []string

// 实现heap.Interface需要的5个方法
func (m MonthHeap) Len() int {
	return len(m)
}

// 注意这里: 比较的是哪个月份在前
func (m MonthHeap) Less(i, j int) bool {
	return monthMap[m[i]] < monthMap[m[j]]
}

func (m MonthHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Push和Pop都要用指针接收者, 因为要在函数内修改slice
func (m *MonthHeap) Push(x interface{}) {
	*m = append(*m, x.(string))
}

func (m *MonthHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
