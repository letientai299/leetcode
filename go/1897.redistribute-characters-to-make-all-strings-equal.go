package main

// Redistribute Characters to Make All Strings Equal
//
// Easy
//
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
//
// You are given an array of strings `words` ( **0-indexed**).
//
// In one operation, pick two **distinct** indices `i` and `j`, where `words[i]`
// is a non-empty string, and move **any** character from `words[i]` to **any**
// position in `words[j]`.
//
// Return `true` _if you can make **every** string in_ `words` _**equal** using
// **any** number of operations_, _and_ `false` _otherwise_.
//
// **Example 1:**
//
// ```
// words[1] to the front of words[2],
// to make words[1]true
// ```
//
// **Example 2:**
//
// ```
// Input: words = ["ab","a"]
// Output: false
// Explanation: It is impossible to make all the strings equal using the
// operation.
//
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 100`
// - `1 <= words[i].length <= 100`
// - `words[i]` consists of lowercase English letters.
func makeEqual(words []string) bool {
	m := [26]int{}
	for _, w := range words {
		for _, c := range w {
			m[c-'a']++
		}
	}

	n := len(words)
	for _, v := range m {
		if v%n != 0 {
			return false
		}
	}

	return true
}
