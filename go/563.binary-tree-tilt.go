package main

/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
 *
 * https://leetcode.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (46.98%)
 * Total Accepted:    52.1K
 * Total Submissions: 110.8K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a binary tree, return the tilt of the whole tree.
 *
 * The tilt of a tree node is defined as the absolute difference between the
 * sum of all left subtree node values and the sum of all right subtree node
 * values. Null node has tilt 0.
 *
 * The tilt of the whole tree is defined as the sum of all nodes' tilt.
 *
 * Example:
 *
 * Input:
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * Output: 1
 * Explanation:
 * Tilt of node 2 : 0
 * Tilt of node 3 : 0
 * Tilt of node 1 : |2-3| = 1
 * Tilt of binary tree : 0 + 0 + 1 = 1
 *
 *
 *
 * Note:
 *
 * The sum of node values in any subtree won't exceed the range of 32-bit
 * integer.
 * All the tilt values won't exceed the range of 32-bit integer.
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
func findTilt(root *TreeNode) int {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	getVal := func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		return t.Val
	}

	// this function update the root node value to be sum of its value and all its
	// child, and return the tilt of its child
	var tilt func(t *TreeNode) int
	tilt = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		tl := tilt(t.Left)
		te := tilt(t.Right)
		t.Val += getVal(t.Left)
		t.Val += getVal(t.Right)
		return tl + te + abs(getVal(t.Left), getVal(t.Right))
	}

	return tilt(root)
}
