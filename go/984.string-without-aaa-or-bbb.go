package main

import "strings"

// String Without AAA or BBB
//
// Medium
//
// https://leetcode.com/problems/string-without-aaa-or-bbb/
//
// Given two integers `a` and `b`, return **any** string `s` such that:
//
// - `s` has length `a + b` and contains exactly `a``'a'` letters, and exactly
// `b``'b'` letters,
// - The substring `'aaa'` does not occur in `s`, and
// - The substring `'bbb'` does not occur in `s`.
//
// **Example 1:**
//
// ```
// Input: a = 1, b = 2
// Output: "abb"
// Explanation: "abb", "bab" and "bba" are all correct answers.
//
// ```
//
// **Example 2:**
//
// ```
// Input: a = 4, b = 1
// Output: "aabaa"
//
// ```
//
// **Constraints:**
//
// - `0 <= a, b <= 100`
// - It is guaranteed such an `s` exists for the given `a` and `b`.
func strWithout3a3b(a int, b int) string {
	x, y := 'a', 'b'
	if a < b {
		a, b = b, a
		x, y = y, x
	}

	var sb strings.Builder

	for a > 0 && b > 0 && (a-b)/2 > 0 {
		sb.WriteByte(byte(x))
		sb.WriteByte(byte(x))
		sb.WriteByte(byte(y))
		a -= 2
		b--
	}

	for a > 0 && b > 0 {
		sb.WriteByte(byte(x))
		sb.WriteByte(byte(y))
		a--
		b--
	}

	for a > 0 {
		sb.WriteByte(byte(x))
		a--
	}

	for b > 0 {
		sb.WriteByte(byte(y))
		b--
	}

	return sb.String()
}
