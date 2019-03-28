package main

/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 *
 * https://leetcode.com/problems/reverse-string/description/
 *
 * algorithms
 * Easy (62.96%)
 * Total Accepted:    393K
 * Total Submissions: 624.2K
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * Write a function that reverses a string. The input string is given as an
 * array of characters char[].
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * You may assume all the characters consist of printable ascii
 * characters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["H","a","n","n","a","h"]
 * Output: ["h","a","n","n","a","H"]
 *
 *
 *
 *
 */
func reverseString(s []byte) {
	x := 0
	y := len(s) - 1
	for x < y {
		s[x], s[y] = s[y], s[x]
		x++
		y--
	}
}
