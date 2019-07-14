package main

/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (34.40%)
 * Total Accepted:    79.2K
 * Total Submissions: 230K
 * Testcase Example:  '"aba"'
 *
 *
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 *
 *
 * Example 1:
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:
 *
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 *
 *
 */
func validPalindrome(s string) bool {
	check := func(left, right int) bool {
		for left < right {
			if s[left] == s[right] {
				left++
				right--
				continue
			}
			return false
		}
		return true
	}

	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}

		return check(left+1, right) || check(left, right-1)
	}

	return true
}
