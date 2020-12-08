package main

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (36.82%)
 * Total Accepted:    163.6K
 * Total Submissions: 413.7K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 *
 * Input: [5,7]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1]
 * Output: 0
 */

func rangeBitwiseAnd(m int, n int) int {
	r := 0
	b := 1
	for b*2 <= n {
		b *= 2
	}

	for m > 0 {
		if b > m {
			break
		}

		r |= b
		m, n = m-b, n-b
		for b > n {
			b /= 2
		}
	}
	return r
}

/*
	Damn, 0ms!

func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= n-1
	}
	return m&n
}
*/
