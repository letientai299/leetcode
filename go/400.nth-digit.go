package main

/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 *
 * https://leetcode.com/problems/nth-digit/description/
 *
 * algorithms
 * Easy (30.16%)
 * Total Accepted:    45.5K
 * Total Submissions: 150.8K
 * Testcase Example:  '3'
 *
 * Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8,
 * 9, 10, 11, ...
 *
 * Note:
 * n is positive and will fit within the range of a 32-bit signed integer (n <
 * 231).
 *
 *
 * Example 1:
 *
 * Input:
 * 3
 *
 * Output:
 * 3
 *
 *
 *
 * Example 2:
 *
 * Input:
 * 11
 *
 * Output:
 * 0
 *
 * Explanation:
 * The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a
 * 0, which is part of the number 10.
 *
 *
 */
func findNthDigit(num int) int {
	n := num
	if n < 10 {
		return n
	}

	nDigit := 1
	base := 10
	requiredDigits := nDigit * (base / 10 * 9)
	for n > requiredDigits {
		n -= requiredDigits
		nDigit++
		base *= 10
		requiredDigits = nDigit * (base / 10 * 9)
	}
	base /= 10
	n--
	base += n / nDigit
	n = nDigit - n%nDigit - 1
	for n > 0 {
		base /= 10
		n--
	}

	return base % 10
}
