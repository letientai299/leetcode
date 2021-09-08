package main

import "fmt"

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 *
 * https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (45.20%)
 * Total Accepted:    394.5K
 * Total Submissions: 774.3K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * Given a binary tree, flatten it to a linked list in-place.
 *
 * For example, given the following tree:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 *
 * The flattened tree should look like:
 *
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
func flatten114(root *TreeNode) {
	var flat func(node *TreeNode) (head, tail *TreeNode)
	flat = func(node *TreeNode) (head, tail *TreeNode) {
		if node == nil {
			return nil, nil
		}
		fmt.Println(node)

		if node.Left == nil && node.Right == nil {
			return node, node
		}

		if node.Left == nil {
			r, t := flat(node.Right)
			node.Right = r
			return node, t
		}

		l, tl := flat(node.Left)
		node.Left = nil
		if node.Right == nil {
			node.Right = l
			return node, tl
		}

		r, tr := flat(node.Right)
		node.Right = l
		tl.Right = r
		return node, tr
	}

	flat(root)
}
