package main

import (
	"fmt"
)

// ListNode /**
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var sz int = 1
	if head.Next == nil {
		return nil
	}
	left, right := head, head.Next
	for offset := 0; offset < n && right != nil; offset++ {
		right = right.Next
		sz++
	}

	for right != nil {
		right = right.Next
		left = left.Next
		sz++
	}

	if n == 1 {
		if left == head && n == sz {
			return head.Next
		}
		left.Next = nil
	} else {
		if left == head && n == sz {
			return head.Next
		}
		left.Next = left.Next.Next
	}

	return head
}

func generateNodes(in []int) *ListNode {

	if len(in) == 0 {
		return nil
	}

	first := &ListNode{
		Val: in[0],
	}
	if len(in) == 1 {
		return first
	}

	prev := first
	for i := 1; i < len(in)-1; i++ {
		cur := &ListNode{Val: in[i]}
		prev.Next = cur
		prev = cur
	}
	prev.Next = &ListNode{Val: in[len(in)-1]}

	return first
}

func (t *ListNode) String() string {

	return fmt.Sprintf("%+v", t.toSlice())
}

func (t *ListNode) toSlice() []int {
	var res []int
	if t == nil {
		return make([]int, 0, 0)
	}
	cur := t
	for cur.Next != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	res = append(res, cur.Val)
	return res
}
