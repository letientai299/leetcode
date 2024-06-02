package main

// Minimum Deletions to Make String Balanced
//
// # Medium
//
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced
//
// You are given a string `s` consisting only of characters `'a'` and `'b'`​​​​.
//
// You can delete any number of characters in `s` to make `s` **balanced**. `s`
// is **balanced** if there is no pair of indices `(i,j)` such that `i < j` and
// `s[i] = 'b'` and `s[j]= 'a'`.
//
// Return _the **minimum** number of deletions needed to make_ `s`
// _**balanced**_.
//
// **Example 1:**
//
// ```
// Input: s = "aababbab"
// Output: 2
// Explanation: You can either:
// Delete the characters at 0-indexed positions 2 and 6 ("aababbab" ->
// "aaabbb"), or
// Delete the characters at 0-indexed positions 3 and 6 ("aababbab" ->
// "aabbbb").
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "bbaaaaabb"
// Output: 2
// Explanation: The only solution is to delete the first two characters.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 105`
// - `s[i]` is `'a'` or `'b'`​​.
func minimumDeletions(s string) int {
	n := len(s)
	as := make([]int, n)
	if s[0] == 'a' {
		as[0] = 1
	}

	for i := 1; i < n; i++ {
		if s[i] == 'a' {
			as[i] = as[i-1] + 1
		} else {
			as[i] = as[i-1]
		}
	}

	b := 0
	best := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'b' {
			b++
		}

		best = max(best, b+as[i])
	}

	return n - best
}
