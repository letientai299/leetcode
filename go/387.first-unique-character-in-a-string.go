package main

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (49.48%)
 * Total Accepted:    243.4K
 * Total Submissions: 491.9K
 * Testcase Example:  '"leetcode"'
 *
 *
 * Given a string, find the first non-repeating character in it and return it's
 * index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 *
 *
 *
 *
 * Note: You may assume the string contain only lowercase letters.
 *
 */
func firstUniqChar(s string) int {
	count := make(map[int32]int)
	for _, c := range s {
		count[c]++
	}

	for i, c := range s {
		if count[c] == 1 {
			return i
		}
	}
	return -1
}
