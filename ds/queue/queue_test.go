package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	const MaxSize = 10

	q := New(MaxSize)
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.Empty())
	defer q.Dequeue()

	for i := 0; i < MaxSize; i++ {
		q.Enqueue(i)
		assert.Equal(t, i+1, q.Size())
	}

	for i := 0; i < MaxSize; i++ {
		d, ok := q.Dequeue()
		assert.True(t, ok)
		assert.Equal(t, i, d.(int))
		assert.Equal(t, MaxSize-i-1, q.Size())
	}

	assert.True(t, q.Empty())
}
