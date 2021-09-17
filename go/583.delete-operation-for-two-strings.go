package main

// Delete Operation for Two Strings
//
// Medium
//
// https://leetcode.com/problems/delete-operation-for-two-strings/
//
// Given two strings `word1` and `word2`, return _the minimum number of
// **steps** required to make_ `word1` _and_ `word2` _the same_.
//
// In one **step**, you can delete exactly one character in either string.
//
// **Example 1:**
//
// ```
// Input: word1 = "sea", word2 = "eat"
// Output: 2
// Explanation: You need one step to make "sea" to "ea" and another step to make
// "eat" to "ea".
//
// ```
//
// **Example 2:**
//
// ```
// Input: word1 = "leetcode", word2 = "etco"
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= word1.length, word2.length <= 500`
// - `word1` and `word2` consist of only lowercase English letters.
func minDistance(word1 string, word2 string) int {
	n := longestCommonSubsequence(word1, word2)
	return len(word1) + len(word2) - 2*n
}
