package main

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (43.16%)
 * Total Accepted:    373.2K
 * Total Submissions: 774.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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

func pathSum(root *TreeNode, sum int) [][]int {
	var find func(s int, t *TreeNode)

	var path []int
	var res [][]int

	find = func(s int, t *TreeNode) {
		if t == nil {
			return
		}

		if t.Left == nil && t.Right == nil {
			if s+t.Val == sum {
				path = append(path, t.Val)
				p := make([]int, len(path))
				copy(p, path)
				res = append(res, p)
				path = path[:len(path)-1]
			}
			return
		}

		s += t.Val
		l := len(path)
		path = append(path, t.Val)
		find(s, t.Left)
		find(s, t.Right)
		path = path[:l]
	}
	find(0, root)
	return res
}
