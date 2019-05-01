package main

/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 *
 * https://leetcode.com/problems/base-7/description/
 *
 * algorithms
 * Easy (44.82%)
 * Total Accepted:    40.4K
 * Total Submissions: 90.2K
 * Testcase Example:  '100'
 *
 * Given an integer, return its base 7 string representation.
 *
 * Example 1:
 *
 * Input: 100
 * Output: "202"
 *
 *
 *
 * Example 2:
 *
 * Input: -7
 * Output: "-10"
 *
 *
 *
 * Note:
 * The input will be in range of [-1e7, 1e7].
 *
 */
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		return "-" + convertToBase7(-num)
	}

	var buf []byte
	for num != 0 {
		buf = append(buf, byte(num%7+'0'))
		num /= 7
	}

	// reverse the buffer
	for i, n := 0, len(buf); i < n/2; i++ {
		buf[i], buf[n-1-i] = buf[n-i-1], buf[i]
	}

	return string(buf)
}
