package main

import "math"

// Maximum Level Sum of a Binary Tree
//
// Medium
//
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
//
// Given the `root` of a binary tree, the level of its root is `1`, the level of
// its children is `2`, and so on.
//
// Return the **smallest** level `x` such that the sum of all the values of
// nodes at level `x` is **maximal**.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/05/03/capture.JPG)
//
// ```
// Input: root = [1,7,0,7,-8,null,null]
// Output: 2
// Explanation:
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
// Output: 2
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 104]`.
// - `-105 <= Node.val <= 105`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	maxSum := math.MinInt32
	minLevel := 1

	level := 0
	q := []*TreeNode{root}
	for len(q) != 0 {
		level++
		sum := 0
		next := make([]*TreeNode, 0, len(q)*2)
		for _, n := range q {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			minLevel = level
		}

		q = next
	}

	return minLevel
}
