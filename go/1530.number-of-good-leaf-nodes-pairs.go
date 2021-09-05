package main

// Number of Good Leaf Nodes Pairs
//
// Medium
//
// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/
//
// Given the `root` of a binary tree and an integer `distance`. A pair of two
// different **leaf** nodes of a binary tree is said to be good if the length of
// **the shortest path** between them is less than or equal to `distance`.
//
// Return _the number of good leaf node pairs_ in the tree.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/07/09/e1.jpg)
//
// ```
// Input: root = [1,2,3,null,4], distance = 3
// Output: 1
// Explanation: The leaf nodes of the tree are 3 and 4 and the length of the
// shortest path between them is 3. This is the only good pair.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/07/09/e2.jpg)
//
// ```
// Input: root = [1,2,3,4,5,6,7], distance = 3
// Output: 2
// Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The
// pair [4,6] is not good because the length of ther shortest path between them
// is 4.
//
// ```
//
// **Example 3:**
//
// ```
// Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
// Output: 1
// Explanation: The only good pair is [2,5].
//
// ```
//
// **Example 4:**
//
// ```
// Input: root = [100], distance = 1
// Output: 0
//
// ```
//
// **Example 5:**
//
// ```
// Input: root = [1,1,1], distance = 2
// Output: 1
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the `tree` is in the range `[1, 2^10].`
// - Each node's value is between `[1, 100]`.
// - `1 <= distance <= 10`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
	leafMap := map[int]int{0: 1}

	var numPairs func(n *TreeNode) (int, map[int]int)
	numPairs = func(n *TreeNode) (int, map[int]int) {
		if n == nil {
			return 0, nil
		}

		if n.Left == nil && n.Right == nil {
			return 0, leafMap
		}

		leftPairs, leftLeaves := numPairs(n.Left)
		rightPairs, rightLeaves := numPairs(n.Right)

		total := leftPairs + rightPairs

		leaves := make(map[int]int, len(leftLeaves) + len(rightLeaves))

		if leftLeaves == nil || rightLeaves == nil {
			if leftLeaves == nil {
				leftLeaves = rightLeaves
			}

			for d, c := range leftLeaves {
				leaves[d+1] = c
			}

			return total, leaves
		}

		for depRight, countRight := range rightLeaves {
			leaves[depRight+1] += countRight
		}

		for depLeft, countLeft := range leftLeaves {
			leaves[depLeft+1] += countLeft

			for depRight, countRight := range rightLeaves {
				if depLeft+depRight+2 <= distance {
					total += countLeft * countRight
				}
			}
		}

		return total, leaves
	}

	pairs, _ := numPairs(root)
	return pairs
}
