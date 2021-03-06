package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (45.87%)
 * Total Accepted:    521.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	root := new(ListNode)
	z := root
	x := l1
	y := l2
	for x != nil && y != nil {
		if x.Val <= y.Val {
			z.Next = x
			x = x.Next
		} else {
			z.Next = y
			y = y.Next
		}

		z = z.Next
	}

	rest := x
	if x == nil {
		rest = y
	}

	for rest != nil {
		z.Next = rest
		rest = rest.Next
		if rest != nil {
			z = z.Next
		}
	}

	return root.Next
}
