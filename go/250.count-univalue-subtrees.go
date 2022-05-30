package main

// Count Univalue Subtrees
//
// Medium
//
// https://leetcode.com/problems/count-univalue-subtrees/
//
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
	total := 0

	var uni func(t *TreeNode) bool
	uni = func(t *TreeNode) bool {
		if t == nil {
			return false
		}

		if t.Left == nil && t.Right == nil {
			total++
			return true
		}

		ok := true
		if uni(t.Left) {
			ok = ok && t.Val == t.Left.Val
		} else if t.Left != nil {
			ok = false
		}

		if uni(t.Right) {
			ok = ok && t.Val == t.Right.Val
		} else if t.Right != nil {
			ok = false
		}

		if ok {
			total++
		}

		return ok
	}

	uni(root)
	return total
}
