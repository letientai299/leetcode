package main

/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (39.50%)
 * Total Accepted:    103.2K
 * Total Submissions: 261.2K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 * Note: Do not use any built-in library function such as sqrt.
 *
 * Example 1:
 *
 *
 *
 * Input: 16
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 14
 * Output: false
 *
 *
 *
 */
func isPerfectSquare(n int) bool {
	if n <= 0 {
		return false
	}

	x := 1
	for ; x*x < n; x++ {
	}
	return x*x == n
}
