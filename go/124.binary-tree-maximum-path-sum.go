package main

import "math"

// Binary Tree Maximum Path Sum
//
// Hard
//
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
//
// A **path** in a binary tree is a sequence of nodes where each pair of
// adjacent nodes in the sequence has an edge connecting them. A node can only
// appear in the sequence **at most once**. Note that the path does not need to
// pass through the root.
//
// The **path sum** of a path is the sum of the node's values in the path.
//
// Given the `root` of a binary tree, return _the maximum **path sum** of any
// path_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)
//
// ```
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 =
// 6.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)
//
// ```
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7
// = 42.
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 3 * 104]`.
// - `-1000 <= Node.val <= 1000`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	best := math.MinInt32
	var walk func(node *TreeNode) int // return max path sum from node to its leaves

	walk = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		left := walk(n.Left)
		right := walk(n.Right)

		if left <= 0 && right <= 0 {
			if best < n.Val {
				best = n.Val
			}
			return n.Val
		}

		path := n.Val
		if left > 0 {
			path += left
		}

		if right > 0 {
			path += right
		}

		if best < path {
			best = path
		}

		if left > right {
			return n.Val + left
		}

		return n.Val + right
	}

	walk(root)
	return best
}
