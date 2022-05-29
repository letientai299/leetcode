package main

import "strings"

// Counting Words With a Given Prefix
//
// Easy
//
// https://leetcode.com/problems/counting-words-with-a-given-prefix/
//
// You are given an array of strings `words` and a string `pref`.
//
// Return _the number of strings in_ `words` _that contain_ `pref` _as a
// **prefix**_.
//
// A **prefix** of a string `s` is any leading contiguous substring of `s`.
//
// **Example 1:**
//
// ```
// pref
// ```
//
// **Example 2:**
//
// ```
// pref
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 100`
// - `1 <= words[i].length, pref.length <= 100`
// - `words[i]` and `pref` consist of lowercase English letters.
func prefixCount(words []string, pref string) int {
	r := 0
	for _, w := range words {
		if len(pref) <= len(words) && strings.HasPrefix(w, pref) {
			r++
		}
	}

	return r
}
