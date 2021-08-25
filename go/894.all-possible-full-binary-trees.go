package main

// All Possible Full Binary Trees
//
// Medium
//
// https://leetcode.com/problems/all-possible-full-binary-trees/
//
// Given an integer `n`, return _a list of all possible **full binary trees**
// with_ `n` _nodes_. Each node of each tree in the answer must have `Node.val
// == 0`.
//
// Each element of the answer is the root node of one possible tree. You may
// return the final list of trees in **any order**.
//
// A **full binary tree** is a binary tree where each node has exactly `0` or
// `2` children.
//
// **Example 1:**
//
// ![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/22/fivetrees.png)
//
// ```
// Input: n = 7
// Output:
// [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 3
// Output: [[0,0,0]]
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 20`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
	var build func(n int) []*TreeNode

	mem := make([][]*TreeNode, 20) // input cap
	mem[1] = []*TreeNode{{}}

	build = func(n int) []*TreeNode {
		if n%2 == 0 {
			return nil // can't create complete tree with n is even
		}

		if mem[n] != nil {
			return mem[n]
		}

		var res []*TreeNode

		for i := 1; i < n-1; i += 2 {
			//   7         5      3
			// 1   5     1   3  1   1
			// 3   3     3   1
			// 5   1
			l, r := i, n-i-1
			lefts := build(l)
			rights := build(r)

			for _, lt := range lefts {
				for _, rt := range rights {
					res = append(res, &TreeNode{Left: lt, Right: rt})
				}
			}
		}

		mem[n] = res
		return res
	}

	return build(n)
}
