package main

// Regular Expression Matching
//
// # Hard
//
// https://leetcode.com/problems/regular-expression-matching
//
// Given an input string `s` and a pattern `p`, implement regular expression
// matching with support for `'.'` and `'*'` where:
//
// - `'.'` Matches any single character.​​​​
// - `'*'` Matches zero or more of the preceding element.
//
// The matching should cover the **entire** input string (not partial).
//
// **Example 1:**
//
// ```
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore,
// by repeating 'a' once, it becomes "aa".
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 20`
// - `1 <= p.length <= 20`
// - `s` contains only lowercase English letters.
// - `p` contains only lowercase English letters, `'.'`, and `'*'`.
// - It is guaranteed for each appearance of the character `'*'`, there will be
// a previous valid character to match.
func isMatch(s string, p string) bool {
	if len(s) == 0 {
		for i := 0; i < len(p); i += 2 {
			if i+1 >= len(p) || p[i+1] != '*' {
				return false
			}
		}
		return true
	}

	if len(p) == 0 {
		return false
	}

	if len(p) == 1 {
		return len(s) == 1 && (p == s || p == ".")
	}

	if p[1] != '*' {
		return (p[0] == '.' || s[0] == p[0]) && isMatch(s[1:], p[1:])
	}

	for i := 0; i <= len(s); i++ {
		if isMatch(s[i:], p[2:]) {
			return true
		}
		if i < len(s) && p[0] != '.' && s[i] != p[0] {
			break
		}
	}
	return false
}

// TODO (tai): think about DP solution
