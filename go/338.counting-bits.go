package main

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Medium (66.09%)
 * Total Accepted:    321.5K
 * Total Submissions: 458.8K
 * Testcase Example:  '2'
 *
 * Given a non negative integer number num. For every numbers i in the range 0
 * ≤ i ≤ num calculate the number of 1's in their binary representation and
 * return them as an array.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,1]
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with run time
 * O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a
 * single pass?
 * Space complexity should be O(n).
 * Can you do it like a boss? Do it without using any builtin function like
 * __builtin_popcount in c++ or in any other language.
 */
func countBits(n int) []int {
	r := make([]int, n+1)
	if n == 0 {
		return r
	}

	b := 1
	for b <= n {
		for i := 0; i <= b && i+b < len(r); i++ {
			r[i+b] = r[i] + 1
		}
		b *= 2
	}

	return r
}
