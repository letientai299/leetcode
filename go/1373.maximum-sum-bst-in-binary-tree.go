package main

import "math"

// Maximum Sum BST in Binary Tree
//
// Hard
//
// https://leetcode.com/problems/maximum-sum-bst-in-binary-tree/
//
// Given a **binary tree** `root`, return _the maximum sum of all keys of
// **any** sub-tree which is also a Binary Search Tree (BST)_.
//
// Assume a BST is defined as follows:
//
// - The left subtree of a node contains only nodes with keys **less than** the
// node's key.
// - The right subtree of a node contains only nodes with keys **greater than**
// the node's key.
// - Both the left and right subtrees must also be binary search trees.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/01/30/sample_1_1709.png)
//
// ```
// Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// Output: 20
// Explanation: Maximum sum in a valid Binary search tree is obtained in root
// node with key equal to 3.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/01/30/sample_2_1709.png)
//
// ```
// Input: root = [4,3,null,1,2]
// Output: 2
// Explanation: Maximum sum in a valid Binary search tree is obtained in a
// single root node with key equal to 2.
//
// ```
//
// **Example 3:**
//
// ```
// Input: root = [-4,-2,-5]
// Output: 0
// Explanation: All values are negatives. Return an empty BST.
//
// ```
//
// **Example 4:**
//
// ```
// Input: root = [2,1,3]
// Output: 6
//
// ```
//
// **Example 5:**
//
// ```
// Input: root = [5,4,8,3,null,6,3]
// Output: 7
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 4 * 104]`.
// - `-4 * 104 <= Node.val <= 4 * 104`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
	var walk func(n *TreeNode) (sum int, valid bool, min, max int)
	best := 0

	walk = func(n *TreeNode) (sum int, valid bool, min, max int) {
		if n == nil {
			return 0, true, math.MinInt32, math.MaxInt32
		}

		if n.Left == nil && n.Right == nil {
			if best < n.Val {
				best = n.Val
			}
			return n.Val, true, n.Val, n.Val
		}

		lsum, lok, lmin, lmax := walk(n.Left)
		rsum, rok, rmin, rmax := walk(n.Right)

		ok := lok && rok && (n.Left == nil || lmax < n.Val) && (n.Right == nil || n.Val < rmin)
		sum = n.Val + lsum + rsum

		if !ok {
			return 0, false, 0, 0
		}

		if best < sum {
			best = sum
		}

		if n.Left == nil {
			lmin = n.Val
		}

		if n.Right == nil {
			rmax = n.Val
		}

		return sum, ok, lmin, rmax
	}

	walk(root)
	return best
}
