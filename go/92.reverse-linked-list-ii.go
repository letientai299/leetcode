package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (36.59%)
 * Total Accepted:    309K
 * Total Submissions: 774.6K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	fake := &ListNode{
		Next: head,
	}

	p := fake
	r := head
	for m > 1 {
		p = r
		r = r.Next
		m--
		n--
	}

	p.Next = nil

	x := r
	c := r
	r = r.Next
	c.Next = nil

	for n > 1 {
		t := c
		c = r.Next
		r.Next = t
		p.Next = r
		r = c
		c = p.Next
		n--
	}
	x.Next = r
	return fake.Next
}
