package main

// Delete Nodes And Return Forest
//
// Medium
//
// https://leetcode.com/problems/delete-nodes-and-return-forest/
//
// Given the `root` of a binary tree, each node in the tree has a distinct
// value.
//
// After deleting all nodes with a value in `to_delete`, we are left with a
// forest (a disjoint union of trees).
//
// Return the roots of the trees in the remaining forest. You may return the
// result in any order.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png)
//
// ```
// Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
// Output: [[1,2,null,4],[6],[7]]
//
// ```
//
// **Example 2:**
//
// ```
// Input: root = [1,2,4,null,3], to_delete = [3]
// Output: [[1,2,4]]
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the given tree is at most `1000`.
// - Each node has a distinct value between `1` and `1000`.
// - `to_delete.length <= 1000`
// - `to_delete` contains distinct values between `1` and `1000`.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	m := make(map[int]bool)
	for _, v := range toDelete {
		m[v] = true
	}

	var res []*TreeNode

	var walk func(n, parent *TreeNode)
	walk = func(n, parent *TreeNode) {
		if n == nil {
			return
		}

		left, right := n.Left, n.Right
		if m[n.Val] {
			if parent == nil {
				// do nothing
			} else if n == parent.Left {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			n = nil
		} else {
			if parent == nil {
				res = append(res, n)
			}
		}

		walk(left, n)
		walk(right, n)
	}

	walk(root, nil)
	return res
}
