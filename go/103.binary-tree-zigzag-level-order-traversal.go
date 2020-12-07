package main

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (44.15%)
 * Total Accepted:    444.4K
 * Total Submissions: 899K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
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
 * return its zigzag level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 16)
	q := []*TreeNode{root}
	res = append(res, []int{root.Val})
	reverse := true
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
		if reverse {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-1-i] = level[len(level)-1-i], level[i]
			}
		}
		reverse = !reverse
		res = append(res, level)
	}
	return res[:len(res)-1]
}
