package bst

import "github.com/GerryLon/learn-go/ds/queue"

// binary search tree
type BST struct {
	// left sub tree and right sub tree
	lChild *BST
	rChild *BST

	data   int // data of tree's root node
	height int // height of the tree
}

func New(d int) *BST {
	return &BST{data: d, height: 1}
}

// 如果有重复数据, 会插入失败
// newRoot: 新节点
func (t *BST) Insert(d int) (ok bool) {
	tmp := t

	for {
		if d == tmp.data {
			return false
		} else if d < tmp.data {
			if tmp.lChild == nil {
				tmp.lChild = New(d)
				break
			} else {
				tmp = tmp.lChild
			}
		} else {
			if tmp.rChild == nil {
				tmp.rChild = New(d)
				break
			} else {
				tmp = tmp.rChild
			}
		}
	}

	lh := t.lChild.Height()
	rh := t.rChild.Height()
	max := lh
	if rh > lh {
		max = lh
	}
	t.height = max + 1

	return true
}

func (t *BST) Height() int {
	if t == nil {
		return 0
	} else {
		return t.height
	}
}

func (t *BST) LevelTraverse(visit func(data int)) {
	tmp := t

	q := queue.New(100)
	q.Enqueue(tmp)

	for !q.Empty() {
		tNode, _ := q.Dequeue()
		node := tNode.(*BST)
		visit(node.data)

		if node.lChild != nil {
			q.Enqueue(tmp.lChild)
		}

		if node.rChild != nil {
			q.Enqueue(node.rChild)
		}
	}
}
