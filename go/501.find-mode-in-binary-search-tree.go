package main

import "math"

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 *
 * https://leetcode.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (39.15%)
 * Total Accepted:    53.6K
 * Total Submissions: 136.8K
 * Testcase Example:  '[1,null,2,2]'
 *
 * Given a binary search tree (BST) with duplicates, find all the mode(s) (the
 * most frequently occurred element) in the given BST.
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than or equal
 * to the node's key.
 * The right subtree of a node contains only nodes with keys greater than or
 * equal to the node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * For example:
 * Given BST [1,null,2,2],
 *
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 *
 *
 * return [2].
 *
 * Note: If a tree has more than one mode, you can return them in any order.
 *
 * Follow up: Could you do that without using any extra space? (Assume that the
 * implicit stack space incurred due to recursion does not count).
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
func findMode(root *TreeNode) []int {
	// map of value in the binary and its number of appearance
	counters := make(map[int]int)
	mode := int(math.MinInt32)

	var travel func(node *TreeNode)

	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		counters[node.Val]++
		if mode < counters[node.Val] {
			mode = counters[node.Val]
		}

		travel(node.Left)
		travel(node.Right)
	}

	travel(root)

	var res []int

	for k, v := range counters {
		if v == mode {
			res = append(res, k)
		}
	}

	return res
}
