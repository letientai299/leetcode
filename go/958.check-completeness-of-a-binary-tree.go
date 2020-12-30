package main

/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
 *
 * https://leetcode.com/problems/check-completeness-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (50.41%)
 * Total Accepted:    69.4K
 * Total Submissions: 132.5K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * Given the root of a binary tree, determine if it is a complete binary tree.
 *
 * In a complete binary tree, every level, except possibly the last, is
 * completely filled, and all nodes in the last level are as far left as
 * possible. It can have between 1 and 2^h nodes inclusive at the last level
 * h.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,4,5,6]
 * Output: true
 * Explanation: Every level before the last is full (ie. levels with
 * node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are
 * as far left as possible.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,3,4,5,null,7]
 * Output: false
 * Explanation: The node with value 7 isn't as far left as possible.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 100].
 * 1 <= Node.val <= 1000
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
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	row := []*TreeNode{root}
	level := 1
	last := false
	for len(row) != 0 {
		level *= 2
		next := make([]*TreeNode, 0, len(row)*2)
		hasNil := false
		for _, n := range row {
			if n.Left != nil {
				if hasNil {
					return false
				}
				next = append(next, n.Left)
			} else {
				hasNil = true
			}

			if n.Right != nil {
				if hasNil {
					return false
				}
				next = append(next, n.Right)
			} else {
				hasNil = true
			}

		}

		if last && len(next) != 0 {
			return false
		}

		if len(next) != level && len(next) != 0 {
			last = true
		}

		// next must be last level
		row = next
	}

	return true
}
