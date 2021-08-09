package main

// Deepest Leaves Sum
//
// Medium
//
// https://leetcode.com/problems/deepest-leaves-sum/
//
// Given the `root` of a binary tree, return _the sum of values of its deepest
// leaves_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/07/31/1483_ex1.png)
//
// ```
// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// Output: 15
//
// ```
//
// **Example 2:**
//
// ```
// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 19
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 104]`.
// - `1 <= Node.val <= 100`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := 0
		next := make([]*TreeNode, 0, len(q)*2)
		for _, n := range q {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}

		if len(next) == 0 {
			return sum
		}

		q = next
	}

	return 0
}
