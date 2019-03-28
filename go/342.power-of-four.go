package main

/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (40.04%)
 * Total Accepted:    108.2K
 * Total Submissions: 270.3K
 * Testcase Example:  '16'
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 *
 * Example 1:
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
 * Input: 5
 * Output: false
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 */
func isPowerOfFour(n int) bool {
	return (n&(n-1) == 0) && (n-1)%3 == 0
}
