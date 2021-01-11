package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.96%)
 * Total Accepted:    216.1K
 * Total Submissions: 586.7K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string s containing only digits, return all possible valid IP
 * addresses that can be obtained from s. You can return them in any order.
 *
 * A valid IP address consists of exactly four integers, each integer is
 * between 0 and 255, separated by single dots and cannot have leading zeros.
 * For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and
 * "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP
 * addresses.
 *
 *
 * Example 1:
 * Input: s = "25525511135"
 * Output: ["255.255.11.135","255.255.111.35"]
 * Example 2:
 * Input: s = "0000"
 * Output: ["0.0.0.0"]
 * Example 3:
 * Input: s = "1111"
 * Output: ["1.1.1.1"]
 * Example 4:
 * Input: s = "010010"
 * Output: ["0.10.0.10","0.100.1.0"]
 * Example 5:
 * Input: s = "101023"
 * Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 3000
 * s consists of digits only.
 *
 *
 */
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	toIP := func(b []int) string {
		var sb strings.Builder
		for i, x := range b {
			if i != 0 {
				sb.WriteByte('.')
			}
			sb.WriteString(strconv.Itoa(x))
		}
		return sb.String()
	}

	buf := make([]int, 0, 4)
	res := make([]string, 0, 10)
	var find func(i int, cur int)
	find = func(i int, cur int) {
		if cur > 255 {
			return // invalid
		}

		if i >= len(s) {
			if len(buf) != 3 {
				return // not enough to make an IP
			}
			buf = append(buf, cur)
			ip := toIP(buf)
			if len(ip)-len(s) != 3 {
				return
			}

			res = append(res, ip)
			return
		}

		x := int(s[i] - '0')
		l := len(buf)
		buf = append(buf, cur)
		find(i+1, x)
		buf = buf[:l]
		find(i+1, cur*10+x)
	}

	find(1, int(s[0]-'0'))
	return res
}
