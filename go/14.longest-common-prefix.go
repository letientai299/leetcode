package main

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.00%)
 * Total Accepted:    417.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// take the first string
	s := strs[0]
	if len(strs) == 1 {
		return s
	}

	ci := 0
	for ; ci < len(s); ci++ {
		c := s[ci]
		for i := 1; i < len(strs); i++ {
			current := strs[i]
			if ci >= len(current) || c != current[ci] {
				return s[0:ci]
			}
		}
	}
	return s[0:ci]
}
