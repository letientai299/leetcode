package main

// medium
// -o goland1 1669. Merge In Between Linked Lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(x *ListNode, a int, b int, y *ListNode) *ListNode {
	if x == nil {
		return y
	}

	fake := &ListNode{Next: x}
	n := fake
	b -= a
	for n.Next != nil && a > 0 {
		n = n.Next
		a--
	}

	r := n.Next
	n.Next = y

	if r == nil {
		return fake.Next
	}

	for r != nil && b > 0 {
		r = r.Next
		b--
	}

	if r == nil {
		return fake.Next
	}

	for y != nil && y.Next != nil {
		y = y.Next
	}

	if y == nil {
		n.Next = r.Next
	} else {
		y.Next = r.Next
	}

	return fake.Next
}
