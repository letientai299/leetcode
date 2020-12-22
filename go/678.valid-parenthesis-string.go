package main

/*
 * @lc app=leetcode id=678 lang=golang
 *
 * [678] Valid Parenthesis String
 *
 * https://leetcode.com/problems/valid-parenthesis-string/description/
 *
 * algorithms
 * Medium (33.51%)
 * Total Accepted:    124K
 * Total Submissions: 394.3K
 * Testcase Example:  '"()"'
 *
 *
 * Given a string containing only three types of characters: '(', ')' and '*',
 * write a function to check whether this string is valid. We define the
 * validity of a string by these rules:
 *
 * Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 * Any right parenthesis ')' must have a corresponding left parenthesis '('.
 * Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 * '*' could be treated as a single right parenthesis ')' or a single left
 * parenthesis '(' or an empty string.
 * An empty string is also valid.
 *
 *
 *
 * Example 1:
 *
 * Input: "()"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "(*)"
 * Output: True
 *
 *
 *
 * Example 3:
 *
 * Input: "(*))"
 * Output: True
 *
 *
 *
 * Note:
 *
 * The string size will be in the range [1, 100].
 *
 *
 */
func checkValidString(s string) bool {
	valid := func(i, diff int) bool {
		cnt := 0
		stars := 0
		open, close := '(', ')'
		if diff < 0 {
			open, close = close, open
		}
		for ; 0 <= i && i < len(s); i += diff {
			switch int32(s[i]) {
			case open:
				cnt++
			case close:
				if cnt > 0 {
					cnt--
					continue
				}
				stars--
				if stars < 0 {
					return false
				}
			case '*':
				stars++
			}
		}
		return true
	}

	return valid(0, 1) && valid(len(s)-1, -1)
}
