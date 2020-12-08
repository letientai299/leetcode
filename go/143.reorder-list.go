package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (33.23%)
 * Total Accepted:    290.2K
 * Total Submissions: 727.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a singly linked list L: L0→L1→…→Ln-1→Ln,
 * reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 * Example 1:
 *
 *
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 *
 * Example 2:
 *
 *
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	var swap func(r *ListNode) *ListNode
	swap = func(r *ListNode) *ListNode {
		if r == nil || r.Next == nil || r.Next.Next == nil {
			return r
		}

		s, f, p := r, r, r
		for f.Next != nil {
			s = s.Next
			p = f
			f = f.Next
			if f.Next == nil {
				break
			}
			p = f
			f = f.Next
		}

		f.Next = r.Next
		p.Next = nil
		r.Next = f
		return r
	}

	r := head
	for r != nil {
		r = swap(r)
		r = r.Next
		if r != nil {
			r = r.Next
		}
	}

	return
}
