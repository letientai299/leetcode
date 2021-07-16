package main

// Check if One String Swap Can Make Strings Equal
//
// Easy
//
// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
//
// You are given two strings `s1` and `s2` of equal length. A **string swap** is
// an operation where you choose two indices in a string (not necessarily
// different) and swap the characters at these indices.
//
// Return `true` _if it is possible to make both strings equal by performing
// **at most one string swap** on **exactly one** of the strings._ Otherwise,
// return `false`.
//
// **Example 1:**
//
// ```
// Input: s1 = "bank", s2 = "kanb"
// Output: true
// Explanation: For example, swap the first character with the last character of
// s2 to make "bank".
//
// ```
//
// **Example 2:**
//
// ```
// Input: s1 = "attack", s2 = "defend"
// Output: false
// Explanation: It is impossible to make them equal with one string swap.
//
// ```
//
// **Example 3:**
//
// ```
// Input: s1 = "kelb", s2 = "kelb"
// Output: true
// Explanation: The two strings are already equal, so no string swap operation
// is required.
//
// ```
//
// **Example 4:**
//
// ```
// Input: s1 = "abcd", s2 = "dcba"
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= s1.length, s2.length <= 100`
// - `s1.length == s2.length`
// - `s1` and `s2` consist of only lowercase English letters.
func areAlmostEqual(s1 string, s2 string) bool {
	j := -1
	count := 0
	for i, c1 := range s1 {
		c2 := int32(s2[i])
		if c1 == c2 {
			continue
		}

		count++
		if count > 2 {
			return false
		}

		if j == -1 {
			j = i
			continue
		}

		if s2[j] != s1[i] || s2[i] != s1[j] {
			return false
		}
	}

	return count != 1
}
