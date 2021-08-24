package main

// Unique Length-3 Palindromic Subsequences
//
// Medium
//
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
//
// Given a string `s`, return _the number of **unique palindromes of length
// three** that are a **subsequence** of_ `s`.
//
// Note that even if there are multiple ways to obtain the same subsequence, it
// is still only counted **once**.
//
// A **palindrome** is a string that reads the same forwards and backwards.
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
// Input: s = "aabca"
// Output: 3
// Explanation: The 3 palindromic subsequences of length 3 are:
// - "aba" (subsequence of "aabca")
// - "aaa" (subsequence of "aabca")
// - "aca" (subsequence of "aabca")
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "adc"
// Output: 0
// Explanation: There are no palindromic subsequences of length 3 in "adc".
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "bbcbaba"
// Output: 4
// Explanation: The 4 palindromic subsequences of length 3 are:
// - "bbb" (subsequence of "bbcbaba")
// - "bcb" (subsequence of "bbcbaba")
// - "bab" (subsequence of "bbcbaba")
// - "aba" (subsequence of "bbcbaba")
//
// ```
//
// **Constraints:**
//
// - `3 <= s.length <= 105`
// - `s` consists of only lowercase English letters.
func countPalindromicSubsequence(s string) int {
	n := len(s)
	pre := 0

	counts := make([][]int, n)
	pair := make([][]int, 26)

	for i, c := range s {
		c -= 'a'
		if pair[c] == nil {
			pair[c] = []int{i, i}
			pre++
		} else {
			pair[c][1] = i
		}
		counts[i] = make([]int, 26)

		if i > 0 {
			copy(counts[i], counts[i-1])
		}

		counts[i][c]++
	}

	r := 0
	for c, p := range pair {
		if p == nil || p[0]+1 >= p[1] {
			continue
		}

		left, right := counts[p[0]], counts[p[1]]
		for i, v := range right {
			v -= left[i]
			if i == c {
				v--
			}

			if v > 0 {
				r++
			}
		}
	}

	return r
}

// TODO (tai): can be faster
