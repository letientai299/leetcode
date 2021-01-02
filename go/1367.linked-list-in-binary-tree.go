package main

// medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	var check func(l *ListNode, t *TreeNode) bool
	check = func(l *ListNode, t *TreeNode) bool {
		if l.Next == nil {
			return true
		}

		l = l.Next
		if t.Left != nil && t.Left.Val == l.Val && check(l, t.Left) {
			return true
		}

		if t.Right != nil && t.Right.Val == l.Val && check(l, t.Right) {
			return true
		}

		return false
	}

	row := []*TreeNode{root}
	for len(row) != 0 {
		next := make([]*TreeNode, 0, 2*len(row))
		for _, n := range row {
			if n.Val == head.Val && check(head, n) {
				return true
			}

			if n.Left != nil {
				next = append(next, n.Left)
			}

			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		row = next
	}

	return false
}
