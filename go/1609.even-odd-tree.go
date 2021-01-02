package main

// medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	row := []*TreeNode{root}
	level := 0
	for len(row) != 0 {
		next := make([]*TreeNode, 0, len(row)*2)
		v := 0
		for i, n := range row {
			if (level^n.Val)&1 == 0 {
				return false
			}

			if i == 0 {
				v = n.Val
			} else if level%2 == 0 {
				if v >= n.Val {
					return false
				}
			} else {
				if v <= n.Val {
					return false
				}
			}

			v = n.Val

			if n.Left != nil {
				next = append(next, n.Left)
			}

			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		level++
		row = next
	}

	return true
}
