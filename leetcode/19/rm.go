package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	offset := 1

	left, right := head, head

	for offset < n+1 {
		right = right.Next
		offset++
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

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

	cur := t
	for cur.Next != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	res = append(res, cur.Val)
	return res
}
