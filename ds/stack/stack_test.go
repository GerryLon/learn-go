package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	const MaxSize = 10
	s := New(MaxSize)
	assert.NotNil(t, s)
	defer s.Destroy()

	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Size())

	for i := 0; i < MaxSize; i++ {
		s.Push(i)
	}
	assert.True(t, s.Full())

	for i := MaxSize - 1; i >= 0; i-- {
		d, ok := s.Pop()
		assert.True(t, ok)
		assert.NotNil(t, d)

		assert.Equal(t, i, d.(int))
	}

	d, ok := s.Pop()
	assert.False(t, ok)
	assert.Nil(t, d)

	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Size())
}
