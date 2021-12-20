package sortlinkedlist

import (
	"fmt"
	. "github.com/asavt7/go-little-tasks/leetcode/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return SortList(head)
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := head, head
	tmp, midNext := r.Next, r.Next
	mid := head

	for r.Next != nil {
		fmt.Printf("%+v\n", l.ToSlice())

		if r.Next.Val <= l.Val {
			tmp = r.Next.Next
			r.Next.Next = l
			l = r.Next
			r.Next = tmp
			continue
		}
		if r.Next.Val > r.Val {
			r = r.Next
			continue
		}

		mid = l
		for r.Next.Val > mid.Next.Val {
			mid = mid.Next
		}
		tmp = r.Next.Next
		midNext = mid.Next
		mid.Next = r.Next
		mid.Next.Next = midNext
		r.Next = tmp
		continue
	}

	return l
}
