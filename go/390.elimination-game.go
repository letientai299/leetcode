package main

/*
 * @lc app=leetcode id=390 lang=golang
 *
 * [390] Elimination Game
 *
 * https://leetcode.com/problems/elimination-game/description/
 *
 * algorithms
 * Medium (44.01%)
 * Total Accepted:    35.1K
 * Total Submissions: 78.4K
 * Testcase Example:  '9'
 *
 *
 * There is a list of sorted integers from 1 to n. Starting from left to right,
 * remove the first number and every other number afterward until you reach the
 * end of the list.
 *
 * Repeat the previous step again, but this time from right to left, remove the
 * right most number and every other number from the remaining numbers.
 *
 * We keep repeating the steps again, alternating left to right and right to
 * left, until a single number remains.
 *
 * Find the last number that remains starting with a list of length n.
 *
 * Example:
 *
 * Input:
 * n = 9,
 * 1 2 3 4 5 6 7 8 9
 * 2 4 6 8
 * 2 6
 * 6
 *
 * 1 2 3 4 5 6 7 8 9 10
 * 2 4 6 8 10
 * 4 8
 * 8
 *
 * 1 2 3 4 5 6 7 8 9 10 11 12
 * 2 4 6 8 10 12
 * 2 6 10
 * 6
 *
 * Output:
 * 6
 *
 *
 */
func lastRemaining(n int) int {
	l, r := 1, n
	d := 1
	up := true
	for l < r {
		if up {
			r -= d - (r-l)%(2*d)
			l += d
		} else {
			l += d - (r-l)%(2*d)
			r -= d
		}
		d *= 2
		up = !up
	}
	return l
}
