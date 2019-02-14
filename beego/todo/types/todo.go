package types

// 一个待办项
type TodoItem struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"` // 是否已经完成
}

type TodoList []*TodoItem
