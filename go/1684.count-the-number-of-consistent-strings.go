package main

// Count the Number of Consistent Strings
//
// Easy
//
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
//
// You are given a string `allowed` consisting of **distinct** characters and an
// array of strings `words`. A string is **consistent** if all characters in the
// string appear in the string `allowed`.
//
// Return _the number of **consistent** strings in the array_ `words`.
//
// **Example 1:**
//
// ```
// Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
// Output: 2
// Explanation: Strings "aaab" and "baa" are consistent since they only contain
// characters 'a' and 'b'.
//
// ```
//
// **Example 2:**
//
// ```
// Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
// Output: 7
// Explanation: All strings are consistent.
//
// ```
//
// **Example 3:**
//
// ```
// Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
// Output: 4
// Explanation: Strings "cc", "acd", "ac", and "d" are consistent.
//
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 104`
// - `1 <= allowed.length <= 26`
// - `1 <= words[i].length <= 10`
// - The characters in `allowed` are **distinct**.
// - `words[i]` and `allowed` contain only lowercase English letters.
func countConsistentStrings(allowed string, words []string) int {
	res := 0
	if len(allowed) == 26 {
		return len(words)
	}

	m := [26]bool{}
	for _, c := range allowed {
		m[c-'a'] = true
	}

	for _, w := range words {
		ok := true
		for _, c := range w {
			ok = m[c-'a']
			if !ok {
				break
			}
		}
		if ok {
			res++
		}
	}
	return res
}
