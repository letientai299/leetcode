package main

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (50.50%)
 * Total Accepted:    358.3K
 * Total Submissions: 647.1K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 *
 * Example:
 *
 *
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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

func rightSideView(root *TreeNode) []int {
	var r []int
	if root == nil {
		return r
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		var next []*TreeNode
		ok := false
		for _, n := range q {
			if !ok {
				r = append(r, n.Val)
				ok = true
			}

			if n.Right != nil {
				next = append(next, n.Right)
			}

			if n.Left != nil {
				next = append(next, n.Left)
			}
		}
		q = next
	}
	return r
}
