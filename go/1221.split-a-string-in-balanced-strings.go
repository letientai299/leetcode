package main

/*
 * @lc app=leetcode id=1221 lang=golang
 *
 * [1221] Split a String in Balanced Strings
 *
 * https://leetcode.com/problems/split-a-string-in-balanced-strings/description/
 *
 * algorithms
 * Easy (78.96%)
 * Total Accepted:    111.4K
 * Total Submissions: 133K
 * Testcase Example:  '"RLRRLLRLRL"'
 *
 * Balanced strings are those who have equal quantity of 'L' and 'R'
 * characters.
 *
 * Given a balanced string s split it in the maximum amount of balanced
 * strings.
 *
 * Return the maximum amount of splitted balanced strings.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "RLRRLLRLRL"
 * Output: 4
 * Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring
 * contains same number of 'L' and 'R'.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "RLLLLRRRLR"
 * Output: 3
 * Explanation: s can be split into "RL", "LLLRRR", "LR", each substring
 * contains same number of 'L' and 'R'.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "LLLLRRRR"
 * Output: 1
 * Explanation: s can be split into "LLLLRRRR".
 *
 *
 * Example 4:
 *
 *
 * Input: s = "RLRRRLLRLL"
 * Output: 2
 * Explanation: s can be split into "RL", "RRRLLRLL", since each substring
 * contains an equal number of 'L' and 'R'
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s[i] = 'L' or 'R'
 *
 *
 */

func balancedStringSplit(s string) int {
	n := 0
	dep := 0
	c := 0
	for _, x := range s {
		if x == 'R' {
			dep++
		} else {
			dep--
		}
		c++
		if c != 0 && dep == 0 {
			n++
		}
	}
	return n
}
