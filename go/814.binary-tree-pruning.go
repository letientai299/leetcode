package main

/*
 * @lc app=leetcode id=814 lang=golang
 *
 * [814] Binary Tree Pruning
 *
 * https://leetcode.com/problems/binary-tree-pruning/description/
 *
 * algorithms
 * Medium (72.41%)
 * Total Accepted:    76.6K
 * Total Submissions: 104.8K
 * Testcase Example:  '[1,null,0,0,1]'
 *
 * We are given the head node root of a binary tree, where additionally every
 * node's value is either a 0 or a 1.
 *
 * Return the same tree where every subtree (of the given tree) not containing
 * a 1 has been removed.
 *
 * (Recall that the subtree of a node X is X, plus every node that is a
 * descendant of X.)
 *
 *
 * Example 1:
 * Input: [1,null,0,0,1]
 * Output: [1,null,0,null,1]
 * ⁠
 * Explanation:
 * Only the red nodes satisfy the property "every subtree not containing a 1".
 * The diagram on the right represents the answer.
 *
 *
 *
 *
 *
 * Example 2:
 * Input: [1,0,1,0,0,0,1]
 * Output: [1,null,1,null,1]
 *
 *
 *
 *
 *
 *
 * Example 3:
 * Input: [1,1,0,1,1,0,1,0]
 * Output: [1,1,0,1,1,null,1]
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The binary tree will have at most 200 nodes.
 * The value of each node will only be 0 or 1.
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

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var prune func(n *TreeNode) (*TreeNode, bool)

	prune = func(n *TreeNode) (*TreeNode, bool) {
		if n == nil {
			return nil, true
		}

		if n.Left == nil && n.Right == nil {
			if n.Val == 0 {
				return nil, true
			}

			return n, false
		}

		_, pruneLeft := prune(n.Left)
		_, pruneRight := prune(n.Right)
		if pruneLeft {
			n.Left = nil
		}

		if pruneRight {
			n.Right = nil
		}

		if n.Val == 0 && pruneLeft && pruneRight {
			return nil, true
		}

		return n, false
	}

	root, _ = prune(root)
	return root
}
