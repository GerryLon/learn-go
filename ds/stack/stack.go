package stack

import "container/list"

type Stack struct {
	l       *list.List
	maxSize int
}

func New(maxSize int) *Stack {
	return &Stack{l: list.New(), maxSize: maxSize}
}

func (s *Stack) Empty() bool { return s.Size() == 0 }
func (s *Stack) Full() bool  { return s.Size() == s.maxSize }

func (s *Stack) Push(v interface{}) (ok bool) {
	if s.Full() {
		return false
	}
	s.l.PushBack(v)
	return true
}

func (s *Stack) Pop() (d interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.l.Remove(s.l.Back()), true
}

func (s *Stack) Peek() (d interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.l.Back().Value, true
}

func (s *Stack) Size() int {
	return s.l.Len()
}

func (s *Stack) Destroy() {
	for s.l.Len() > 0 {
		s.l.Remove(s.l.Front())
	}
}
