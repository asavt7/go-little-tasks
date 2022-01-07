package merge

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

type HeadOfNodes []*ListNode

func (h *HeadOfNodes) Len() int {
	return len(*h)
}

func (h *HeadOfNodes) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *HeadOfNodes) Swap(i, j int) {
	(*h)[i].Val, (*h)[j].Val = (*h)[j].Val, (*h)[i].Val
}

func (h *HeadOfNodes) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *HeadOfNodes) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	h := make(HeadOfNodes, 0)

	for _, list := range lists {
		for list != nil {
			heap.Push(&h, list)
			list = list.Next
		}
	}

	if len(h) == 0 {
		return nil
	}

	var head *ListNode = heap.Pop(&h).(*ListNode)
	cur := head
	for len(h) > 0 {
		cur.Next = heap.Pop(&h).(*ListNode)
		cur = cur.Next
	}
	cur.Next = nil

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
