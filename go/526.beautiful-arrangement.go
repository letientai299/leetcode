package main

/*
 * @lc app=leetcode id=526 lang=golang
 *
 * [526] Beautiful Arrangement
 *
 * https://leetcode.com/problems/beautiful-arrangement/description/
 *
 * algorithms
 * Medium (56.07%)
 * Total Accepted:    60.6K
 * Total Submissions: 101.5K
 * Testcase Example:  '2'
 *
 * Suppose you have n integers from 1 to n. We define a beautiful arrangement
 * as an array that is constructed by these n numbers successfully if one of
 * the following is true for the i^th position (1 <= i <= n) in this
 * array:
 *
 *
 * The number at the i^th position is divisible by i.
 * i is divisible by the number at the i^th position.
 *
 *
 * Given an integer n, return the number of the beautiful arrangements that you
 * can construct.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: 2
 * Explanation:
 * The first beautiful arrangement is [1, 2]:
 * Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).
 * Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).
 * The second beautiful arrangement is [2, 1]:
 * Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).
 * Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 15
 *
 *
 */

func countArrangement(n int) int {
	count := 0
	used := 0
	candidates := make([][]int, n)
	for i := range candidates {
		x := i + 1
		candidates[i] = append(candidates[i], x)

		for k := x + 1; k <= n; k++ {
			if k%x == 0 {
				candidates[i] = append(candidates[i], k)
				candidates[k-1] = append(candidates[k-1], x)
			}
		}
	}

	var walk func(i int)
	walk = func(i int) {
		where := i + 1
		if where > n {
			count++
			return
		}

		for _, k := range candidates[i] {
			mask := 1 << (k - 1)
			if used&mask == 0 {
				used |= mask
				walk(i + 1)
				used &= ^mask
			}
		}
	}

	walk(0)
	return count
}
