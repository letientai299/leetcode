package main

import "math"

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 *
 * https://leetcode.com/problems/insertion-sort-list/description/
 *
 * algorithms
 * Medium (38.93%)
 * Total Accepted:    219K
 * Total Submissions: 498.9K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Sort a linked list using insertion sort.
 *
 *
 *
 *
 *
 * A graphical example of insertion sort. The partial sorted list (black)
 * initially contains only the first element in the list.
 * With each iteration one element (red) is removed from the input data and
 * inserted in-place into the sorted list
 *
 *
 *
 *
 *
 * Algorithm of Insertion Sort:
 *
 *
 * Insertion sort iterates, consuming one input element each repetition, and
 * growing a sorted output list.
 * At each iteration, insertion sort removes one element from the input data,
 * finds the location it belongs within the sorted list, and inserts it
 * there.
 * It repeats until no input elements remain.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 4->2->1->3
 * Output: 1->2->3->4
 *
 *
 * Example 2:
 *
 *
 * Input: -1->5->3->4->0
 * Output: -1->0->3->4->5
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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{
		Val:  math.MinInt32,
		Next: head,
	}

	pr := head
	r := head.Next
	for r != nil {
		c := pre
		for c.Next != nil && c.Next.Val < r.Val {
			c = c.Next
		}

		if c.Next == r {
			pr = r
			r = r.Next
			continue
		}

		n := r.Next
		pr.Next = n
		r.Next = c.Next
		c.Next = r
		r = n
	}

	return pre.Next
}
