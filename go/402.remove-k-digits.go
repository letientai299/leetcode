package main

/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 *
 * https://leetcode.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (27.22%)
 * Total Accepted:    165.7K
 * Total Submissions: 580.2K
 * Testcase Example:  '"1432219"\n3'
 *
 * Given a non-negative integer num represented as a string, remove k digits
 * from the number so that the new number is the smallest possible.
 *
 *
 * Note:
 *
 * The length of num is less than 10002 and will be ≥ k.
 * The given num does not contain any leading zero.
 *
 *
 *
 *
 * Example 1:
 *
 * Input: num = "1432219", k = 3
 * Output: "1219"
 * Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219
 * which is the smallest.
 *
 *
 *
 * Example 2:
 *
 * Input: num = "10200", k = 1
 * Output: "200"
 * Explanation: Remove the leading 1 and the number is 200. Note that the
 * output must not contain leading zeroes.
 *
 *
 *
 * Example 3:
 *
 * Input: num = "10", k = 2
 * Output: "0"
 * Explanation: Remove all the digits from the number and it is left with
 * nothing which is 0.
 *
 *
 */
func removeKdigits(s string, k int) string {
	n := len(s)
	if k >= n || n == 0 {
		return "0"
	}

	var buf []byte

	j := 1
	last := s[0]
	for ; k > 0 && j < len(s); j++ {
		cur := s[j]
		if last <= cur {
			buf = append(buf, last)
			last = cur
		} else {
			if len(buf) != 0 {
				j--
			}
			if len(buf) > 0 {
				last = buf[len(buf)-1]
				buf = buf[:len(buf)-1]
			} else {
				last = cur
			}
			k--
		}
	}

	buf = append(buf, last)
	buf = append(buf, s[j:]...)
	for len(buf) > 0 && buf[0] == '0' {
		buf = buf[1:]
	}
	if k >= len(buf) {
		return "0"
	}

	return string(buf)[:len(buf)-k]
}
