package main

/*
* @lc app=leetcode id=205 lang=golang
*
* [205] Isomorphic Strings
*
* https://leetcode.com/problems/isomorphic-strings/description/
*
* algorithms
* Easy (36.88%)
* Total Accepted:    191.6K
* Total Submissions: 519.5K
* Testcase Example:  '"egg"\n"add"'
*
* Given two strings s and t, determine if they are isomorphic.
*
* Two strings are isomorphic if the characters in s can be replaced to get t.
*
* All occurrences of a character must be replaced with another character while
* preserving the order of characters. No two characters may map to the same
* character but a character may map to itself.
*
* Example 1:
*
*
* Input: s = "egg", t = "add"
* Output: true
*
*
* Example 2:
*
*
* Input: s = "foo", t = "bar"
* Output: false
*
* Example 3:
*
*
* Input: s = "paper", t = "title"
* Output: true
*
* Note:
* You may assume both s and t have the same length.
*
 */
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smt := make(map[int32]byte)
	tms := make(map[byte]int32)
	for i, c := range s {
		if val, ok := smt[c]; ok {
			if t[i] != val {
				return false
			}
		} else {
			smt[c] = t[i]
		}

		if val, ok := tms[t[i]]; ok {
			if c != val {
				return false
			}
		} else {
			tms[t[i]] = c
		}
	}

	return true
}
