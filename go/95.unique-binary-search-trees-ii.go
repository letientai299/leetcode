package main

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (37.63%)
 * Total Accepted:    211.6K
 * Total Submissions: 507.2K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 8
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	mem := make(map[int][]*TreeNode)

	var gen func(l, r int) []*TreeNode
	gen = func(l, r int) []*TreeNode {
		if l > r {
			return []*TreeNode{nil}
		}

		if l == r {
			return []*TreeNode{{Val: l}}
		}

		if know, ok := mem[l*1337+r]; ok {
			return know
		}

		res := make([]*TreeNode, 0, n*n*n)
		for i := l; i <= r; i++ {
			left := gen(l, i-1)
			right := gen(i+1, r)
			for _, a := range left {
				for _, b := range right {
					res = append(res,
						&TreeNode{Val: i, Left: a, Right: b},
					)
				}
			}
		}

		mem[l*1337+r] = res
		return res
	}

	return gen(1, n)
}
