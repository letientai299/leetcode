package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (45.27%)
 * Total Accepted:    214.8K
 * Total Submissions: 474.4K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	prefix := fmt.Sprintf("%d->", root.Val)
	left := binaryTreePaths(root.Left)
	for i := range left {
		left[i] = prefix + left[i]
	}

	right := binaryTreePaths(root.Right)
	for i := range right {
		right[i] = prefix + right[i]
	}

	left = append(left, right...)
	return left
}
