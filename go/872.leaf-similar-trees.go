package main

/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 *
 * https://leetcode.com/problems/leaf-similar-trees/description/
 *
 * algorithms
 * Easy (64.27%)
 * Total Accepted:    105.3K
 * Total Submissions: 163.2K
 * Testcase Example:  '[3,5,1,6,2,9,8,null,null,7,4]\n' +
  '[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]'
 *
 * Consider all the leaves of a binary tree, from left to right order, the
 * values of those leaves form a leaf value sequence.
 *
 *
 *
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4,
 * 9, 8).
 *
 * Two binary trees are considered leaf-similar if their leaf value sequence is
 * the same.
 *
 * Return true if and only if the two given trees with head nodes root1 and
 * root2 are leaf-similar.
 *
 *
 * Example 1:
 *
 *
 * Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 =
 * [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root1 = [1], root2 = [1]
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: root1 = [1], root2 = [2]
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: root1 = [1,2], root2 = [2,2]
 * Output: true
 *
 *
 * Example 5:
 *
 *
 * Input: root1 = [1,2,3], root2 = [1,3,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in each tree will be in the range [1, 200].
 * Both of the given trees will have values in the range [0, 200].
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

func leafSimilar(a *TreeNode, b *TreeNode) bool {
	leaves := func(n *TreeNode) []int {
		if n == nil {
			return nil
		}

		stack := []*TreeNode{n}
		var l []int
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if node.Left == nil && node.Right == nil {
				l = append(l, node.Val)
				continue
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
		return l
	}

	la := leaves(a)
	lb := leaves(b)

	if len(la) != len(lb) {
		return false
	}

	for i, x := range la {
		y := lb[i]
		if x != y {
			return false
		}
	}

	return true
}
