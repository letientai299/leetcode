package main

import "fmt"

/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 *
 * https://leetcode.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (47.74%)
 * Total Accepted:    78.8K
 * Total Submissions: 152K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * Given the root of a binary tree, return all duplicate subtrees.
 *
 * For each kind of duplicate subtrees, you only need to return the root node
 * of any one of them.
 *
 * Two trees are duplicate if they have the same structure with the same node
 * values.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,4,null,2,4,null,null,4]
 * Output: [[2,4],[4]]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [2,1,1]
 * Output: [[1]]
 *
 *
 * Example 3:
 *
 *
 * Input: root = [2,2,2,3,null,3,null]
 * Output: [[2,3],[3]]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of the nodes in the tree will be in the range [1, 10^4]
 * -200 <= Node.val <= 200
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var height func(n *TreeNode) int

	m := make(map[[3]int][]*TreeNode)
	height = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := height(n.Left)
		r := height(n.Right)

		k := [3]int{n.Val, l, r}
		m[k] = append(m[k], n)

		if l < r {
			l = r
		}

		l++
		return 1 + l
	}

	height(root)

	for _, ts := range m {
		r := len(res)
		fmt.Println(ts)
		seen := make(map[int]struct{})
		for i := 0; i < len(ts)-1; i++ {
			if _, ok := seen[i]; ok {
				continue
			}

			for j := i + 1; j < len(ts); j++ {
				if _, ok := seen[j]; ok {
					continue
				}

				if !isSameTree(ts[i], ts[j]) {
					continue
				}

				seen[i] = struct{}{}
				seen[j] = struct{}{}

				ok := true
				for k := r; ok && k < len(res); k++ {
					if isSameTree(ts[i], res[k]) {
						ok = false
					}
				}

				if ok {
					res = append(res, ts[i])
				}
				break
			}
		}
	}
	return res
}
