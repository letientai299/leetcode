package main

/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (41.06%)
 * Total Accepted:    556.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * Given a binary tree, find the lowest common ancestor (LCA) of two given
 * nodes in the tree.
 *
 * According to the definition of LCA on Wikipedia: “The lowest common ancestor
 * is defined between two nodes p and q as the lowest node in T that has both p
 * and q as descendants (where we allow a node to be a descendant of
 * itself).”
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * Output: 3
 * Explanation: The LCA of nodes 5 and 1 is 3.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * Output: 5
 * Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant
 * of itself according to the LCA definition.
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1,2], p = 1, q = 2
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [2, 10^5].
 * -10^9 <= Node.val <= 10^9
 * All Node.val are unique.
 * p != q
 * p and q will exist in the tree.
 *
 *
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == nil || q == nil {
		return nil
	}

	if p.Val == q.Val {
		return p
	}

	var find func(path []*TreeNode, n *TreeNode) []*TreeNode
	find = func(path []*TreeNode, n *TreeNode) []*TreeNode {
		last := path[len(path)-1]
		if last.Val == n.Val {
			return path
		}

		if last.Left == nil && last.Right == nil {
			return nil
		}

		l := len(path)

		if last.Left != nil {
			path = append(path, last.Left)
			if r := find(path, n); len(r) != 0 {
				return r
			}
			path = path[:l]
		}

		if last.Right != nil {
			path = append(path, last.Right)
			if r := find(path, n); len(r) != 0 {
				return r
			}
			path = path[:l]
		}

		return nil
	}

	ps := find([]*TreeNode{root}, p)
	if len(ps) == 0 {
		return nil
	}

	qs := find([]*TreeNode{root}, q)
	if len(qs) == 0 {
		return nil
	}

	i := 0
	for ; i < len(ps) && i < len(qs); i++ {
		if ps[i].Val != qs[i].Val {
			break
		}
	}
	i--

	return ps[i]
}
