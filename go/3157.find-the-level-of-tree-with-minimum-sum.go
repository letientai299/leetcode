package main

import (
	"math"
)

// Find the Level of Tree with Minimum Sum
//
// # Medium
//
// https://leetcode.com/problems/find-the-level-of-tree-with-minimum-sum/
func minimumLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	lvl := 0
	best := 0
	bestSum := math.MaxInt
	var next []*TreeNode
	for len(queue) > 0 {
		lvl++

		s := 0
		for _, node := range queue {
			s += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if s < bestSum {
			bestSum = s
			best = lvl
		}

		queue, next = next, queue[:0]
	}

	return best
}
