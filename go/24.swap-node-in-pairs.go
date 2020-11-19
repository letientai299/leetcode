package main

// medium

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head.Next

	for a != nil && b != nil {
		a.Val, b.Val = b.Val, a.Val
		a = b.Next
		if a != nil {
			b = a.Next
		}
	}
	return head
}
