package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New()
	assert.NotNil(t, s)

	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Size())

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	for i := 9; i >=0; i-- {
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
