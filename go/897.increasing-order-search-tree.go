package main

/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 *
 * https://leetcode.com/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Easy (66.37%)
 * Total Accepted:    79.7K
 * Total Submissions: 110K
 * Testcase Example:  '[5,3,6,2,4,null,8,1,null,null,null,7,9]\r'
 *
 * Given a binary search tree, rearrange the tree in in-order so that the
 * leftmost node in the tree is now the root of the tree, and every node has no
 * left child and only 1 right child.
 *
 *
 * Example 1:
 * Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]
 *
 * ⁠      5
 * ⁠     / \
 * ⁠   3    6
 * ⁠  / \    \
 * ⁠ 2   4    8
 * /        / \
 * 1        7   9
 *
 * Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 * ⁠1
 * \
 * 2
 * \
 * 3
 * \
 * 4
 * \
 * 5
 * \
 * 6
 * \
 * 7
 * \
 * 8
 * \
 * ⁠                9
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree will be between 1 and 100.
 * Each node will have a unique integer value from 0 to 1000.
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
func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	p := head

	var travel func(n *TreeNode)
	travel = func(n *TreeNode) {
		if n == nil {
			return
		}
		travel(n.Left)
		n.Left = nil
		p.Right = n
		right := n.Right
		n.Right = nil
		p = p.Right
		travel(right)
	}

	travel(root)
	return head.Right
}
