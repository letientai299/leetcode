package main

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 *
 * https://leetcode.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (48.80%)
 * Total Accepted:    119.5K
 * Total Submissions: 244.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *      ⁠   1
 *      ⁠  / \
 *      ⁠ 2   3
 *   ⁠   /  \
 * ⁠    4    5
 *
 * There are two left leaves in the binary tree, with values 9 and 15
 * respectively. Return 24.
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}

		f(node.Right)

		if node.Left == nil {
			return
		}

		node = node.Left

		if node.Left == nil && node.Right == nil {
			sum += node.Val
		}
		f(node)
	}
	f(root)
	return sum
}
