package main

// Equal Tree Partition
//
// Medium
//
// https://leetcode.com/problems/equal-tree-partition/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkEqualTree(root *TreeNode) bool {
	m := make(map[int]*TreeNode) // sum -> node

	var sum func(cur *TreeNode) int

	sum = func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		s := t.Val + sum(t.Left) + sum(t.Right)
		if m[s] == nil {
			m[s] = t
		}
		return s
	}

	all := sum(root)
	if all%2 != 0 {
		return false
	}

	half := all / 2
	node := m[half]
	return node != nil && node != root
}
