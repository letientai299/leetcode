package main

/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 *
 * https://leetcode.com/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (46.81%)
 * Total Accepted:    249K
 * Total Submissions: 533.1K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * Given two strings S and T, return if they are equal when both are typed into
 * empty text editors. # means a backspace character.
 *
 * Note that after backspacing an empty text, the text will continue empty.
 *
 *
 * Example 1:
 *
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 * Explanation: Both S and T become "".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 * Explanation: Both S and T become "c".
 *
 *
 *
 * Example 4:
 *
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 *
 *
 * Follow up:
 *
 *
 * Can you solve it in O(N) time and O(1) space?
 *
 *
 *
 *
 *
 *
 */
func backspaceCompare(s string, t string) bool {
	a := make([]byte, 0, len(s))
	b := make([]byte, 0, len(t))
	for _, x := range s {
		if x != '#' {
			a = append(a, byte(x))
			continue
		}
		if len(a) != 0 {
			a = a[:len(a)-1]
		}
	}

	for _, x := range t {
		if x != '#' {
			b = append(b, byte(x))
			continue
		}
		if len(b) != 0 {
			b = b[:len(b)-1]
		}
	}

	return string(a) == string(b)
}
