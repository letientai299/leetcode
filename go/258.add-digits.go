package main

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 *
 * https://leetcode.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (53.67%)
 * Total Accepted:    230.4K
 * Total Submissions: 429.4K
 * Testcase Example:  '38'
 *
 * Given a non-negative integer num, repeatedly add all its digits until the
 * result has only one digit.
 *
 * Example:
 *
 *
 * Input: 38
 * Output: 2
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
 * Since 2 has only one digit, return it.
 *
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	x := num % 9
	if x == 0 {
		return 9
	}

	return x
}
