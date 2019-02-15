package types

// 一个待办项
type TodoItem struct {
	Id    int
	Title string
	Done  bool // 是否已经完成
}

type TodoList []*TodoItem
