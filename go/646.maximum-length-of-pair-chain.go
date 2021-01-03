package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=646 lang=golang
 *
 * [646] Maximum Length of Pair Chain
 *
 * https://leetcode.com/problems/maximum-length-of-pair-chain/description/
 *
 * algorithms
 * Medium (50.05%)
 * Total Accepted:    68.2K
 * Total Submissions: 129.7K
 * Testcase Example:  '[[1,2], [2,3], [3,4]]'
 *
 *
 * You are given n pairs of numbers. In every pair, the first number is always
 * smaller than the second number.
 *
 *
 *
 * Now, we define a pair (c, d) can follow another pair (a, b) if and only if b
 * < c. Chain of pairs can be formed in this fashion.
 *
 *
 *
 * Given a set of pairs, find the length longest chain which can be formed. You
 * needn't use up all the given pairs. You can select pairs in any order.
 *
 *
 *
 * Example 1:
 *
 * Input: [[1,2], [2,3], [3,4]]
 * Output: 2
 * Explanation: The longest chain is [1,2] -> [3,4]
 *
 *
 *
 * Note:
 *
 * The number of given pairs will be in the range [1, 1000].
 *
 *
 */

// TODO (tai): slow, 45%, there much a nlogn solution.
func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	if n < 2 {
		return n
	}

	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})

	dp := make([]int, n)
	dp[0] = 1
	best := 1
	for i := 1; i < n; i++ {
		cur := pairs[i]
		pre := pairs[i-1]
		if cur[0] > pre[1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}

		for j := i - 2; j >= 0; j-- {
			a := pairs[j]
			if a[1] < cur[0] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}

		if best < dp[i] {
			best = dp[i]
		}
	}

	fmt.Println(dp)
	fmt.Println(pairs)

	return best
}
