package main

// medium
// https://leetcode.com/problems/rotate-list/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	n := 0
	r := head
	tail := r
	for r != nil {
		tail = r
		r = r.Next
		n++
	}

	k %= n
	if k == 0 {
		return head
	}

	k = n - k
	r = head
	for k > 1 {
		r = r.Next
		k--
	}
	t := r.Next
	r.Next = nil
	tail.Next = head
	return t
}
