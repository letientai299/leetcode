package main

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 *
 * https://leetcode.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (46.76%)
 * Total Accepted:    131.5K
 * Total Submissions: 281.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 *
 * Given a binary tree, you need to compute the length of the diameter of the
 * tree. The diameter of a binary tree is the length of the longest path
 * between any two nodes in a tree. This path may or may not pass through the
 * root.
 *
 *
 *
 * Example:
 * Given a binary tree
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 *
 * Note:
 * The length of path between two nodes is represented by the number of edges
 * between them.
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
func diameterOfBinaryTree(root *TreeNode) int {
	var max func(arr ...int) int
	max = func(arr ...int) int {
		m := 0
		for _, n := range arr {
			if m < n {
				m = n
			}
		}
		return m
	}

	var depth func(t *TreeNode, maxDiam int) (int, int)
	depth = func(t *TreeNode, maxDiam int) (int, int) {
		if t == nil {
			return 0, maxDiam
		}

		l, ld := depth(t.Left, maxDiam)
		r, rd := depth(t.Right, maxDiam)
		return max(l, r) + 1, max(ld, rd, l+r, maxDiam)
	}

	_, diam := depth(root, 0)
	return diam
}
