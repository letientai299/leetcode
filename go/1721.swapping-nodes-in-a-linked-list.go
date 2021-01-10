package main

// medium
// 1721. Swapping Nodes in a Linked List
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	k--
	fake := &ListNode{Next: head}
	left := 0
	full := 0
	f, s := head, head
	for f != nil && left < k {
		f = f.Next
		full++
		if f != nil {
			f = f.Next
			left++
			full++
			s = s.Next
		}
	}

	if f != nil {
		// k < half of the list
		t := s
		for f != nil {
			f = f.Next
			full++
			if f != nil {
				f = f.Next
				left++
				full++
				s = s.Next
			}
		}

		left -= k
		f = s
		s = t
		for left > 1-full%2 {
			f = f.Next
			left--
		}
	} else {
		// k >= half of the list
		full -= k + 1
		f = head
		for full > 0 {
			f = f.Next
			full--
		}

		for left < k {
			s = s.Next
			left++
		}
	}

	f.Val, s.Val = s.Val, f.Val
	return fake.Next
}
