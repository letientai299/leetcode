package main

/*
 * @lc app=leetcode id=1134 lang=golang
 *
 * [1134] Armstrong Number
 *
 * https://leetcode.com/problems/armstrong-number/description/
 *
 * algorithms
 * Easy (77.74%)
 * Total Accepted:    14.2K
 * Total Submissions: 18.2K
 * Testcase Example:  '153'
 *
 * The k-digit number N is an Armstrong number if and only if the k-th power of
 * each digit sums to N.
 *
 * Given a positive integer N, return true if and only if it is an Armstrong
 * number.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 153
 * Output: true
 * Explanation:
 * 153 is a 3-digit number, and 153 = 1^3 + 5^3 + 3^3.
 *
 *
 * Example 2:
 *
 *
 * Input: 123
 * Output: false
 * Explanation:
 * 123 is a 3-digit number, and 123 != 1^3 + 2^3 + 3^3 = 36.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 10^8
 *
 *
 */
func isArmstrong(a int) bool {
	n := a
	var ds []int
	for n > 0 {
		ds = append(ds, n%10)
		n /= 10
	}

	for _, x := range ds {
		c := 1
		for range ds {
			c *= x
		}
		a -= c
	}
	return a == 0
}
