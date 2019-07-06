package main

/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (52.64%)
 * Total Accepted:    91.8K
 * Total Submissions: 174.3K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given a Binary Search Tree and a target number, return true if there exist
 * two elements in the BST such that their sum is equal to the given target.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 9
 *
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 28
 *
 * Output: False
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
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	var exist func(t *TreeNode, k int) *TreeNode
	exist = func(t *TreeNode, k int) *TreeNode {
		if t == nil {
			return nil
		}

		if t.Val == k {
			return t
		}

		if t.Val < k {
			return exist(t.Right, k)
		}

		return exist(t.Left, k)
	}

	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			res := k - n.Val
			t := exist(root, res)
			if t != nil && t != n {
				return true
			}

			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
	}

	return false
}
