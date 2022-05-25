package main

// Minimum Moves to Convert String
//
// Easy
//
// https://leetcode.com/problems/minimum-moves-to-convert-string/
//
// You are given a string `s` consisting of `n` characters which are either
// `'X'` or `'O'`.
//
// A **move** is defined as selecting **three** **consecutive characters** of
// `s` and converting them to `'O'`. Note that if a move is applied to the
// character `'O'`, it will stay the **same**.
//
// Return _the **minimum** number of moves required so that all the characters
// of_ `s` _are converted to_`'O'`.
//
// **Example 1:**
//
// ```
// Input: s = "XXX"
// Output: 1
// Explanation: XXX -> OOO
// We select all the 3 characters and convert them in one move.
//
// ```
//
// **Example 2:**
//
// ```
// 'O''O'
// ```
//
// **Example 3:**
//
// ```
// 'X'ss
// ```
//
// **Constraints:**
//
// - `3 <= s.length <= 1000`
// - `s[i]` is either `'X'` or `'O'`.
func minimumMoves(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'X' {
			r++
			i += 2
		}
	}
	return r
}
