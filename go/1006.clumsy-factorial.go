package main

/*
 * @lc app=leetcode id=1006 lang=golang
 *
 * [1006] Clumsy Factorial
 *
 * https://leetcode.com/problems/clumsy-factorial/description/
 *
 * algorithms
 * Medium (53.39%)
 * Total Accepted:    14K
 * Total Submissions: 26.1K
 * Testcase Example:  '4'
 *
 * Normally, the factorial of a positive integer n is the product of all
 * positive integers less than or equal to n.  For example, factorial(10) = 10
 * * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1.
 *
 * We instead make a clumsy factorial: using the integers in decreasing order,
 * we swap out the multiply operations for a fixed rotation of operations:
 * multiply (*), divide (/), add (+) and subtract (-) in this order.
 *
 * For example, clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1.  However,
 * these operations are still applied using the usual order of operations of
 * arithmetic: we do all multiplication and division steps before any addition
 * or subtraction steps, and multiplication and division steps are processed
 * left to right.
 *
 * Additionally, the division that we use is floor division such that 10 * 9 /
 * 8 equals 11.  This guarantees the result is an integer.
 *
 * Implement the clumsy function as defined above: given an integer N, it
 * returns the clumsy factorial of N.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 7
 * Explanation: 7 = 4 * 3 / 2 + 1
 *
 *
 * Example 2:
 *
 *
 * Input: 10
 * Output: 12
 * Explanation: 12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 10000
 * -2^31 <= answer <= 2^31 - 1  (The answer is guaranteed to fit within a
 * 32-bit integer.)
 *
 *
 */
func clumsy(n int) int {
	if n < 0 {
		return 0
	}

	switch n {
	case 0, 1, 2:
		return n
	case 3:
		return 6
	case 4:
		return 7
	}

	add := 2
	v := n
	n %= 4
	if n == 3 {
		add = -1
	} else if n == 0 {
		add = 1
	}

	return v + add
}
