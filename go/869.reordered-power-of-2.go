package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=869 lang=golang
 *
 * [869] Reordered Power of 2
 *
 * https://leetcode.com/problems/reordered-power-of-2/description/
 *
 * algorithms
 * Medium (51.46%)
 * Total Accepted:    15.9K
 * Total Submissions: 29.5K
 * Testcase Example:  '1'
 *
 * Starting with a positive integer N, we reorder the digits in any order
 * (including the original order) such that the leading digit is not zero.
 *
 * Return true if and only if we can do this in a way such that the resulting
 * number is a power of 2.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 10
 * Output: false
 *
 *
 *
 * Example 3:
 *
 *
 * Input: 16
 * Output: true
 *
 *
 *
 * Example 4:
 *
 *
 * Input: 24
 * Output: false
 *
 *
 *
 * Example 5:
 *
 *
 * Input: 46
 * Output: true
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 10^9
 *
 *
 *
 *
 *
 *
 *
 */

func reorderedPowerOf2(n int) bool {
	if n < 10 {
		return n == 1 || n == 2 || n == 4 || n == 8
	}

	b := 8
	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	for i := 0; i <= 28; i++ {
		b <<= 1
		if b%9 != n%9 {
			continue
		}
		t := []byte(strconv.Itoa(b))
		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		if string(s) == string(t) {
			return true
		}
	}
	return false
}
