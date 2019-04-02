package main

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 *
 * https://leetcode.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (43.34%)
 * Total Accepted:    88.6K
 * Total Submissions: 204.5K
 * Testcase Example:  '"0"\n"0"'
 *
 * Given two non-negative integers num1 and num2 represented as string, return
 * the sum of num1 and num2.
 *
 * Note:
 *
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */
func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}

	if len(num2) == 0 {
		return num1
	}

	n := len(num1) + 1
	if len(num2) > len(num1) {
		n = len(num2) + 1
	}

	res := make([]byte, n)
	n--
	i, j := len(num1)-1, len(num2)-1
	carry := byte(0)
	for i >= 0 || j >= 0 {
		next := carry
		if i >= 0 {
			next += num1[i] - '0'
			i--
		}

		if j >= 0 {
			next += num2[j] - '0'
			j--
		}

		carry = next / 10
		res[n] = (next % 10) + '0'
		n--
	}
	if carry != 0 {
		res[0] = carry + '0'
		return string(res)
	}

	return string(res[1:])
}
