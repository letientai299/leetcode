package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=655 lang=golang
 *
 * [655] Print Binary Tree
 *
 * https://leetcode.com/problems/print-binary-tree/description/
 *
 * algorithms
 * Medium (52.97%)
 * Total Accepted:    36.1K
 * Total Submissions: 64.9K
 * Testcase Example:  '[1,2]'
 *
 * Print a binary tree in an m*n 2D string array following these rules:
 *
 *
 * The row number m should be equal to the height of the given binary tree.
 * The column number n should always be an odd number.
 * The root node's value (in string format) should be put in the exactly middle
 * of the first row it can be put. The column and the row where the root node
 * belongs will separate the rest space into two parts (left-bottom part and
 * right-bottom part). You should print the left subtree in the left-bottom
 * part and print the right subtree in the right-bottom part. The left-bottom
 * part and the right-bottom part should have the same size. Even if one
 * subtree is none while the other is not, you don't need to print anything for
 * the none subtree but still need to leave the space as large as that for the
 * other subtree. However, if two subtrees are none, then you don't need to
 * leave space for both of them.
 * Each unused space should contain an empty string "".
 * Print the subtrees following the same rules.
 *
 *
 * Example 1:
 *
 * Input:
 * ⁠    1
 * ⁠   /
 * ⁠  2
 * Output:
 * [["", "1", ""],
 * ⁠["2", "", ""]]
 *
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ⁠    1
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \
 * ⁠    4
 * Output:
 * [["", "", "", "1", "", "", ""],
 * ⁠["", "2", "", "", "", "3", ""],
 * ⁠["", "", "4", "", "", "", ""]]
 *
 *
 *
 * Example 3:
 *
 * Input:
 * ⁠     1
 * ⁠    / \
 * ⁠   2   5
 * ⁠  /
 * ⁠ 3
 * ⁠/
 * 4
 * Output:
 *
 * [["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 * ⁠["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 * ⁠["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 * ⁠["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
 *
 *
 *
 * Note:
 * The height of binary tree is in the range of [1, 10].
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
func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	s := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return [][]string{{s}}
	}

	var rightWider bool
	left := printTree(root.Left)
	height := len(left)
	halfWidth := 0
	if height > 0 {
		halfWidth = len(left[0])
	}

	right := printTree(root.Right)
	if height < len(right) {
		height = len(right)
		rightWider = true
	}

	var pad int
	if len(right) > 0 && halfWidth < len(right[0]) {
		if len(left) > 0 {
			pad = (len(right[0]) - halfWidth) / 2
		}
		halfWidth = len(right[0])
	} else if len(right) > 0 {
		pad = (halfWidth - len(right[0])) / 2
	}

	res := make([][]string, 0, height*2+1)
	line := make([]string, halfWidth*2+1)
	line[halfWidth] = s
	res = append(res, line)

	for i := 0; i < height; i++ {
		line = make([]string, 0, halfWidth*2+1)
		if rightWider {
			if i >= len(left) {
				line = append(line, make([]string, halfWidth)...)
			} else {
				mid := left[i]
				m := len(mid) / 2
				if mid[m] == "" {
					line = append(line, make([]string, pad/2)...)
					line = append(line, mid[:m]...)
					line = append(line, make([]string, pad)...)
					line = append(line, mid[m:]...)
					line = append(line, make([]string, pad/2)...)
				} else {
					line = append(line, make([]string, pad)...)
					line = append(line, mid...)
					line = append(line, make([]string, pad)...)
				}
			}

			line = append(line, "")
			line = append(line, right[i]...)
		} else {
			line = append(line, left[i]...)
			line = append(line, "")

			if i >= len(right) {
				line = append(line, make([]string, halfWidth)...)
			} else {
				mid := right[i]
				m := len(mid) / 2
				if mid[m] == "" {
					line = append(line, make([]string, pad/2)...)
					line = append(line, mid[:m]...)
					line = append(line, make([]string, pad)...)
					line = append(line, mid[m:]...)
					line = append(line, make([]string, pad/2)...)
				} else {
					line = append(line, make([]string, pad)...)
					line = append(line, mid...)
					line = append(line, make([]string, pad)...)
				}
			}
		}

		res = append(res, line)
	}

	return res
}
