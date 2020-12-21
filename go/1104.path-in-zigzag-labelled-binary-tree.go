package main

/*
 * @lc app=leetcode id=1104 lang=golang
 *
 * [1104] Path In Zigzag Labelled Binary Tree
 *
 * https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/description/
 *
 * algorithms
 * Medium (70.47%)
 * Total Accepted:    21.6K
 * Total Submissions: 29.6K
 * Testcase Example:  '14'
 *
 * In an infinite binary tree where every node has two children, the nodes are
 * labelled in row order.
 *
 * In the odd numbered rows (ie., the first, third, fifth,...), the labelling
 * is left to right, while in the even numbered rows (second, fourth,
 * sixth,...), the labelling is right to left.
 *
 *
 *
 * Given the label of a node in this tree, return the labels in the path from
 * the root of the tree to the node with that label.
 *
 *
 * Example 1:
 *
 *
 * Input: label = 14
 * Output: [1,3,4,14]
 *
 *
 * Example 2:
 *
 *
 * Input: label = 26
 * Output: [1,2,6,10,26]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= label <= 10^6
 *
 *
 */
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}

	i := 0
	for 1<<i <= label {
		i++
	}

	r := make([]int, i)
	r[i-1] = label
	r[0] = 1
	for j := i - 2; j > 0; j-- {
		// range: 1<<j .. 1 << (j+1)-1
		label /= 2
		label = 1<<j + (1<<(j+1) - 1 - label)
		r[j] = label
	}
	return r
}
