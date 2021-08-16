package main

// Number of Matching Subsequences
//
// Medium
//
// https://leetcode.com/problems/number-of-matching-subsequences/
//
// Given a string `s` and an array of strings `words`, return _the number of_
// `words[i]` _that is a subsequence of_ `s`.
//
// A **subsequence** of a string is a new string generated from the original
// string with some characters (can be none) deleted without changing the
// relative order of the remaining characters.
//
// - For example, `"ace"` is a subsequence of `"abcde"`.
//
// **Example 1:**
//
// ```
// Input: s = "abcde", words = ["a","bb","acd","ace"]
// Output: 3
// Explanation: There are three strings in words that are a subsequence of s:
// "a", "acd", "ace".
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
// Output: 2
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 5 * 104`
// - `1 <= words.length <= 5000`
// - `1 <= words[i].length <= 50`
// - `s` and `words[i]` consist of only lowercase English letters.
func numMatchingSubseq(s string, words []string) int {
	seen := make(map[string]int)
	i := 0
	for j := 0; j < len(words); j++ {
		w := words[j]
		seen[w]++
		if seen[w] == 1 {
			words[i] = w
			i++
		}
	}

	words = words[:i]
	ys := make([]int, len(words))

	for x := 0; x < len(s); x++ {
		a := s[x]
		for i, y := range ys {
			w := words[i]
			if y < len(w) && w[y] == a {
				ys[i]++
			}
		}
	}

	r := 0
	for i, y := range ys {
		if y >= len(words[i]) {
			r += seen[words[i]]
		}
	}

	return r
}

// TODO (tai): slow, 564 ms, faster than 35.14%
