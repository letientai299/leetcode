package main

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 *
 * https://leetcode.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (50.96%)
 * Total Accepted:    78.7K
 * Total Submissions: 154.5K
 * Testcase Example:  '[5,2,13]'
 *
 * Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 * every key of the original BST is changed to the original key plus sum of all
 * keys greater than the original key in BST.
 *
 *
 * Example:
 *
 * Input: The root of a Binary Search Tree like this:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * Output: The root of a Greater Tree like this:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
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
func convertBST(root *TreeNode) *TreeNode {
	nodes := make(map[int]*TreeNode)
	var cache []int
	var travel func(t *TreeNode)

	travel = func(t *TreeNode) {
		if t == nil {
			return
		}

		travel(t.Left)
		cache = append(cache, t.Val)
		nodes[t.Val] = t
		travel(t.Right)
	}

	travel(root)

	sum := 0
	for i := len(cache) - 1; i >= 0; i-- {
		nodes[cache[i]].Val += sum
		sum += cache[i]
	}
	return root
}
