package main

// Convert Binary Search Tree to Sorted Doubly Linked List
//
// # Medium
//
// https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list
func treeToDoublyList(root *Node) *Node {
	var conv func(n *Node) (*Node, *Node)
	conv = func(n *Node) (*Node, *Node) {
		if n == nil {
			return nil, nil
		}

		var head *Node
		var tail *Node

		if n.Left == nil {
			head = n
		} else {
			head, n.Left = conv(n.Left)
			n.Left.Right = n
		}

		if n.Right == nil {
			tail = n
		} else {
			n.Right, tail = conv(n.Right)
			n.Right.Left = n
		}

		head.Left = tail
		tail.Right = head
		return head, tail
	}

	root, _ = conv(root)
	return root
}
