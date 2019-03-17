package main

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (45.80%)
 * Total Accepted:    212.3K
 * Total Submissions: 463.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var levels [][]int
	var level []int
	for {
		length := len(nodes)
		if length == 0 {
			break
		}

		level = make([]int, length)

		for i := 0; i < length; i++ {
			n := nodes[i]
			level[i] = n.Val
			if n.Left != nil {
				nodes = append(nodes, n.Left)
			}
			if n.Right != nil {
				nodes = append(nodes, n.Right)
			}
		}
		levels = append(levels, level)
		nodes = nodes[length:]
	}

	n := len(levels)
	for i := 0; i < n/2; i++ {
		left := levels[i]
		right := levels[n-1-i]

		levels[i] = right
		levels[n-1-i] = left
	}
	return levels
}
