package main

/*
 * @lc app=leetcode id=712 lang=golang
 *
 * [712] Minimum ASCII Delete Sum for Two Strings
 *
 * https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description/
 *
 * algorithms
 * Medium (56.17%)
 * Total Accepted:    41.8K
 * Total Submissions: 70.5K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * Given two strings s1, s2, find the lowest ASCII sum of deleted characters to
 * make two strings equal.
 *
 * Example 1:
 *
 * Input: s1 = "sea", s2 = "eat"
 * Output: 231
 * Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to
 * the sum.
 * Deleting "t" from "eat" adds 116 to the sum.
 * At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum
 * possible to achieve this.
 *
 *
 *
 * Example 2:
 *
 * Input: s1 = "delete", s2 = "leet"
 * Output: 403
 * Explanation: Deleting "dee" from "delete" to turn the string into "let",
 * adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e]
 * to the sum.
 * At the end, both strings are equal to "let", and the answer is
 * 100+101+101+101 = 403.
 * If instead we turned both strings into "lee" or "eet", we would get answers
 * of 433 or 417, which are higher.
 *
 *
 *
 * Note:
 * 0 < s1.length, s2.length .
 * All elements of each string will have an ASCII value in [97, 122].
 *
 */

func minimumDeleteSum(a string, b string) int {
	m, n := len(a), len(b)
	r := 0
	if m == 0 {
		for _, c := range b {
			r += int(c)
		}
		return r
	}

	if n == 0 {
		for _, c := range a {
			r += int(c)
		}
		return r
	}

	var dp [][]int
	dp = append(dp, make([]int, n))

	dp[0][0] = int(a[0]) + int(b[0])
	if a[0] == b[0] {
		dp[0][0] = 0
	}

	// first line
	cur := int(b[0])
	for i := 1; i < n; i++ {
		v := 0
		if b[i] == a[0] {
			v = cur
		} else {
			v = int(b[i]) + dp[0][i-1]
		}

		cur += int(b[i])
		dp[0][i] = v
	}

	if m == 1 {
		return dp[0][n-1]
	}

	dp = append(dp, make([]int, n))
	cur = int(a[0])
	for y := 1; y < m; y++ {
		if a[y] == b[0] {
			dp[1][0] = cur
		} else {
			dp[1][0] = dp[0][0] + int(a[y])
		}
		cur += int(a[y])

		for x := 1; x < n; x++ {
			if a[y] == b[x] {
				dp[1][x] = dp[0][x-1]
				continue
			}

			min := dp[0][x] + int(a[y])
			if min > dp[1][x-1]+int(b[x]) {
				min = dp[1][x-1] + int(b[x])
			}

			dp[1][x] = min
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	return dp[0][n-1]
}
