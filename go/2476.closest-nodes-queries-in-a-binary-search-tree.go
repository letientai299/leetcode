package main

import "sort"

// Closest Nodes Queries in a Binary Search Tree
//
// Medium
//
// https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree
//
// You are given the `root` of a **binary search tree** and an array `queries`
// of size `n` consisting of positive integers.
//
// Find a **2D** array `answer` of size `n` where `answer[i] = [mini, maxi]`:
//
// - `mini` is the **largest** value in the tree that is smaller than or equal
// to `queries[i]`. If a such value does not exist, add `-1` instead.
// - `maxi` is the **smallest** value in the tree that is greater than or equal
// to `queries[i]`. If a such value does not exist, add `-1` instead.
//
// Return _the array_ `answer`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2022/09/28/bstreeedrawioo.png)
//
// ```
// Input: root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries =
// [2,5,16]
// Output: [[2,2],[4,6],[15,-1]]
// Explanation: We answer the queries in the following way:
// - The largest number that is smaller or equal than 2 in the tree is 2, and
// the smallest number that is greater or equal than 2 is still 2. So the answer
// for the first query is [2,2].
// - The largest number that is smaller or equal than 5 in the tree is 4, and
// the smallest number that is greater or equal than 5 is 6. So the answer for
// the second query is [4,6].
// - The largest number that is smaller or equal than 16 in the tree is 15, and
// the smallest number that is greater or equal than 16 does not exist. So the
// answer for the third query is [15,-1].
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2022/09/28/bstttreee.png)
//
// ```
// Input: root = [4,null,9], queries = [3]
// Output: [[-1,4]]
// Explanation: The largest number that is smaller or equal to 3 in the tree
// does not exist, and the smallest number that is greater or equal to 3 is 4.
// So the answer for the query is [-1,4].
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[2, 105]`.
// - `1 <= Node.val <= 106`
// - `n == queries.length`
// - `1 <= n <= 105`
// - `1 <= queries[i] <= 106`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	var all []int
	var walk func(root *TreeNode)
	walk = func(n *TreeNode) {
		if n != nil {
			walk(n.Left)
			all = append(all, n.Val)
			walk(n.Right)
		}
	}

	walk(root)
	res := make([][]int, len(queries))
	for i, v := range queries {
		j := sort.SearchInts(all, v)
		if j == len(all) {
			res[i] = []int{all[j-1], -1}
			continue
		}

		if all[j] == v {
			res[i] = []int{all[j], all[j]}
			continue
		}

		if j == 0 {
			res[i] = []int{-1, all[j]}
			continue
		}

		if all[j-1] == v {
			res[i] = []int{all[j-1], all[j-1]}
			continue
		}

		res[i] = []int{all[j-1], all[j]}
	}

	return res
}
