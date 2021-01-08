package main

/*
 * @lc app=leetcode id=863 lang=golang
 *
 * [863] All Nodes Distance K in Binary Tree
 *
 * https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/
 *
 * algorithms
 * Medium (51.10%)
 * Total Accepted:    110.5K
 * Total Submissions: 192.6K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n2'
 *
 * We are given a binary tree (with root node root), a target node, and an
 * integer value K.
 *
 * Return a list of the values of all nodes that have a distance K from the
 * target node.  The answer can be returned in any order.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
 *
 * Output: [7,4,1]
 *
 * Explanation:
 * The nodes that are a distance 2 from the target node (with value 5)
 * have values 7, 4, and 1.
 *
 *
 *
 * Note that the inputs "root" and "target" are actually TreeNodes.
 * The descriptions of the inputs above are just serializations of these
 * objects.
 *
 *
 *
 *
 * Note:
 *
 *
 * The given tree is non-empty.
 * Each node in the tree has unique values 0 <= node.val <= 500.
 * The target node is a node in the tree.
 * 0 <= K <= 1000.
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	var findParents func(*TreeNode, int, []*TreeNode) []*TreeNode
	findParents = func(n *TreeNode, v int, cur []*TreeNode) []*TreeNode {
		if n == nil {
			return nil
		}

		cur = append(cur, n)
		if n.Val == v {
			return cur
		}

		l := len(cur)
		if s := findParents(n.Left, v, cur); s != nil {
			return s
		}

		cur = cur[:l]
		return findParents(n.Right, v, cur)
	}
	chains := findParents(root, target.Val, []*TreeNode{})

	if len(chains) == 0 {
		return nil
	}

	var res []int
	visited := make(map[int]struct{})
	var findHeight func(n *TreeNode, h int)

	findHeight = func(n *TreeNode, h int) {
		if n == nil || h < 0 {
			return
		}

		if _, ok := visited[n.Val]; ok {
			return
		}

		visited[n.Val] = struct{}{}
		if h == 0 {
			res = append(res, n.Val)
			return
		}
		findHeight(n.Left, h-1)
		findHeight(n.Right, h-1)
	}

	l := len(chains)
	for i := l - 1; i >= l-K-1 && i >= 0; i-- {
		findHeight(chains[i], K-(l-i-1))
	}

	return res
}
