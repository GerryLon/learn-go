package types

// 一个待办项
type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"` // 是否已经完成
}

type TodoList []*TodoItem
