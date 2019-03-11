package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (39.63%)
 * Total Accepted:    262K
 * Total Submissions: 661.2K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 * 6.     312211
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 * Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the
 * count-and-say sequence.
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "1"
 *
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "1211"
 *
 */
func countAndSay(n int) string {
	s := "1"
	for i := 2; i <= n; i++ {
		c := s[0]
		count := 1
		sb := strings.Builder{}
		for j := 1; j < len(s); j++ {
			if s[j] == c {
				count ++
				continue
			}

			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(c)
			c = s[j]
			count = 1
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(c)
		s = sb.String()
	}
	return s
}
