package main

/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (43.58%)
 * Total Accepted:    270.3K
 * Total Submissions: 548K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given the head of a singly linked list where elements are sorted in
 * ascending order, convert it to a height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [-10,-3,0,5,9]
 * Output: [0,-3,9,-10,null,5]
 * Explanation: One possible answer is [0,-3,9,-10,null,5], which represents
 * the shown height balanced BST.
 *
 *
 * Example 2:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [0]
 * Output: [0]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,3]
 * Output: [3,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in head is in the range [0, 2 * 10^4].
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	tmp := head
	n := 0
	for head != nil {
		n++
		head = head.Next
	}

	var build func(s, e int) *TreeNode

	build = func(s, e int) *TreeNode {
		if s > e {
			return nil
		}

		mid := (s + e) / 2
		t := &TreeNode{}
		t.Left = build(s, mid-1)
		t.Val = tmp.Val
		tmp = tmp.Next
		t.Right = build(mid+1, e)
		return t
	}

	return build(0, n-1)
}
