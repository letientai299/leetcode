package main

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (38.25%)
 * Total Accepted:    320.9K
 * Total Submissions: 706.9K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Given the head of a linked list, return the list after sorting it in
 * ascending order.
 *
 * Follow up: Can you sort the linked list in O(n logn) time and O(1) memory
 * (i.e. constant space)?
 *
 *
 * Example 1:
 *
 *
 * Input: head = [4,2,1,3]
 * Output: [1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [-1,5,3,4,0]
 * Output: [-1,0,3,4,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 5 * 10^4].
 * -10^5 <= Node.val <= 10^5
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

func sortList(head *ListNode) *ListNode {
	var sort func(r *ListNode) *ListNode

	sort = func(h *ListNode) *ListNode {
		if h == nil || h.Next == nil {
			return h
		}

		s, f := h, h.Next
		for f != nil {
			f = f.Next
			if f != nil {
				f = f.Next
			} else {
				break
			}
			s = s.Next
		}

		h1 := sort(s.Next)
		s.Next = nil
		h2 := sort(h)
		pre := &ListNode{}
		r := pre

		for h1 != nil && h2 != nil {
			if h1.Val < h2.Val {
				r.Next = h1
				h1 = h1.Next
			} else {
				r.Next = h2
				h2 = h2.Next
			}
			r = r.Next
		}

		for h1 != nil {
			r.Next = h1
			h1 = h1.Next
			r = r.Next
		}

		for h2 != nil {
			r.Next = h2
			h2 = h2.Next
			r = r.Next
		}

		return pre.Next

	}

	return sort(head)
}
