package main

// Inorder Successor in BST
//
// # Medium
//
// https://leetcode.com/problems/inorder-successor-in-bst
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	for root != nil {
		if root.Val <= p.Val {
			root = root.Right
		} else {
			res = root
			root = root.Left
		}
	}
	return res
}
