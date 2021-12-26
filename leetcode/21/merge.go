package merge

import . "github.com/asavt7/go-little-tasks/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var h *ListNode
	if list1.Val < list2.Val {
		h = list1
		list1 = list1.Next
	} else {
		h = list2
		list2 = list2.Next
	}
	cur := h
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	for list1 != nil {
		cur.Next = list1
		list1 = list1.Next
		cur = cur.Next
	}

	for list2 != nil {
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}

	return h
}
