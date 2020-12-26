package main

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (44.32%)
 * Total Accepted:    426K
 * Total Submissions: 840.8K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder= [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7

			1
		   3
      5 6

pre: 1 3 5 6
in:  1 5 3 6
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(lp, rp, li, ri int) *TreeNode

	build = func(lp, rp, li, ri int) *TreeNode {
		if lp > rp {
			return nil
		}

		root := &TreeNode{Val: preorder[lp]}
		if lp == rp {
			return root
		}

		mid := li
		for mid < len(inorder) && inorder[mid] != preorder[lp] {
			mid++
		}

		root.Left = build(lp+1, lp+mid-li, li, mid-1)
		root.Right = build(lp+1+mid-li, rp, mid+1, ri)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
