package main

// medium

// 32.7%
//func detectCycle(head *ListNode) *ListNode {
//	m := make(map[*ListNode]struct{})
//	n := head
//	for n != nil {
//		if _, ok := m[n]; ok {
//			return n
//		}
//		m[n] = struct{}{}
//		n = n.Next
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	f := head.Next
	if f != nil {
		f = f.Next
	}
	s := head.Next
	for f != nil && f != s {
		s = s.Next
		f = f.Next
		if f == nil {
			return nil
		}
		f = f.Next
	}

	if f == nil {
		return nil
	}

	s = head
	for f != s {
		s = s.Next
		f = f.Next
	}

	// p + c
	// s: k = p + xc + m
	// f: 2k = p + yc + m = p + 2xc + 2m
	// => k = (y-x)c
	// => (p+m)%c = 0
	// => let s run another k
	return s
}
