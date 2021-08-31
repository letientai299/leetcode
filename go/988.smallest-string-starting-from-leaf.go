package main

// Smallest String Starting From Leaf
//
// Medium
//
// https://leetcode.com/problems/smallest-string-starting-from-leaf/
//
// You are given the `root` of a binary tree where each node has a value in the
// range `[0, 25]` representing the letters `'a'` to `'z'`.
//
// Return _the **lexicographically smallest** string that starts at a leaf of
// this tree and ends at the root_.
//
// As a reminder, any shorter prefix of a string is **lexicographically
// smaller**.
//
// - For example, `"ab"` is lexicographically smaller than `"aba"`.
//
// A leaf of a node is a node that has no children.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/01/30/tree1.png)
//
// ```
// Input: root = [0,1,2,3,4,3,4]
// Output: "dba"
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2019/01/30/tree2.png)
//
// ```
// Input: root = [25,1,3,1,3,0,2]
// Output: "adz"
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2019/02/01/tree3.png)
//
// ```
// Input: root = [2,2,1,null,1,0,null,0]
// Output: "abc"
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[1, 8500]`.
// - `0 <= Node.val <= 25`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	best := ""
	if root == nil {
		return best
	}

	var cur []byte
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}

		cur = append(cur, byte(n.Val+'a'))

		if n.Left == nil && n.Right == nil {
			bs := make([]byte, len(cur))
			for i, l := 0, len(cur); i < l; i++ {
				bs[l-1-i] = cur[i]
			}

			s := string(bs)
			if best == "" || s < best {
				best = s
			}
			cur = cur[:len(cur)-1]
			return
		}

		walk(n.Left)
		walk(n.Right)
		cur = cur[:len(cur)-1]
	}

	walk(root)
	return best
}
