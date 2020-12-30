package main

/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 *
 * https://leetcode.com/problems/most-frequent-subtree-sum/description/
 *
 * algorithms
 * Medium (55.97%)
 * Total Accepted:    78.9K
 * Total Submissions: 134.1K
 * Testcase Example:  '[5,2,-3]'
 *
 *
 * Given the root of a tree, you are asked to find the most frequent subtree
 * sum. The subtree sum of a node is defined as the sum of all the node values
 * formed by the subtree rooted at that node (including the node itself). So
 * what is the most frequent subtree sum value? If there is a tie, return all
 * the values with the highest frequency in any order.
 *
 *
 * Examples 1
 * Input:
 *
 * ⁠ 5
 * ⁠/  \
 * 2   -3
 *
 * return [2, -3, 4], since all the values happen only once, return all of them
 * in any order.
 *
 *
 * Examples 2
 * Input:
 *
 * ⁠ 5
 * ⁠/  \
 * 2   -5
 *
 * return [2], since 2 happens twice, however -5 only occur once.
 *
 *
 * Note:
 * You may assume the sum of values in any subtree is in the range of 32-bit
 * signed integer.
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

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	m := make(map[int]int)
	best := 0
	var treeSum func(n *TreeNode) int
	treeSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		v := n.Val + treeSum(n.Left) + treeSum(n.Right)
		m[v]++
		if m[v] > best {
			best = m[v]
		}
		return v
	}

	treeSum(root)

	var res []int
	for k, v := range m {
		if v == best {
			res = append(res, k)
		}
	}
	return res
}
