package main

/*
 * @lc app=leetcode id=270 lang=golang
 *
 * [270] Closest Binary Search Tree Value
 *
 * https://leetcode.com/problems/closest-binary-search-tree-value/description/
 *
 * algorithms
 * Easy (45.40%)
 * Total Accepted:    160.8K
 * Total Submissions: 327.7K
 * Testcase Example:  '[4,2,5,1,3]\n3.714286'
 *
 * Given a non-empty binary search tree and a target value, find the value in
 * the BST that is closest to the target.
 *
 * Note:
 *
 *
 * Given target value is a floating point.
 * You are guaranteed to have only one unique value in the BST that is closest
 * to the target.
 *
 *
 * Example:
 *
 *
 * Input: root = [4,2,5,1,3], target = 3.714286
 *
 * ⁠   4
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \
 * 1   3
 *
 * Output: 4
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	dis := func(v int) float64 {
		fv := float64(v)
		if fv >= target {
			return fv - target
		}
		return target - fv
	}

	best := root.Val
	minDiff := dis(root.Val)
	var find func(root *TreeNode) int

	find = func(root *TreeNode) int {
		diff := dis(root.Val)
		if diff < minDiff {
			best = root.Val
			minDiff = diff
		}

		if float64(root.Val) == target {
			return root.Val
		}

		if root.Left == nil && root.Right == nil {
			return root.Val
		}

		if root.Left == nil {
			if target < float64(root.Val) {
				return root.Val
			}
			return find(root.Right)
		}

		if root.Right == nil {
			if target > float64(root.Val) {
				return root.Val
			}
			return find(root.Left)
		}

		next := root.Left
		if target > float64(root.Val) {
			next = root.Right
		}

		return find(next)
	}

	find(root)
	return best
}
