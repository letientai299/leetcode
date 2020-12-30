package main

/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 *
 * https://leetcode.com/problems/maximum-width-of-binary-tree/description/
 *
 * algorithms
 * Medium (39.57%)
 * Total Accepted:    102.5K
 * Total Submissions: 256K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * Given a binary tree, write a function to get the maximum width of the given
 * tree. The maximum width of a tree is the maximum width among all levels.
 *
 * The width of one level is defined as the length between the end-nodes (the
 * leftmost and right most non-null nodes in the level, where the null nodes
 * between the end-nodes are also counted into the length calculation.
 *
 * It is guaranteed that the answer will in the range of 32-bit signed
 * integer.
 *
 * Example 1:
 *
 *
 * Input:
 *
 * ⁠          1
 * ⁠        /   \
 * ⁠       3     2
 * ⁠      / \     \
 * ⁠     5   3     9
 *
 * Output: 4
 * Explanation: The maximum width existing in the third level with the length 4
 * (5,3,null,9).
 *
 *
 * Example 2:
 *
 *
 * Input:
 *
 * ⁠         1
 * ⁠        /
 * ⁠       3
 * ⁠      / \
 * ⁠     5   3
 *
 * Output: 2
 * Explanation: The maximum width existing in the third level with the length 2
 * (5,3).
 *
 *
 * Example 3:
 *
 *
 * Input:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /
 * ⁠     5
 *
 * Output: 2
 * Explanation: The maximum width existing in the second level with the length
 * 2 (3,2).
 *
 *
 * Example 4:
 *
 *
 * Input:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /     \
 * ⁠     5       9
 * ⁠    /         \
 * ⁠   6           7
 * Output: 8
 * Explanation:The maximum width existing in the fourth level with the length 8
 * (6,null,null,null,null,null,null,7).
 *
 *
 *
 * Constraints:
 *
 *
 * The given binary tree will have between 1 and 3000 nodes.
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
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	best := 1
	row := []*TreeNode{root}
	col := []int{0}
	for len(row) != 0 {
		next := make([]*TreeNode, 0, len(row)*2)
		nextCol := make([]int, 0, len(row)*2)
		for i, n := range row {
			c := col[i]

			if n.Left != nil {
				next = append(next, n.Left)
				nextCol = append(nextCol, 2*c)
			}

			if n.Right != nil {
				next = append(next, n.Right)
				nextCol = append(nextCol, 2*c+1)
			}
		}

		row = next
		col = nextCol
		if len(col) >= 2 {
			v := col[len(col)-1] - col[0] + 1
			if v > best {
				best = v
			}
		}
	}

	return best
}
