package main

/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 *
 * https://leetcode.com/problems/min-cost-climbing-stairs/description/
 *
 * algorithms
 * Easy (47.62%)
 * Total Accepted:    89.5K
 * Total Submissions: 187.4K
 * Testcase Example:  '[0,0,0,0]'
 *
 *
 * On a staircase, the i-th step has some non-negative cost cost[i] assigned (0
 * indexed).
 *
 * Once you pay the cost, you can either climb one or two steps. You need to
 * find minimum cost to reach the top of the floor, and you can either start
 * from the step with index 0, or the step with index 1.
 *
 *
 * Example 1:
 *
 * Input: cost = [10, 15, 20]
 * Output: 15
 * Explanation: Cheapest is start on cost[1], pay that cost and go to the
 * top.
 *
 *
 *
 * Example 2:
 *
 * Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
 * Output: 6
 * Explanation: Cheapest is start on cost[0], and only step on 1s, skipping
 * cost[3].
 *
 *
 *
 * Note:
 *
 * cost will have a length in the range [2, 1000].
 * Every cost[i] will be an integer in the range [0, 999].
 *
 *
 */
func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	a, b := cost[0], cost[1]
	if a > b {
		a, b = b, a
	}

	for i := 2; i < n; i++ {
		cost[i] = a + cost[i]
		if cost[i-1] < cost[i] {
			a, b = cost[i-1], cost[i]
		} else {
			a, b = cost[i], cost[i-1]
		}
	}

	return a
}
