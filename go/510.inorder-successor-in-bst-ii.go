package main

// Inorder Successor in BST II
//
// # Medium
//
// https://leetcode.com/problems/inorder-successor-in-bst-ii
func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		for node = node.Right; node != nil && node.Left != nil; node = node.Left {
		}
		return node
	}

	// this uses BST property
	// p := node.Parent
	// for p != nil && p.Val < node.Val {
	//   p = p.Parent
	// }

	// this only care whether current node is its parent's left.
	// which solve the follow-up question of "not checking node value"
	p := node.Parent
	for p != nil && p.Right == node {
		node = p
		p = p.Parent
	}
	return p
}
