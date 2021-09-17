package main

// Edit Distance
//
// Hard
//
// https://leetcode.com/problems/edit-distance/
//
// Given two strings `word1` and `word2`, return _the minimum number of
// operations required to convert `word1` to `word2`_.
//
// You have the following three operations permitted on a word:
//
// - Insert a character
// - Delete a character
// - Replace a character
//
// **Example 1:**
//
// ```
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
//
// ```
//
// **Example 2:**
//
// ```
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')
//
// ```
//
// **Constraints:**
//
// - `0 <= word1.length, word2.length <= 500`
// - `word1` and `word2` consist of lowercase English letters.
func minDistance(a string, b string) int {
	m := len(a)
	n := len(b)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, m+1)
	for i := 0; i < m; i++ {
		dp[0][i+1] = i + 1
	}

	dp[1] = make([]int, m+1)

	for j := 0; j < n; j++ {
		y := b[j]
		dp[1][0] = j+1
		for i := 0; i < m; i++ {
			x := a[i]
			cost := 0
			if x != y {
				cost = 1
			}

			dp[1][i+1] = min(
				dp[0][i+1]+1,
				dp[1][i]+1,
				dp[0][i]+cost,
			)
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	return dp[0][m]
}
