package main

// medium
// https://leetcode.com/problems/partition-list/

func partition(head *ListNode, x int) *ListNode {
	a := &ListNode{Next: head}
	b := &ListNode{}
	r := b
	parent := a
	n := head
	for n != nil {
		if n.Val >= x {
			parent.Next = n.Next
			r.Next = n
			r = n
			n.Next, n = nil, n.Next
		} else {
			parent = n
			n = n.Next
		}
	}

	parent.Next = b.Next
	return a.Next
}
