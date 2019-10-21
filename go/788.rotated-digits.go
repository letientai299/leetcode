package main

/*
 * @lc app=leetcode id=788 lang=golang
 *
 * [788] Rotated Digits
 *
 * https://leetcode.com/problems/rotated-digits/description/
 *
 * algorithms
 * Easy (55.61%)
 * Total Accepted:    35.6K
 * Total Submissions: 64K
 * Testcase Example:  '10'
 *
 * X is a good number if after rotating each digit individually by 180 degrees,
 * we get a valid number that is different from X.  Each digit must be rotated
 * - we cannot choose to leave it alone.
 *
 * A number is valid if each digit remains a digit after rotation. 0, 1, and 8
 * rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each
 * other, and the rest of the numbers do not rotate to any other number and
 * become invalid.
 *
 * Now given a positive number N, how many numbers X from 1 to N are good?
 *
 *
 * Example:
 * Input: 10
 * Output: 4
 * Explanation:
 * There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
 * Note that 1 and 10 are not good numbers, since they remain unchanged after
 * rotating.
 *
 *
 * Note:
 *
 *
 * N  will be in range [1, 10000].
 *
 *
 */
func rotatedDigits(N int) int {
	isGood := func(x int) bool {
		// a number if good if its digits don't contains 4, 7 and contains either
		// 2, 5, 6 or 9
		hasGoodDigit := false

		for x > 0 {
			d := x % 10
			switch d {
			case 3, 4, 7:
				return false
			case 2, 5, 6, 9:
				hasGoodDigit = true
			}
			x /= 10
		}

		return hasGoodDigit
	}

	count := 0
	for i := 1; i <= N; i++ {
		if isGood(i) {
			count++
		}
	}
	return count
}
