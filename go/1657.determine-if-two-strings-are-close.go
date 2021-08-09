package main

// Determine if Two Strings Are Close
//
// Medium
//
// https://leetcode.com/problems/determine-if-two-strings-are-close/
//
// Two strings are considered **close** if you can attain one from the other
// using the following operations:
//
// - Operation 1: Swap any two **existing** characters.
//
//
//     - For example, `abcde -> aecdb`
// - Operation 2: Transform **every** occurrence of one **existing** character
// into another **existing** character, and do the same with the other
// character.
//
//     - For example, `aacabb -> bbcbaa` (all `a`'s turn into `b`'s, and all
// `b`'s turn into `a`'s)
//
// You can use the operations on either string as many times as necessary.
//
// Given two strings, `word1` and `word2`, return `true` _if_ `word1` _and_
// `word2` _are **close**, and_ `false` _otherwise._
//
// **Example 1:**
//
// ```
// Input: word1 = "abc", word2 = "bca"
// Output: true
// Explanation: You can attain word2 from word1 in 2 operations.
// Apply Operation 1: "abc" -> "acb"
// Apply Operation 1: "acb" -> "bca"
//
// ```
//
// **Example 2:**
//
// ```
// Input: word1 = "a", word2 = "aa"
// Output: false
// Explanation: It is impossible to attain word2 from word1, or vice versa, in
// any number of operations.
//
// ```
//
// **Example 3:**
//
// ```
// Apply Operation 2: "
// ```
//
// **Example 4:**
//
// ```
// Input: word1 = "cabbba", word2 = "aabbss"
// Output: false
// Explanation: It is impossible to attain word2 from word1, or vice versa, in
// any amount of operations.
//
// ```
//
// **Constraints:**
//
// - `1 <= word1.length, word2.length <= 105`
// - `word1` and `word2` contain only lowercase English letters.
func closeStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	f1 := make([]int, 26)
	f2 := make([]int, 26)
	for i := 0; i < len(a); i++ {
		f1[a[i]-'a']++
		f2[b[i]-'a']++
	}

	m := make(map[int]int)
	for i, v := range f1 {
		if (v == 0) != (f2[i] == 0) {
			return false // has different character set
		}
		m[v]++
		m[f2[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

// Good for interview
