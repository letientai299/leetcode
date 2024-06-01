package main

// Split BST
//
// # Medium
//
// https://leetcode.com/problems/split-bst
func splitBST(root *TreeNode, target int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	}

	if root.Val <= target {
		ts := splitBST(root.Right, target)
		l, r := ts[0], ts[1]
		root.Right = l
		return []*TreeNode{root, r}
	}

	// root.Val > target
	ts := splitBST(root.Left, target)
	l, r := ts[0], ts[1]
	root.Left = r
	return []*TreeNode{l, root}
}
