package main

// Binary Tree Upside Down
//
// Medium
//
// https://leetcode.com/problems/binary-tree-upside-down/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	l := root.Left
	newRoot := upsideDownBinaryTree(l)
	r := upsideDownBinaryTree(root.Right)
	root.Left = nil
	root.Right = nil

	l.Left = r
	l.Right = root
	return newRoot
}
