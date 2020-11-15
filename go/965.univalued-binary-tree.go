package main

/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
 *
 * https://leetcode.com/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (66.90%)
 * Total Accepted:    103.9K
 * Total Submissions: 153.7K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * A binary tree is univalued if every node in the tree has the same value.
 *
 * Return true if and only if the given tree is univalued.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,1,1,1,1,null,1]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,5,2]
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the given tree will be in the range [1, 100].
 * Each node's value will be an integer in the range [0, 99].
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

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	v := root.Val

	var check func(n *TreeNode) bool

	check = func(n *TreeNode) bool {
		if n == nil {
			return true
		}
		if n.Val != v {
			return false
		}
		return check(n.Left) && check(n.Right)
	}

	return check(root.Left) && check(root.Right)
}
