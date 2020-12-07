package main

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Easy (42.24%)
 * Total Accepted:    97.7K Total Submissions: 231.3K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
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
func pathSum_437(root *TreeNode, sum int) int {
	var allPaths int
	var scan func(*TreeNode)
	var f func(*TreeNode, int)

	f = func(n *TreeNode, rest int) {
		if n == nil {
			return
		}

		rest -= n.Val

		if rest == 0 {
			allPaths++
		}

		f(n.Left, rest)
		f(n.Right, rest)
	}

	scan = func(n *TreeNode) {
		if n == nil {
			return
		}

		f(n, sum)
		scan(n.Left)
		scan(n.Right)
	}

	scan(root)

	return allPaths
}
