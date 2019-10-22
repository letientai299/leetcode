package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.00%)
 * Total Accepted:    1.1M
 * Total Submissions: 4M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// DP problem, memo contains length of longest substring that end at the index.
	memo := make([]int, len(s))
	memo[0] = 1

	// contains last appearance of a character
	lasts := make(map[uint8]int, len(s))
	lasts[s[0]] = 0
	longest := memo[0]

	// the simplified logic is:
	for i := 1; i < len(s); i++ {
		c := s[i]
		last, ok := lasts[c]
		if !ok { // c is not appeared yet
			last = -1
		}
		lasts[c] = i

		memo[i] = i - last
		if memo[i] > memo[i-1]+1 {
			memo[i] = memo[i-1] + 1
		}

		if memo[i] > longest {
			longest = memo[i]
		}
	}

	return longest
}
