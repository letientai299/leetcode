package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Check If a String Contains All Binary Codes of Size K
//
// Medium
//
// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
//
// Given a binary string `s` and an integer `k`.
//
// Return `true` _if every binary code of length_ `k` _is a substring of_ `s`.
// Otherwise, return `false`.
//
// **Example 1:**
//
// ```
// Input: s = "00110110", k = 2
// Output: true
// Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They
// can be all found as substrings at indicies 0, 1, 3 and 2 respectively.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "00110", k = 2
// Output: true
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "0110", k = 1
// Output: true
// Explanation: The binary codes of length 1 are "0" and "1", it is clear that
// both exist as a substring.
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = "0110", k = 2
// Output: false
// Explanation: The binary code "00" is of length 2 and doesn't exist in the
// array.
//
// ```
//
// **Example 5:**
//
// ```
// Input: s = "0000000001011100", k = 4
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 5 * 105`
// - `s[i]` is either `'0'` or `'1'`.
// - `1 <= k <= 20`
func hasAllCodes(s string, k int) bool {
	n := 1 << k
	format := fmt.Sprintf("%%0%ds", k)
	for i := 0; i < n; i++ {
		v := strconv.FormatInt(int64(i), 2)
		sub := fmt.Sprintf(format, v)
		if !strings.Contains(s, sub) {
			return false
		}
	}

	return true
}

// TODO (tai): too slow, but somehow still pass all the test: 4724 ms, 6.67%
