package main

import "fmt"

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (48.65%)
 * Total Accepted:    335.1K
 * Total Submissions: 622K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 19
 *
 *
 */
func numTrees(n int) int {
	// https://stackoverflow.com/a/16004841/3869533

	if n < 3 {
		return n
	}

	mem := make([]int, 0, n+1)
	mem = append(mem, 1, 1, 2)
	for i := 3; i <= n; i++ {
		v := 0
		for j := 0; j <= (i-1)/2; j++ {
			k := i - 1 - j
			v += 2 * mem[j] * mem[k]
		}

		if i%2 == 1 {
			v -= mem[i/2] * mem[i/2]
		}
		mem = append(mem, v)
	}

	for i, v := range mem {
		fmt.Println(i, v)
	}

	return mem[n]
}
