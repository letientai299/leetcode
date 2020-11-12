package main

/*
 * @lc app=leetcode id=1180 lang=golang
 *
 * [1180] Count Substrings with Only One Distinct Letter
 *
 * https://leetcode.com/problems/count-substrings-with-only-one-distinct-letter/description/
 *
 * algorithms
 * Easy (76.66%)
 * Total Accepted:    11.4K
 * Total Submissions: 14.8K
 * Testcase Example:  '"aaaba"'
 *
 * Given a string S, return the number of substrings that have only one
 * distinct letter.
 *
 *
 * Example 1:
 *
 *
 * Input: S = "aaaba"
 * Output: 8
 * Explanation: The substrings with one distinct letter are "aaa", "aa", "a",
 * "b".
 * "aaa" occurs 1 time.
 * "aa" occurs 2 times.
 * "a" occurs 4 times.
 * "b" occurs 1 time.
 * So the answer is 1 + 2 + 4 + 1 = 8.
 *
 *
 * Example 2:
 *
 *
 * Input: S = "aaaaaaaaaa"
 * Output: 55
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= S.length <= 1000
 * S[i] consists of only lowercase English letters.
 *
 *
 */
func countLetters(s string) int {
	if len(s) == 1 {
		return 1
	}
	cnt := 0
	dup := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dup++
			continue
		}
		cnt += dup * (dup + 1) / 2
		dup = 1
	}

	cnt += dup * (dup + 1) / 2
	return cnt
}
