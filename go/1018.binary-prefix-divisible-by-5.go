package main

/*
 * @lc app=leetcode id=1018 lang=golang
 *
 * [1018] Binary Prefix Divisible By 5
 *
 * https://leetcode.com/problems/binary-prefix-divisible-by-5/description/
 *
 * algorithms
 * Easy (46.87%)
 * Total Accepted:    26.5K
 * Total Submissions: 55.6K
 * Testcase Example:  '[0,1,1]'
 *
 * Given an array A of 0s and 1s, consider N_i: the i-th subarray from A[0] to
 * A[i] interpreted as a binary number (from most-significant-bit to
 * least-significant-bit.)
 *
 * Return a list of booleans answer, where answer[i] is true if and only if N_i
 * is divisible by 5.
 *
 * Example 1:
 *
 *
 * Input: [0,1,1]
 * Output: [true,false,false]
 * Explanation:
 * The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in
 * base-10.  Only the first number is divisible by 5, so answer[0] is true.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,1,1]
 * Output: [false,false,false]
 *
 *
 * Example 3:
 *
 *
 * Input: [0,1,1,1,1,1]
 * Output: [true,false,false,false,true,false]
 *
 *
 * Example 4:
 *
 *
 * Input: [1,1,1,0,1]
 * Output: [false,false,false,false,false]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 30000
 * A[i] is 0 or 1
 *
 *
 */
func prefixesDivBy5(a []int) []bool {
	res := make([]bool, 0, len(a))
	v := 0
	for _, x := range a {
		v = (v<<1 + x) % 5
		res = append(res, v == 0)
	}
	return res
}
