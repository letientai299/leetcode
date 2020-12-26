package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (41.97%)
 * Total Accepted:    270.4K
 * Total Submissions: 553.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var build func(lp, rp, li, ri int) *TreeNode

	build = func(lp, rp, li, ri int) *TreeNode {
		if lp > rp {
			return nil
		}

		root := &TreeNode{Val: postorder[rp]}
		if lp == rp {
			return root
		}

		mid := li
		for mid < len(inorder) && inorder[mid] != postorder[rp] {
			mid++
		}

		root.Left = build(lp, lp+mid-li-1, li, mid-1)
		root.Right = build(lp+mid-li, rp-1, mid+1, ri)
		return root
	}

	return build(0, len(postorder)-1, 0, len(inorder)-1)
}
