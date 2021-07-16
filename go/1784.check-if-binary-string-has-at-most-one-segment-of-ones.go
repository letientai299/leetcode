package main

// Check if Binary String Has at Most One Segment of Ones
//
// Easy
//
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
//
// Given a binary string `s` **​​​​​without leading zeros**, return `true`​​​
// _if_ `s` _contains **at most one contiguous segment of ones**_. Otherwise,
// return `false`.
//
// **Example 1:**
//
// ```
// Input: s = "1001"
// Output: false
// Explanation: The ones do not form a contiguous segment.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "110"
// Output: true
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 100`
// - `s[i]`​​​​ is either `'0'` or `'1'`.
// - `s[0]` is `'1'`.
func checkOnesSegment(s string) bool {
	count := 0
	run := false

	for _, c := range s {
		if c == '1' {
			if !run {
				count++
				if count == 2 {
					return false
				}
				run = true
			}
			continue
		}

		run = false
	}

	return true
}
