package main

// Split a String Into the Max Number of Unique Substrings
//
// Medium
//
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
//
// Given a string `s`, return _the maximum number of unique substrings that the
// given string can be split into_.
//
// You can split string `s` into any list of **non-empty substrings**, where the
// concatenation of the substrings forms the original string. However, you must
// split the substrings such that all of them are **unique**.
//
// A **substring** is a contiguous sequence of characters within a string.
//
// **Example 1:**
//
// ```
// Input: s = "ababccc"
// Output: 5
// Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc'].
// Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a'
// and 'b' multiple times.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "aba"
// Output: 2
// Explanation: One way to split maximally is ['a', 'ba'].
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "aa"
// Output: 1
// Explanation: It is impossible to split the string any further.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 16`
//
// - `s` contains only lower case English letters.
func maxUniqueSplit(s string) int {
	best := 0
	seen := make(map[string]struct{})
	var walk func(l, r int)

	walk = func(l, r int) {
		sub := s[l:r]
		seen[sub] = struct{}{}

		if r >= len(s) {
			if best < len(seen) {
				best = len(seen)
			}
			delete(seen, sub)
			return
		}

		l = r
		for i := r + 1; i <= len(s); i++ {
			next := s[l:i]
			if _, ok := seen[next]; !ok {
				walk(l, i)
			}
		}
		delete(seen, sub)
	}

	for i := 1; i <= len(s); i++ {
		walk(0, i)
	}
	return best
}

// TODO (tai): slow, 55ms, 17.66%
