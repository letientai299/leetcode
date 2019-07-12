package main

/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
 *
 * https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/description/
 *
 * algorithms
 * Easy (43.30%)
 * Total Accepted:    51.1K
 * Total Submissions: 118K
 * Testcase Example:  '[2,2,5,null,null,5,7]'
 *
 * Given a non-empty special binary tree consisting of nodes with the
 * non-negative value, where each node in this tree has exactly two or zero
 * sub-node. If the node has two sub-nodes, then this node's value is the
 * smaller value among its two sub-nodes. More formally, the property root.val
 * = min(root.left.val, root.right.val) always holds.
 *
 * Given such a binary tree, you need to output the second minimum value in the
 * set made of all the nodes' value in the whole tree.
 *
 * If no such second minimum value exists, output -1 instead.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   5
 * ⁠    / \
 * ⁠   5   7
 *
 * Output: 5
 * Explanation: The smallest value is 2, the second smallest value is 5.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   2
 *
 * Output: -1
 * Explanation: The smallest value is 2, but there isn't any second smallest
 * value.
 *
 *
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
func findSecondMinimumValue(root *TreeNode) int {
	var min func(t *TreeNode, bound int) int
	min = func(t *TreeNode, bound int) int {
		if t == nil {
			return bound
		}

		left := min(t.Left, bound)
		right := min(t.Right, bound)
		res := bound

		if t.Val > bound {
			res = t.Val
		}

		if (left < res || res == bound) && left > bound {
			res = left
		}

		if (right < res || res == bound) && right > bound {
			res = right
		}

		return res
	}

	m1 := min(root, -1)
	m2 := min(root, m1)
	if m1 == m2 {
		return -1
	}

	return m2
}
