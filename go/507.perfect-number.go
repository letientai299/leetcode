package main

import "math"

/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 *
 * https://leetcode.com/problems/perfect-number/description/
 *
 * algorithms
 * Easy (34.13%)
 * Total Accepted:    39.5K
 * Total Submissions: 115.7K
 * Testcase Example:  '28'
 *
 * We define the Perfect Number is a positive integer that is equal to the sum
 * of all its positive divisors except itself.
 *
 * Now, given an integer n, write a function that returns true when it is a
 * perfect number and false when it is not.
 *
 *
 * Example:
 *
 * Input: 28
 * Output: True
 * Explanation: 28 = 1 + 2 + 4 + 7 + 14
 *
 *
 *
 * Note:
 * The input number n will not exceed 100,000,000. (1e8)
 *
 */
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	i := 2
	sum := 1
	for i <= int(math.Sqrt(float64(num))) && sum <= num {
		if num%i == 0 {
			sum += i
			sum += num / i
		}
		i++
	}
	return sum == num
}
