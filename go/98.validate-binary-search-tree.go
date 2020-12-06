package main

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (26.69%)
 * Total Accepted:    826.3K
 * Total Submissions: 2.9M
 * Testcase Example:  '[2,1,3]'
 *
 * Given the root of a binary tree, determine if it is a valid binary search
 * tree (BST).
 *
 * A valid BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is
 * 4.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -2^31 <= Node.val <= 2^31 - 1
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

func isValidBST(t *TreeNode) bool {
	var valid func(t *TreeNode) (bool, int, int)

	valid = func(t *TreeNode) (bool, int, int) {
		if t == nil {
			return true, math.MinInt32, math.MaxInt32
		}

		if t.Left == nil && t.Right == nil {
			return true, t.Val, t.Val
		}

		leftOK, leftMin, leftMax := valid(t.Left)
		rightOK, rightMin, rightMax := valid(t.Right)

		if t.Left == nil {
			if !rightOK || rightMin <= t.Val {
				return false, 0, 0
			}

			return true, t.Val, rightMax
		}

		if t.Right == nil {
			if !leftOK || leftMax >= t.Val {
				return false, 0, 0
			}

			return true, leftMin, t.Val
		}

		if !(leftOK && rightOK) || rightMin <= t.Val || leftMax >= t.Val {
			return false, 0, 0
		}

		return true, leftMin, rightMax
	}

	ok, _, _ := valid(t)
	return ok
}
