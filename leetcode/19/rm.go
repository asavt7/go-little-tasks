package main

import . "github.com/asavt7/go-little-tasks/leetcode/utils"

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
