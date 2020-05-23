package queue

import "container/list"

type Queue struct {
	l       *list.List
	maxSize int
}

func New(maxSize int) *Queue {
	return &Queue{l: list.New(), maxSize: maxSize}
}

func (q *Queue) Empty() bool { return q.l.Len() == 0 }
func (q *Queue) Full() bool  { return q.l.Len() == q.maxSize }

func (q *Queue) Enqueue(v interface{}) (ok bool) {
	if q.Full() {
		return false
	}
	q.l.PushBack(v)
	return true
}

func (q *Queue) Dequeue() (d interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}

	return q.l.Remove(q.l.Front()), true
}

func (q *Queue) Size() int { return q.l.Len() }

// peek the queue head
func (q *Queue) Peek() (d interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}
	return q.l.Front().Value, true
}

func (q *Queue) Destroy() {
	for q.l.Len() > 0 {
		q.l.Remove(q.l.Front())
	}
}
