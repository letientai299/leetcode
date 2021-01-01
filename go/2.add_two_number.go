package main

/**
 * Definition for singly-linked list.
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	n := new(ListNode)
	root := n

	x, y, r := l1, l2, 0

	for {
		t := x.Val + y.Val + r
		n.Val = t % 10
		r = t / 10
		x = x.Next
		y = y.Next
		if x == nil || y == nil {
			break
		}
		n.Next = new(ListNode)
		n = n.Next
	}

	var z *ListNode
	if x == nil {
		z = y
	}

	if y == nil {
		z = x
	}

	if z != nil {
		n.Next = new(ListNode)
		n = n.Next

		for z != nil {
			t := z.Val + r
			r = t / 10
			n.Val = t % 10
			if z.Next != nil {
				n.Next = new(ListNode)
				n = n.Next
				z = z.Next
			} else {
				break
			}
		}
	}

	if r != 0 {
		n.Next = new(ListNode)
		n.Next.Val = r
	}

	return root
}
