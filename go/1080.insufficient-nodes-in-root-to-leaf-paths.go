package main

// Insufficient Nodes in Root to Leaf Paths
//
// Medium
//
// https://leetcode.com/problems/insufficient-nodes-in-root-to-leaf-paths/
//
// Given the `root` of a binary tree and an integer `limit`, delete all
// **insufficient nodes** in the tree simultaneously, and return _the root of
// the resulting binary tree_.
//
// A node is **insufficient** if every root to **leaf** path intersecting this
// node has a sum strictly less than `limit`.
//
// A **leaf** is a node with no children.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/06/05/insufficient-11.png)
//
// ```
// Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
// Output: [1,2,3,4,null,null,7,8,9,null,14]
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2019/06/05/insufficient-3.png)
//
// ```
// Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
// Output: [5,4,8,11,null,17,4,7,null,null,null,5]
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2019/06/11/screen-shot-2019-06-11-at-83301-pm.png)
//
// ```
// Input: root = [1,2,-3,-5,null,4,null], limit = -1
// Output: [1,null,-3,4]
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 5000]`.
// - `-105 <= Node.val <= 105`
// - `-109 <= limit <= 109`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var walk func(n, parent *TreeNode, isLeft bool, sum int) []int
	walk = func(n, parent *TreeNode, isLeft bool, sum int) []int {
		if n == nil {
			return nil
		}

		sum += n.Val

		if n.Left == nil && n.Right == nil {
			if sum < limit {
				if isLeft {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}
			return []int{sum}
		}

		l := walk(n.Left, n, true, sum)
		r := walk(n.Right, n, false, sum)
		l = append(l, r...)

		ok := false
		for _, v := range l {
			if v >= limit {
				ok = true
				break
			}
		}

		if !ok {
			if isLeft {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}

		return l
	}

	pre := &TreeNode{Left: root}
	walk(root, pre, true, 0)
	return pre.Left
}
