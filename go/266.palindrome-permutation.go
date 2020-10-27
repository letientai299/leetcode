package main

/*
 * @lc app=leetcode id=266 lang=golang
 *
 * [266] Palindrome Permutation
 *
 * https://leetcode.com/problems/palindrome-permutation/description/
 *
 * algorithms
 * Easy (60.81%)
 * Total Accepted:    98.9K
 * Total Submissions: 159K
 * Testcase Example:  '"code"'
 *
 * Given a string, determine if a permutation of the string could form a
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "code"
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: "aab"
 * Output: true
 *
 * Example 3:
 *
 *
 * Input: "carerac"
 * Output: true
 *
 */
func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	hasOdd := false
	for _, v := range m {
		if v%2 == 1 {
			if hasOdd {
				return false
			}
			hasOdd = true
		}
	}
	return true
}
