package summ

//https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	var val, overflow int

	for l1 != nil && l2 != nil {
		val = l1.Val + l2.Val + overflow
		overflow = val / 10
		cur.Val = val % 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			newCur := &ListNode{}
			cur.Next = newCur
			cur = newCur
		}
	}
	for l1 != nil {
		val = l1.Val + overflow
		overflow = val / 10
		cur.Val = val % 10

		if l1.Next != nil {
			newCur := &ListNode{}
			cur.Next = newCur
			cur = newCur
		}
		l1 = l1.Next
	}
	for l2 != nil {
		val = l2.Val + overflow
		overflow = val / 10
		cur.Val = val % 10

		if l2.Next != nil {
			newCur := &ListNode{}
			cur.Next = newCur
			cur = newCur
		}
		l2 = l2.Next
	}

	if overflow > 0 {
		newCur := &ListNode{Val: overflow}
		cur.Next = newCur
		cur = newCur
	}

	return head
}
