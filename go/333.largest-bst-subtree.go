package main

// Largest BST Subtree
//
// Medium
//
// https://leetcode.com/problems/largest-bst-subtree
//
//
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestBSTSubtree(root *TreeNode) int {
	var walk func(*TreeNode) (int, bool, int, int)
	//                  node counts, is valid, min, max
	best := 0
	walk = func(r *TreeNode) (int, bool, int, int) {
		if r == nil {
			return 0, true, 0, 0
		}

		if r.Left == nil && r.Right == nil {
			best = max(best, 1)
			return 1, true, r.Val, r.Val
		}

		leftNodes, leftValid, leftMin, leftMax := walk(r.Left)
		rightNodes, rightValid, rightMin, rightMax := walk(r.Right)

		nodes := 1 + leftNodes + rightNodes
		valid := leftValid && rightValid &&
			(r.Left == nil || r.Val > leftMax) &&
			(r.Right == nil || r.Val < rightMin)

		if r.Left == nil {
			leftMin = r.Val
		}

		if r.Right == nil {
			rightMax = r.Val
		}

		if valid {
			best = max(best, nodes)
		}

		return nodes, valid, leftMin, rightMax
	}

	walk(root)
	return best
}
