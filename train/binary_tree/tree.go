package main

import (
	"errors"
	"fmt"
)

var ErrNilTreeNode = errors.New("ERROR - node is nil")
var ErrNotSuchNode = errors.New("ERROR - no such node")

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return ErrNilTreeNode
	}
	cur := t
	prev := t

	for cur != nil {
		prev = cur
		if value >= cur.val {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	if value >= prev.val {
		prev.right = &TreeNode{
			val: value,
		}
	} else {
		prev.left = &TreeNode{
			val: value,
		}
	}
	return nil
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}
func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMin()
}

func (t *TreeNode) Walk(walkFunc func(t *TreeNode)) {
	if t != nil {
		t.left.Walk(walkFunc)
		walkFunc(t)
		t.right.Walk(walkFunc)
	}
}

func (t *TreeNode) Remove(val int) (*TreeNode, error) {
	if t == nil {
		return nil, ErrNotSuchNode
	}
	return t.remove(t, val), nil

}

func (t *TreeNode) remove(root *TreeNode, val int) *TreeNode {
	if t == nil {
		return nil
	}

	if val < t.val {
		t.left = t.left.remove(root, val)
		return t
	}
	if val > t.val {
		t.right = t.right.remove(root, val)
		return t
	}

	if t.left == nil && t.right == nil {
		return nil
	}

	if t.left == nil {
		return t.right
	}

	if t.right == nil {
		return t.left
	}

	// val == t.val
	prev, nextValue := t, t.right
	leftChild := false
	for nextValue.left != nil {
		prev = nextValue
		nextValue = nextValue.left
		leftChild = true
	}

	// no left children of right child of root node
	if nextValue == t.right {
		nextValue.left = t.left
		return nextValue
	}

	if leftChild {
		prev.left = nil
	} else {
		prev.right = nil
	}
	nextValue.left = t.left
	nextValue.right = t.right
	return nextValue
}

func main() {
	var t *TreeNode = &TreeNode{
		val: 5,
	}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)

	t.Insert(7)
	t.Insert(58)
	t.Insert(8)
	t.Insert(31)
	t.Insert(13)

	t, _ = t.Remove(4)
	//fmt.Println(t.FindMin())

	t.Walk(func(t *TreeNode) {
		fmt.Println(t.val)
	})
}
