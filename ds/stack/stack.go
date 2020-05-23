package stack

import "container/list"

type Stack struct {
	l *list.List
}

func New() *Stack {
	return &Stack{l: list.New()}
}

func (s *Stack)  Empty() bool { return s.Size() == 0 }

func (s *Stack) Push(v interface{}) { s.l.PushBack(v) }

func (s*Stack) Pop() (d interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.l.Remove(s.l.Back()), true
}

func (s *Stack) Peek() (d interface{}, ok bool){
	if s.Empty() {
		return nil, false
	}
	return s.l.Back().Value, true
}

func (s *Stack) Size() int {
	return s.l.Len()
}