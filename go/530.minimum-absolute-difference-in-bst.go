package main

import "math"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (50.51%)
 * Total Accepted:    59.4K
 * Total Submissions: 117.6K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * Output:
 * 1
 *
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 *
 *
 *
 *
 * Note: There are at least two nodes in this BST.
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	var cache []int
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {

		if root == nil {
			return
		}

		travel(root.Left)
		cache = append(cache, root.Val)
		travel(root.Right)
	}

	travel(root)

	var min = func(a, b int) int {
		if a > b {
			return b
		}

		return a
	}

	var abs = func(a, b int) int {
		if a > b {
			return a - b
		}

		return b - a
	}

	minDiff := math.MaxInt32
	for i := 0; i < len(cache)-1; i++ {
		current := abs(cache[i], cache[i+1])
		minDiff = min(current, minDiff)
	}

	return minDiff
}

// This solution is only 93.58% speed and 69.84 mem
/*
func getMinimumDifference(root *TreeNode) int {
	abs := func(a, b int) int {
		if a >= b {
			return a - b
		}

		return b - a
	}

	var leftMost, rightMost, minDiff func(t *TreeNode) int

	leftMost = func(t *TreeNode) int {
		// assuming t is not nil
		if t.Left != nil {
			return leftMost(t.Left)
		}

		return t.Val
	}

	rightMost = func(t *TreeNode) int {
		// assuming t is not nil
		if t.Right != nil {
			return rightMost(t.Right)
		}

		return t.Val
	}

	minDiff = func(t *TreeNode) int {
		var left int
		var right int

		if t.Left != nil {
			right = abs(t.Val, rightMost(t.Left))
		} else {
			right = math.MaxInt32
		}

		if t.Right != nil {
			left = abs(t.Val, leftMost(t.Right))
		} else {
			left = math.MaxInt32
		}

		var currentMin = left
		if left > right {
			currentMin = right
		}

		if currentMin == 0 {
			return currentMin
		}

		if t.Left != nil {
			minLeft := minDiff(t.Left)
			if currentMin > minLeft {
				currentMin = minLeft
			}
		}

		if t.Right != nil {
			minRight := minDiff(t.Right)
			if currentMin > minRight {
				currentMin = minRight
			}
		}

		return currentMin
	}

	return minDiff(root)
}
*/
