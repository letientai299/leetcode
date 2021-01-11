package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (35.90%)
 * Total Accepted:    527.4K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */
func isValid(s string) bool {
	stack := make([]int32, len(s))
	top := 0

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack[top] = c
			top++
			continue
		}

		if top == 0 {
			return false
		}

		if stack[top-1] == '(' && c != ')' {
			return false
		}

		if stack[top-1] == '{' && c != '}' {
			return false
		}

		if stack[top-1] == '[' && c != ']' {
			return false
		}

		stack[top-1] = 0
		top--
	}

	return top == 0
}
