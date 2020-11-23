package main

// medium
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fake := &ListNode{}

	node := head
	prev := fake
	for node != nil {
		c := 1
		for node.Next != nil && node.Val == node.Next.Val {
			c++
			node = node.Next
		}

		if c == 1 {
			prev.Next = node
			prev = node
		}
		node = node.Next
	}
	prev.Next = nil
	return fake.Next
}
