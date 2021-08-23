package main

import "strings"

// Longest Happy String
//
// Medium
//
// https://leetcode.com/problems/longest-happy-string/
//
// A string is called _happy_ if it does not have any of the strings `'aaa'`,
// `'bbb'` or `'ccc'` as a substring.
//
// Given three integers `a`, `b` and `c`, return **any** string `s`, which
// satisfies following conditions:
//
// - `s` is _happy_ and longest possible.
// - `s` contains **at most** `a` occurrences of the letter `'a'`, **at most**
// `b` occurrences of the letter `'b'` and **at most** `c` occurrences of the
// letter `'c'`.
// - `s `will only contain `'a'`, `'b'` and `'c'` letters.
//
// If there is no such string `s` return the empty string `""`.
//
// **Example 1:**
//
// ```
// Input: a = 1, b = 1, c = 7
// Output: "ccaccbcc"
// Explanation: "ccbccacc" would also be a correct answer.
//
// ```
//
// **Example 2:**
//
// ```
// Input: a = 2, b = 2, c = 1
// Output: "aabbc"
//
// ```
//
// **Example 3:**
//
// ```
// Input: a = 7, b = 1, c = 0
// Output: "aabaa"
// Explanation: It's the only correct answer in this case.
//
// ```
//
// **Constraints:**
//
// - `0 <= a, b, c <= 100`
// - `a + b + c > 0`
func longestDiverseString(a int, b int, c int) string {
	var sb strings.Builder

	pick := func(pre []byte) byte {
		if a >= b && a >= c && string(pre) != "aa" {
			return 'a'
		}

		if b >= a && b >= c && string(pre) != "bb" {
			return 'b'
		}

		if c >= a && c >= b && string(pre) != "cc" {
			return 'c'
		}

		switch string(pre) {
		case "aa":
			if b >= c && b > 0 {
				return 'b'
			}
			if c > 0 {
				return 'c'
			}
		case "bb":
			if a >= c && a > 0 {
				return 'a'
			}

			if c > 0 {
				return 'c'
			}
		case "cc":
			if a >= b && a > 0 {
				return 'a'
			}
			if b > 0 {
				return 'b'
			}
		}

		return 0
	}

	pre := []byte("zx")
	for a > 0 || b > 0 || c > 0 {
		x := pick(pre)
		switch x {
		case 'a':
			a--
		case 'b':
			b--
		case 'c':
			c--
		default:
			return sb.String()
		}
		sb.WriteByte(x)
		pre[0], pre[1] = pre[1], x
	}

	return sb.String()
}
