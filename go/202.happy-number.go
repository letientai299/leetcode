package main

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (44.50%)
 * Total Accepted:    218.5K
 * Total Submissions: 491.1K
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number is "happy".
 *
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 *
 * Example:
 *
 *
 * Input: 19
 * Output: true
 * Explanation:
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 */
func isHappy(n int) bool {
	sumSquareDigit := func(n int) int {
		s := 0

		for n > 0 {
			d := n % 10
			s += d * d
			n /= 10
		}

		return s
	}

	m := make(map[int]bool)
	for {
		x := sumSquareDigit(n)
		if m[x] == false {
			m[x] = true
			n = x
		} else {
			return x == 1
		}
	}
}
