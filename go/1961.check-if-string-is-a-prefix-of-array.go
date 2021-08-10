package main

// Check If String Is a Prefix of Array
//
// Easy
//
// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/
//
// Given a string `s` and an array of strings `words`, determine whether `s` is
// a **prefix string** of `words`.
//
// A string `s` is a **prefix string** of `words` if `s` can be made by
// concatenating the first `k` strings in `words` for some **positive** `k` no
// larger than `words.length`.
//
// Return `true` _if_ `s` _is a **prefix string** of_ `words` _, or_ `false`
// _otherwise_.
//
// **Example 1:**
//
// ```
// Input: s = "iloveleetcode", words = ["i","love","leetcode","apples"]
// Output: true
// Explanation:
// s can be made by concatenating "i", "love", and "leetcode" together.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "iloveleetcode", words = ["apples","i","love","leetcode"]
// Output: false
// Explanation:
// It is impossible to make s using a prefix of arr.
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 100`
// - `1 <= words[i].length <= 20`
// - `1 <= s.length <= 1000`
// - `words[i]` and `s` consist of only lowercase English letters.
func isPrefixString(s string, words []string) bool {
	i := 0
	wi := 0
	j := 0
	for i < len(s) && wi < len(words) {
		w := words[wi]
		for i < len(s) && j < len(w) {
			if s[i] != w[j] {
				return false
			}
			i++
			j++
		}

		if j == len(w) {
			wi++
			j = 0
		}
	}

	return i == len(s) && wi >= 1 && j == 0
}
