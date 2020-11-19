package main

// medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	s, f := 0, 0
	for slow != nil && fast != nil {
		fast = fast.Next
		f++
		if fast != nil {
			fast = fast.Next
			f++
		} else {
			break
		}
		slow = slow.Next
		s++
	}

	if n > f {
		return head
	}

	if n == f {
		return head.Next
	}

	if n < f/2 {
		n = s - n
		for n > 1-f%2 {
			slow = slow.Next
			n--
		}
		if slow.Next != nil {
			slow.Next = slow.Next.Next
		}
		return head
	}

	n = f - n - 1
	slow = head
	for n > 0 {
		slow = slow.Next
		n--
	}
	slow.Next = slow.Next.Next

	return head
}
