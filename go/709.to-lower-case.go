package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 *
 * https://leetcode.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (77.08%)
 * Total Accepted:    123.7K
 * Total Submissions: 160.4K
 * Testcase Example:  '"Hello"'
 *
 * Implement function ToLowerCase() that has a string parameter str, and
 * returns the same string in lowercase.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "Hello"
 * Output: "hello"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "here"
 * Output: "here"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "LOVELY"
 * Output: "lovely"
 *
 *
 *
 *
 *
 */
func toLowerCase(str string) string {
	sb := strings.Builder{}
	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			c = c + ('a' - 'A')
		}
		sb.WriteByte(byte(c))
	}
	return sb.String()
}
