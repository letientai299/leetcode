package main

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.74%)
 * Total Accepted:    360.3K
 * Total Submissions: 883.9K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */
func plusOne(digits []int) []int {
	n := len(digits)
	var i int
	for i = n - 1; i >= 0; i -= 1 {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}
	if i == -1 {
		newDigits := make([]int, len(digits)+1)
		newDigits[0] = 1
		return newDigits
	}
	digits[i-1] += 1
	return digits
}
