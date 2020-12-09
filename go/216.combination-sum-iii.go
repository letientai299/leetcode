package main

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (53.74%)
 * Total Accepted:    208.1K
 * Total Submissions: 349.3K
 * Testcase Example:  '3\n7'
 *
 * Find all valid combinations of k numbers that sum up to n such that the
 * following conditions are true:
 *
 *
 * Only numbers 1 through 9 are used.
 * Each number is used at most once.
 *
 *
 * Return a list of all possible valid combinations. The list must not contain
 * the same combination twice, and the combinations may be returned in any
 * order.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 * Explanation:
 * 1 + 2 + 4 = 7
 * There are no other valid combinations.
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6],[1,3,5],[2,3,4]]
 * Explanation:
 * 1 + 2 + 6 = 9
 * 1 + 3 + 5 = 9
 * 2 + 3 + 4 = 9
 * There are no other valid combinations.
 *
 *
 * Example 3:
 *
 *
 * Input: k = 4, n = 1
 * Output: []
 * Explanation: There are no valid combinations. [1,2,1] is not valid because 1
 * is used twice.
 *
 *
 * Example 4:
 *
 *
 * Input: k = 3, n = 2
 * Output: []
 * Explanation: There are no valid combinations.
 *
 *
 * Example 5:
 *
 *
 * Input: k = 9, n = 45
 * Output: [[1,2,3,4,5,6,7,8,9]]
 * Explanation:
 * 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
 * ​​​​​​​There are no other valid combinations.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= k <= 9
 * 1 <= n <= 60
 *
 *
 */

func combinationSum3(k int, n int) [][]int {
	var r [][]int
	buf := make([]int, 0, k)
	var find func(i, s int)

	find = func(i, s int) {
		if s == n {
			if len(buf) == k {
				cb := make([]int, k)
				copy(cb, buf)
				r = append(r, cb)
			}
			return
		}

		if len(buf) == k {
			return
		}

		l := len(buf)
		for j := i; j <= 9 && j <= n-s; j++ {
			buf = append(buf, j)
			find(j+1, s+j)
			buf = buf[:l]
		}
	}

	find(1, 0)
	return r
}
