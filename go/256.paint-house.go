package main

/*
 * @lc app=leetcode id=256 lang=golang
 *
 * [256] Paint House
 *
 * https://leetcode.com/problems/paint-house/description/
 *
 * algorithms
 * Easy (50.28%)
 * Total Accepted:    66.7K
 * Total Submissions: 132.1K
 * Testcase Example:  '[[17,2,17],[16,16,5],[14,3,19]]'
 *
 * There are a row of n houses, each house can be painted with one of the three
 * colors: red, blue or green. The cost of painting each house with a certain
 * color is different. You have to paint all the houses such that no two
 * adjacent houses have the same color.
 *
 * The cost of painting each house with a certain color is represented by a n x
 * 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with
 * color red; costs[1][2] is the cost of painting house 1 with color green, and
 * so on... Find the minimum cost to paint all houses.
 *
 * Note:
 * All costs are positive integers.
 *
 * Example:
 *
 *
 * Input: [[17,2,17],[16,16,5],[14,3,19]]
 * Output: 10
 * Explanation: Paint house 0 into blue, paint house 1 into green, paint house
 * 2 into blue.
 * Minimum cost: 2 + 5 + 3 = 10.
 *
 *
 */
func minCost256(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	col := costs[0]
	for i := 1; i < n; i++ {
		prev := costs[i-1]
		col = costs[i]
		col[0] += min(prev[1], prev[2])
		col[1] += min(prev[2], prev[0])
		col[2] += min(prev[0], prev[1])
	}

	return min(min(col[0], col[1]), col[2])
}
