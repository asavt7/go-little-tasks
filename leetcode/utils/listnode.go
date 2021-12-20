package utils

import "fmt"

// ListNode /**
type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateNodes(in []int) *ListNode {

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

	return fmt.Sprintf("%+v", t.ToSlice())
}

func (t *ListNode) ToSlice() []int {
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
