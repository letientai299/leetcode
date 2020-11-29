package main

/*
 * @lc app=leetcode id=1017 lang=golang
 *
 * [1017] Convert to Base -2
 *
 * https://leetcode.com/problems/convert-to-base-2/description/
 *
 * algorithms
 * Medium (57.10%)
 * Total Accepted:    12.7K
 * Total Submissions: 21.4K
 * Testcase Example:  '2'
 *
 * Given a number N, return a string consisting of "0"s and "1"s that
 * represents its value in base -2 (negative two).
 *
 * The returned string must have no leading zeroes, unless the string is
 * "0".
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: "110"
 * Explantion: (-2) ^ 2 + (-2) ^ 1 = 2
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: "111"
 * Explantion: (-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3
 *
 *
 *
 * Example 3:
 *
 *
 * Input: 4
 * Output: "100"
 * Explantion: (-2) ^ 2 = 4
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= N <= 10^9
 *
 *
 *
 *
 */
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	buf := make([]byte, 0, 32)
	for n != 0 {
		if n%-2 != 0 {
			buf = append(buf, '1')
			n -= 1
		} else {
			buf = append(buf, '0')
		}
		n /= -2
	}
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	return string(buf)
}
