package main

import "math"

/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 *
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (51.00%)
 * Total Accepted:    42.4K
 * Total Submissions: 83.2K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * Given a Binary Search Tree (BST) with the root node root, return the minimum
 * difference between the values of any two different nodes in the tree.
 *
 * Example :
 *
 *
 * Input: root = [4,2,6,1,3,null,null]
 * Output: 1
 * Explanation:
 * Note that root is a TreeNode object, not an array.
 *
 * The given tree [4,2,6,1,3,null,null] is represented by the following
 * diagram:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * while the minimum difference in this tree is 1, it occurs between node 1 and
 * node 2, also between node 3 and node 2.
 *
 *
 * Note:
 *
 *
 * The size of the BST will be between 2 and 100.
 * The BST is always valid, each node's value is an integer, and each node's
 * value is different.
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
func minDiffInBST(root *TreeNode) int {
	var inOrder func(n *TreeNode) []int
	inOrder = func(n *TreeNode) []int {
		if n == nil {
			return nil
		}

		vals := inOrder(n.Left)
		vals = append(vals, n.Val)
		return append(vals, inOrder(n.Right)...)
	}

	vals := inOrder(root)

	m := math.MaxInt32
	for i := 0; i < len(vals)-1; i++ {
		diff := vals[i+1] - vals[i]
		if diff == 0 {
			continue
		}

		if m > diff {
			m = diff
		}
	}
	return m
}
