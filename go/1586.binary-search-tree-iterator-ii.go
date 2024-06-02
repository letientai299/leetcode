package main

// Binary Search Tree Iterator II
//
// # Medium
//
// https://leetcode.com/problems/binary-search-tree-iterator-ii
type BSTIterator struct {
	cur *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	// convert the tree into doubly linked list
	var conv func(*TreeNode) (*TreeNode, *TreeNode)
	conv = func(n *TreeNode) (*TreeNode, *TreeNode) {
		if n == nil {
			return nil, nil
		}

		var head, tail *TreeNode

		if n.Left != nil {
			head, n.Left = conv(n.Left)
			n.Left.Right = n
		} else {
			head = n
		}

		if n.Right != nil {
			n.Right, tail = conv(n.Right)
			n.Right.Left = n
		} else {
			tail = n
		}

		return head, tail
	}

	head, _ := conv(root)
	cur := &TreeNode{Right: head, Val: head.Val - 1}
	head.Left = cur
	return BSTIterator{
		cur: cur,
	}
}

func (this *BSTIterator) HasNext() bool {
	return this.cur.Right != nil
}

func (this *BSTIterator) Next() int {
	this.cur = this.cur.Right
	return this.cur.Val
}

func (this *BSTIterator) HasPrev() bool {
	return this.cur.Left != nil
}

func (this *BSTIterator) Prev() int {
	this.cur = this.cur.Left
	return this.cur.Val
}
