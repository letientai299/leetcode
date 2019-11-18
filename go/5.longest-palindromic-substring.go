package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (28.10%)
 * Total Accepted:    690.4K
 * Total Submissions: 2.5M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */
func longestPalindrome(str string) string {
	// Manacher’s Algorithm
	// https://medium.com/hackernoon/manachers-algorithm-explained-longest-palindromic-substring-22cb27a5e96f

	// build another string with a special character inserted in betwwen old
	// string characters, to make the algorithm work.
	var sb strings.Builder
	for _, c := range str {
		sb.WriteRune('|')
		sb.WriteRune(c)
	}

	sb.WriteRune('|')
	s := sb.String()

	// prepare the dynamic memo
	memo := make([]int, len(s))
	center := 0 // center of the longest palindrome till now
	right := 0  // right boundary of the longest palindrome till now

	for i := 0; i < len(s); i++ {
		mirror := 2*center - i

		// mirror can't be negative while it is smaller than right
		if i < right {
			// this logic really took some time to understand!
			memo[i] = memo[mirror]
			if memo[i] > right-i {
				memo[i] = right - i
			}
		}

		for x := i + memo[i] + 1; x < len(s); x++ {
			y := 2*i - x
			if y < 0 {
				break
			}

			if s[x] != s[y] {
				break
			}

			memo[i]++
		}

		if memo[i] > right-center {
			center = i
			right = i + memo[i]
		}
	}

	var result strings.Builder

	for i := 2*center - right; i <= right; i++ {
		if s[i] != '|' {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}
