package main

import "math"

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (38.08%)
 * Total Accepted:    280.7K
 * Total Submissions: 736.3K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	ai, bi := len(a)-1, len(b)-1
	n := int(math.Max(float64(len(a)), float64(len(b))))
	c := make([]byte, n+1)
	var r byte
	for n >= 0 {
		var x = r
		if ai != -1 {
			if a[ai] == '1' {
				x += 1
			}
			ai--
		}

		if bi != -1 {
			if b[bi] == '1' {
				x += 1
			}
			bi--
		}

		if x >= 2 {
			r = x / 2
			x %= 2
		} else {
			r = 0
		}

		c[n] = x + '0'
		n--
	}

	c[0] += r

	if c[0] == '0' {
		return string(c[1:])
	}
	return string(c)
}
