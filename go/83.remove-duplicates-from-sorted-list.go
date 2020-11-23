package main

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (42.00%)
 * Total Accepted:    306.3K
 * Total Submissions: 729.2K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
func deleteDuplicates_83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head.Next
	prev := head
	for node != nil {
		if node.Val == prev.Val {
			node = node.Next
			prev.Next = node
			continue
		}

		prev = node
		node = node.Next
	}
	return head
}
