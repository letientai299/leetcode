package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (51.89%)
 * Total Accepted:    644.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given an array of distinct integers candidates and a target integer target,
 * return a list of all unique combinations of candidates where the chosen
 * numbers sum to target. You may return the combinations in any order.
 *
 * The same number may be chosen from candidates an unlimited number of times.
 * Two combinations are unique if the frequency of at least one of the chosen
 * numbers is different.
 *
 * It is guaranteed that the number of unique combinations that sum up to
 * target is less than 150 combinations for the given input.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7
 * Output: [[2,2,3],[7]]
 * Explanation:
 * 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
 * times.
 * 7 is a candidate, and 7 = 7.
 * These are the only two combinations.
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8
 * Output: [[2,2,2,2],[2,3,3],[3,5]]
 *
 *
 * Example 3:
 *
 *
 * Input: candidates = [2], target = 1
 * Output: []
 *
 *
 * Example 4:
 *
 *
 * Input: candidates = [1], target = 1
 * Output: [[1]]
 *
 *
 * Example 5:
 *
 *
 * Input: candidates = [1], target = 2
 * Output: [[1,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * All elements of candidates are distinct.
 * 1 <= target <= 500
 *
 *
 */
func combinationSum(arr []int, n int) [][]int {
	sort.Ints(arr)
	if arr[0] > n {
		return nil
	}

	var gen func(i, n int)
	res := make([][]int, 0, len(arr)*n/arr[0]) // guess

	buf := make([]int, 0, n/arr[0])
	gen = func(i, n int) {
		if i >= len(arr) || n < 0 {
			return
		}

		if n == 0 {
			b := make([]int, len(buf))
			copy(b, buf)
			res = append(res, b)
			return
		}

		l := len(buf)
		for ; i < len(arr) && arr[i] <= n && l != 0; i++ {
			buf = append(buf, arr[i])
			gen(i, n-arr[i])
			buf = buf[:l]
		}

		return
	}

	for i, x := range arr {
		buf = append(buf, x)
		gen(i, n-x)
		buf = buf[:0]
	}

	return res
}
