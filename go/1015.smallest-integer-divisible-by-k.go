package main

/*
 * @lc app=leetcode id=1015 lang=golang
 *
 * [1015] Smallest Integer Divisible by K
 *
 * https://leetcode.com/problems/smallest-integer-divisible-by-k/description/
 *
 * algorithms
 * Medium (30.27%)
 * Total Accepted:    25.9K
 * Total Submissions: 62.2K
 * Testcase Example:  '1'
 *
 * Given a positive integer K, you need to find the length of the smallest
 * positive integer N such that N is divisible by K, and N only contains the
 * digit 1.
 *
 * Return the length of N. If there is no such N, return -1.
 *
 * Note: N may not fit in a 64-bit signed integer.
 *
 *
 * Example 1:
 *
 *
 * Input: K = 1
 * Output: 1
 * Explanation: The smallest answer is N = 1, which has length 1.
 *
 *
 * Example 2:
 *
 *
 * Input: K = 2
 * Output: -1
 * Explanation: There is no such positive integer N divisible by 2.
 *
 *
 * Example 3:
 *
 *
 * Input: K = 3
 * Output: 3
 * Explanation: The smallest answer is N = 111, which has length 3.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= K <= 10^5
 *
 *
 */

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	m := 1
	n := 1
	for m%k != 0 {
		n++
		m = (((10 * m) % k) + 1) % k
	}

	return n
}
