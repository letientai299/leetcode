package main

// Longest Common Subsequence
//
// Medium
//
// https://leetcode.com/problems/longest-common-subsequence/
//
// Given two strings `text1` and `text2`, return _the length of their longest
// **common subsequence**._ If there is no **common subsequence**, return `0`.
//
// A **subsequence** of a string is a new string generated from the original
// string with some characters (can be none) deleted without changing the
// relative order of the remaining characters.
//
// - For example, `"ace"` is a subsequence of `"abcde"`.
//
// A **common subsequence** of two strings is a subsequence that is common to
// both strings.
//
// **Example 1:**
//
// ```
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
// ```
//
// **Example 2:**
//
// ```
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
// ```
//
// **Example 3:**
//
// ```
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
//
// ```
//
// **Constraints:**
//
// - `1 <= text1.length, text2.length <= 1000`
// - `text1` and `text2` consist of only lowercase English characters.
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	dp := make([][]int, 2)

	// much start with 0 (means empty string),
	// otherwise, the logic will be very clumsy
	dp[0] = make([]int, m+1)
	dp[1] = make([]int, m+1)

	for _, y := range text2 {
		for i, x := range text1 {
			if x == y {
				dp[1][i+1] = dp[0][i] + 1
			} else {
				a, b := dp[1][i], dp[0][i+1]
				if a < b {
					a = b
				}
				dp[1][i+1] = a
			}
		}

		dp[0], dp[1] = dp[1], dp[0]
	}

	return dp[0][m]
}
