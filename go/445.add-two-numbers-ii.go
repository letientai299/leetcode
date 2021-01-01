package main

/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (51.99%)
 * Total Accepted:    211K
 * Total Submissions: 377.4K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var a, b []int
	for n := l1; n != nil; n = n.Next {
		a = append(a, n.Val)
	}

	for n := l2; n != nil; n = n.Next {
		b = append(b, n.Val)
	}

	if len(a) < len(b) {
		a, b = b, a
	}

	rem := 0
	for i, j := len(a)-1, len(b)-1; i >= 0; i-- {
		if j >= 0 {
			rem += b[j]
			j--
		}
		rem += a[i]
		a[i] = rem % 10
		rem /= 10
	}

	head := &ListNode{Val: rem}
	n := head
	for _, v := range a {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}

	if head.Val != 0 {
		return head
	}

	return head.Next
}
