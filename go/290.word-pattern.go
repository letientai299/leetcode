package main

import (
	"strings"
)

/*
* @lc app=leetcode id=290 lang=golang
*
* [290] Word Pattern
*
* https://leetcode.com/problems/word-pattern/description/
*
* algorithms
* Easy (34.64%)
* Total Accepted:    133.4K
* Total Submissions: 385.2K
* Testcase Example:  '"abba"\n"dog cat cat dog"'
*
* Given a pattern and a string str, find if str follows the same pattern.
*
* Here follow means a full match, such that there is a bijection between a
* letter in pattern and a non-empty word in str.
*
* Example 1:
*
*
* Input: pattern = "abba", str = "dog cat cat dog"
* Output: true
*
* Example 2:
*
*
* Input:pattern = "abba", str = "dog cat cat fish"
* Output: false
*
* Example 3:
*
*
* Input: pattern = "aaaa", str = "dog cat cat dog"
* Output: false
*
* Example 4:
*
*
* Input: pattern = "abba", str = "dog dog dog dog"
* Output: false
*
* Notes:
* You may assume pattern contains only lowercase letters, and str contains
* lowercase letters that may be separated by a single space.
*
 */
func wordPattern(pattern string, str string) bool {
	map1 := make(map[string]byte)
	map2 := make(map[byte]string)
	i := 0
	for _, s := range strings.Split(str, " ") {
		if i >= len(pattern) {
			return false
		}

		x := pattern[i]

		if c, OK := map1[s]; !OK {
			map1[s] = x
		} else if c != x {
			return false
		}

		if t, OK := map2[x]; !OK {
			map2[x] = s
		} else if t != s {
			return false
		}

		i++
	}

	return i == len(pattern)
}
