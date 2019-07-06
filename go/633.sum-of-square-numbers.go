package main

import (
	"math"
)

/*
 * @lc app=leetcode id=633 lang=golang
 *
 * [633] Sum of Square Numbers
 *
 * https://leetcode.com/problems/sum-of-square-numbers/description/
 *
 * algorithms
 * Easy (32.70%)
 * Total Accepted:    46.1K
 * Total Submissions: 141.2K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer c, your task is to decide whether there're two
 * integers a and b such that a2 + b2 = c.
 *
 * Example 1:
 *
 *
 * Input: 5
 * Output: True
 * Explanation: 1 * 1 + 2 * 2 = 5
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: False
 *
 *
 *
 *
 */
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	if c <= 2 {
		return true
	}

	isSquare := func(x int) bool {
		sq := int(math.Floor(math.Sqrt(float64(x))))
		return sq*sq == x
	}

	bound := int(math.Floor(math.Sqrt(float64(c))))
	for i := 0; i <= bound; i++ {
		if isSquare(c - i*i) {
			return true
		}
	}

	return false
}
