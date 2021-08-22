package main

import "strings"

// DI String Match
//
// Easy
//
// https://leetcode.com/problems/di-string-match/
//
// A permutation `perm` of `n + 1` integers of all the integers in the range
// `[0, n]` can be represented as a string `s` of length `n` where:
//
// - `s[i] == 'I'` if `perm[i] < perm[i + 1]`, and
// - `s[i] == 'D'` if `perm[i] > perm[i + 1]`.
//
// Given a string `s`, reconstruct the permutation `perm` and return it. If
// there are multiple valid permutations perm, return **any of them**.
//
// **Example 1:**
//
// ```
// Input: s = "IDID"
// Output: [0,4,1,3,2]
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "III"
// Output: [0,1,2,3]
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "DDI"
// Output: [3,2,0,1]
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 105`
// - `s[i]` is either `'I'` or `'D'`.
func diStringMatch(s string) []int {
	n := len(s) + 1
	res := make([]int, n)

	// 2 queues, 1 up (x, x+1, x+2, ...) , 1 down (x-1, x-2, ...).
	// If we see a 'D', append next value to down queue.
	// If we see an 'I', append next to up queue.
	// Now, how long are the queues? Should be numbers of 'I' and 'D'.
	// From this, we will know which x should we start with
	up := strings.Count(s, "D")
	res[0] = up
	down := up
	for i, c := range s {
		if c == 'I' {
			up++
			res[i+1] = up
		} else {
			down--
			res[i+1] = down
		}
	}

	return res
}
