package models

import (
	"errors"
	"fmt"
	"github.com/GerryLon/learn-go/beego/todo/types"
)

type todoListManager struct {
	todoList *types.TodoList
	lastId   int // todoitem按整数递增, lastId是最后一条todoitem的id
}

var (
	m *todoListManager
)

func init() {
	if m == nil {
		m = &todoListManager{
			lastId: -1,
		}
		m.todoList = &types.TodoList{}

		AddTodo(&types.TodoItem{
			Id:    0,
			Title: "test todo item",
			Done:  false,
		})
	}
}

func GetAll() *types.TodoList {
	return m.todoList
}

// 查看给定的totoitem在不在当前列表中
func findTodo(todo *types.TodoItem) (success bool, index int, err error) {
	if todo == nil {
		return false, -1, errors.New("todo item can not be nil")
	}

	if len(*m.todoList) == 0 {
		return false, -1, errors.New("todo list is empty")
	}

	for i, t := range *m.todoList {
		if t.Id == todo.Id {
			return true, i, nil
		}
	}

	return false, -1, fmt.Errorf("todo not found, id=%d", todo.Id)
}

// 添加一个待办项, 不用传递id
func AddTodo(todo *types.TodoItem) (success bool, err error) {
	if todo == nil {
		return false, errors.New("todo item can not be nil")
	}

	// 如果已经有了, 直接修改内容
	b, i, _ := findTodo(todo)
	if b {
		(*m.todoList)[i].Done = false
		(*m.todoList)[i].Title = todo.Title
		return true, nil
	}

	// 新增加的
	m.lastId++
	todo.Id = m.lastId // 从0开始计数
	*m.todoList = append(*m.todoList, todo)

	return true, nil
}

// 删除一个待办项
func DelTodo(id int) (success bool, err error) {
	todo := &types.TodoItem{Id: id}
	b, _, e := findTodo(todo)
	if !b {
		return false, e
	}

	for i, t := range *m.todoList {
		if t.Id == todo.Id {
			*m.todoList = append((*m.todoList)[:i], (*m.todoList)[i+1:]...)
			return true, nil
		}
	}

	return false, errors.New("todo not found")
}
