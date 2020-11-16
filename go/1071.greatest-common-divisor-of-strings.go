package main

/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 *
 * https://leetcode.com/problems/greatest-common-divisor-of-strings/description/
 *
 * algorithms
 * Easy (53.80%)
 * Total Accepted:    35.9K
 * Total Submissions: 69.5K
 * Testcase Example:  '"ABCABC"\n"ABC"'
 *
 * For two strings s and t, we say "t divides s" if and only if s = t + ... +
 * t  (t concatenated with itself 1 or more times)
 *
 * Given two strings str1 and str2, return the largest string x such that x
 * divides both str1 and str2.
 *
 *
 * Example 1:
 * Input: str1 = "ABCABC", str2 = "ABC"
 * Output: "ABC"
 * Example 2:
 * Input: str1 = "ABABAB", str2 = "ABAB"
 * Output: "AB"
 * Example 3:
 * Input: str1 = "LEET", str2 = "CODE"
 * Output: ""
 * Example 4:
 * Input: str1 = "ABCDEF", str2 = "ABC"
 * Output: ""
 *
 *
 * Constraints:
 *
 *
 * 1 <= str1.length <= 1000
 * 1 <= str2.length <= 1000
 * str1 and str2 consist of English uppercase letters.
 *
 *
 */

func gcdOfStrings(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	if len(b) == 0 {
		return a
	}

	for i, y := range b {
		x := a[i]
		if x != byte(y) {
			return ""
		}
	}
	a = a[len(b):]
	return gcdOfStrings(b, a)
}
