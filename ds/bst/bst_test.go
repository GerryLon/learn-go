package bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(t *testing.T) {
	tree := New(50)
	assert.NotNil(t, tree)

	assert.Equal(t, 1, tree.Height())

	// 插入重复数据失败
	assert.False(t, tree.Insert(50))

	assert.True(t, tree.Insert(20))
	assert.True(t, tree.Insert(60))
	assert.Equal(t, 2, tree.height)

	tree.LevelTraverse(func(data int) {
		fmt.Println(data)
	})

	for i := 10; i > 0; i-- {
		assert.True(t, tree.Insert(i))
	}
	assert.Equal(t, 12, tree.height)
}
