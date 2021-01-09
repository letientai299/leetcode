package main

/*
 * @lc app=leetcode id=738 lang=golang
 *
 * [738] Monotone Increasing Digits
 *
 * https://leetcode.com/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (42.78%)
 * Total Accepted:    26.6K
 * Total Submissions: 58.5K
 * Testcase Example:  '10'
 *
 *
 * Given a non-negative integer N, find the largest number that is less than or
 * equal to N with monotone increasing digits.
 *
 * (Recall that an integer has monotone increasing digits if and only if each
 * pair of adjacent digits x and y satisfy x .)
 *
 *
 * Example 1:
 *
 * Input: N = 10
 * Output: 9
 *
 *
 *
 * Example 2:
 *
 * Input: N = 1234
 * Output: 1234
 *
 *
 *
 * Example 3:
 *
 * Input: N = 332
 * Output: 299
 *
 *
 *
 * Note:
 * N is an integer in the range [0, 10^9].
 *
 */

func monotoneIncreasingDigits(N int) int {
	base := 1
	r := 0
	p := 10
	for n := N; n > 0; base *= 10 {
		d := n % 10
		n /= 10
		if d <= p {
			r += base * d
			p = d
			continue
		}
		d--
		p = d
		r = base - 1 + d*base
	}
	if r > N {
		r -= base / 10
	}
	return r
}
