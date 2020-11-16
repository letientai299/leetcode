package main

/*
 * @lc app=leetcode id=1118 lang=golang
 *
 * [1118] Number of Days in a Month
 *
 * https://leetcode.com/problems/number-of-days-in-a-month/description/
 *
 * algorithms
 * Easy (56.86%)
 * Total Accepted:    4.5K
 * Total Submissions: 7.9K
 * Testcase Example:  '1992\n7'
 *
 * Given a year Y and a month M, return how many days there are in that
 * month.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: Y = 1992, M = 7
 * Output: 31
 *
 *
 * Example 2:
 *
 *
 * Input: Y = 2000, M = 2
 * Output: 29
 *
 *
 * Example 3:
 *
 *
 * Input: Y = 1900, M = 2
 * Output: 28
 *
 *
 *
 *
 * Note:
 *
 *
 * 1583 <= Y <= 2100
 * 1 <= M <= 12
 *
 *
 */
func numberOfDays(Y int, M int) int {
	switch M {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if (Y%4 == 0 && Y%100 != 0) || (Y%400 == 0) {
			return 29
		}
		return 28
	}
	return 0
}
