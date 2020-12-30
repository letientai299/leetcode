package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (51.04%)
 * Total Accepted:    721.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
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

func levelOrder102(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 16)
	q := []*TreeNode{root}
	res = append(res, []int{root.Val})
	for len(q) > 0 {
		var level []int
		var next []*TreeNode
		for _, r := range q {
			if r.Left != nil {
				level = append(level, r.Left.Val)
				next = append(next, r.Left)
			}

			if r.Right != nil {
				level = append(level, r.Right.Val)
				next = append(next, r.Right)
			}
		}
		q = next
		res = append(res, level)
	}
	return res[:len(res)-1]
}
