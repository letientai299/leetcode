package main

/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 *
 * https://leetcode.com/problems/cousins-in-binary-tree/description/
 *
 * algorithms
 * Easy (51.93%)
 * Total Accepted:    124.6K
 * Total Submissions: 239.2K
 * Testcase Example:  '[1,2,3,4]\n4\n3'
 *
 * In a binary tree, the root node is at depth 0, and children of each depth k
 * node are at depth k+1.
 *
 * Two nodes of a binary tree are cousins if they have the same depth, but have
 * different parents.
 *
 * We are given the root of a binary tree with unique values, and the values x
 * and y of two different nodes in the tree.
 *
 * Return true if and only if the nodes corresponding to the values x and y are
 * cousins.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: root = [1,2,3,4], x = 4, y = 3
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 *
 * Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: root = [1,2,3,null,4], x = 2, y = 3
 * Output: false
 *
 *
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree will be between 2 and 100.
 * Each node has a unique integer value from 1 to 100.
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
func isCousins(root *TreeNode, x int, y int) bool {
	var find func(n *TreeNode, v int, d int) (parent *TreeNode, node *TreeNode, dep int)

	find = func(n *TreeNode, v int, d int) (parent *TreeNode, node *TreeNode, dep int) {
		if n == nil {
			return nil, nil, d
		}

		if n.Left != nil && n.Left.Val == v {
			return n, n.Left, d
		}

		if n.Right != nil && n.Right.Val == v {
			return n, n.Right, d
		}

		if n.Left != nil {
			p, c, dl := find(n.Left, v, d+1)
			if c != nil {
				return p, c, dl
			}
		}

		if n.Right != nil {
			p, c, dr := find(n.Right, v, d+1)
			if c != nil {
				return p, c, dr
			}
		}

		return nil, nil, d
	}

	px, nx, dx := find(root, x, 0)
	py, ny, dy := find(root, y, 0)
	if px == py || nx == nil || ny == nil || dx != dy {
		return false
	}

	return true
}
