package merge

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, cur *ListNode

	var curMin int = math.MaxInt64
	//init step
	for _, node := range lists {
		if node.Val < curMin{
			curMin = node.Val
		}
	}

	empties:= make([]int, 0,len(lists))
	for len(lists)>0 {

		for i, node := range lists {
			if node == nil{
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}
		}

		empties = empties[:0]


	}

	return head
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

func GenerateNodeLists(in [][]int) []*ListNode {
	res := make([]*ListNode, 0, len(in))
	for _, ints := range in {
		res = append(res, GenerateNodes(ints))
	}
	return res
}
