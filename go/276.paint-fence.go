package main

/*
 * @lc app=leetcode id=276 lang=golang
 *
 * [276] Paint Fence
 *
 * https://leetcode.com/problems/paint-fence/description/
 *
 * algorithms
 * Easy (37.24%)
 * Total Accepted:    59.3K
 * Total Submissions: 153.6K
 * Testcase Example:  '3\n2'
 *
 * There is a fence with n posts, each post can be painted with one of the k
 * colors.
 *
 * You have to paint all the posts such that no more than two adjacent fence
 * posts have the same color.
 *
 * Return the total number of ways you can paint the fence.
 *
 * Note:
 * n and k are non-negative integers.
 *
 * Example:
 *
 *
 * Input: n = 3, k = 2
 * Output: 6
 * Explanation: Take c1 as color 1, c2 as color 2. All possible ways
 * are:
 *
 *            post1  post2  post3
 * -----      -----  -----  -----
 *   1         c1     c1     c2
 *   2         c1     c2     c1
 *   3         c1     c2     c2
 *   4         c2     c1     c1
 *   5         c2     c1     c2
 *   6         c2     c2     c1
 *
 *
 */
func numWays(n int, k int) int {
	dp1 := k
	dp2 := k * k

	switch n {
	case 0:
		return 0
	case 1:
		return dp1
	case 2:
		return dp2
	}

	for ; n > 2; n-- {
		// i-3 | i-2 | i-1  | ... |
		// dp0 | dp1 | dp2 |
		v := (k-1)*dp2 + // i has different color with i-1
			+(k-1)*dp1 // i has different color with i-2
		dp1, dp2 = dp2, v
	}

	return dp2
}
