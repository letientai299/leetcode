package main

// Second Largest Digit in a String
//
// Easy
//
// https://leetcode.com/problems/second-largest-digit-in-a-string/
//
// Given an alphanumeric string `s`, return _the **second largest** numerical
// digit that appears in_ `s` _, or_`-1` _if it does not exist_.
//
// An **alphanumeric** string is a string consisting of lowercase English
// letters and digits.
//
// **Example 1:**
//
// ```
// Input: s = "dfa12321afd"
// Output: 2
// Explanation: The digits that appear in s are [1, 2, 3]. The second largest
// digit is 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "abc1111"
// Output: -1
// Explanation: The digits that appear in s are [1]. There is no second largest
// digit.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 500`
// - `s` consists of only lowercase English letters and/or digits.
func secondHighest(s string) int {
	bm := 0 // bitmap
	top := -1
	for _, c := range s {
		if '0' <= c && c <= '9' {
			d := int(c - '0')
			if d > top {
				top = d
			}
			bm |= 1 << (d + 1)
		}
	}

	top--
	for top >= 0 {
		if bm&(1<<(top+1)) != 0 {
			return top
		}
		top--
	}

	return -1
}
