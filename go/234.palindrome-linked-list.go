package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (35.51%)
 * Total Accepted:    239.4K
 * Total Submissions: 674.3K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var values []int
	fast := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			head = head.Next
			break
		}
		fast = fast.Next
		values = append(values, head.Val)
		head = head.Next
	}

	if values == nil {
		return true
	}

	i := len(values) - 1
	for head != nil && i >= 0 {
		if head.Val != values[i] {
			return false
		}
		head = head.Next
		i--
	}
	return true
}
