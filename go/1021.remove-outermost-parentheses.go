package main

import "strings"

/*
 * @lc app=leetcode id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 *
 * https://leetcode.com/problems/remove-outermost-parentheses/description/
 *
 * algorithms
 * Easy (75.80%)
 * Total Accepted:    109.2K
 * Total Submissions: 139.1K
 * Testcase Example:  '"(()())(())"'
 *
 * A valid parentheses string is either empty (""), "(" + A + ")", or A + B,
 * where A and B are valid parentheses strings, and + represents string
 * concatenation.  For example, "", "()", "(())()", and "(()(()))" are all
 * valid parentheses strings.
 *
 * A valid parentheses string S is primitive if it is nonempty, and there does
 * not exist a way to split it into S = A+B, with A and B nonempty valid
 * parentheses strings.
 *
 * Given a valid parentheses string S, consider its primitive decomposition: S
 * = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.
 *
 * Return S after removing the outermost parentheses of every primitive string
 * in the primitive decomposition of S.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "(()())(())"
 * Output: "()()()"
 * Explanation:
 * The input string is "(()())(())", with primitive decomposition "(()())" +
 * "(())".
 * After removing outer parentheses of each part, this is "()()" + "()" =
 * "()()()".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "(()())(())(()(()))"
 * Output: "()()()()(())"
 * Explanation:
 * The input string is "(()())(())(()(()))", with primitive decomposition
 * "(()())" + "(())" + "(()(()))".
 * After removing outer parentheses of each part, this is "()()" + "()" +
 * "()(())" = "()()()()(())".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "()()"
 * Output: ""
 * Explanation:
 * The input string is "()()", with primitive decomposition "()" + "()".
 * After removing outer parentheses of each part, this is "" + "" = "".
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * S.length <= 10000
 * S[i] is "(" or ")"
 * S is a valid parentheses string
 *
 *
 *
 *
 *
 *
 *
 */

func removeOuterParentheses(s string) string {
	var sb strings.Builder
	var buf []byte
	var n int
	for _, x := range s {
		if x == '(' {
			buf = append(buf, byte(x))
			n++
			continue
		}

		n--
		buf = append(buf, byte(x))
		if n == 0 {
			sb.Write(buf[1 : len(buf)-1])
			buf = buf[:0]
		}
	}

	return sb.String()
}
