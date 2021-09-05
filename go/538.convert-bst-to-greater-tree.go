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
	// 1038 is same with this. And I found a better solution than this one,
	// after 2 years

	var leftLeaf func(root *TreeNode) *TreeNode
	var plus func(root *TreeNode, n int)

	leftLeaf = func(root *TreeNode) *TreeNode {
		if root.Left == nil {
			return root
		}

		return leftLeaf(root.Left)
	}

	plus = func(root *TreeNode, n int) {
		if root == nil {
			return
		}

		root.Val += n
		plus(root.Left, n)
		plus(root.Right, n)
	}
	if root == nil {
		return nil
	}

	right := convertBST(root.Right)
	if right != nil {
		root.Val += leftLeaf(right).Val
	}
	left := convertBST(root.Left)
	plus(left, root.Val)
	return root
}
