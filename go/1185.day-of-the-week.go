package main

import "time"

/*
 * @lc app=leetcode id=1185 lang=golang
 *
 * [1185] Day of the Week
 *
 * https://leetcode.com/problems/day-of-the-week/description/
 *
 * algorithms
 * Easy (64.24%)
 * Total Accepted:    26.7K
 * Total Submissions: 42.6K
 * Testcase Example:  '31\n8\n2019'
 *
 * Given a date, return the corresponding day of the week for that date.
 *
 * The input is given as three integers representing the day, month and year
 * respectively.
 *
 * Return the answer as one of the following values {"Sunday", "Monday",
 * "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.
 *
 *
 * Example 1:
 *
 *
 * Input: day = 31, month = 8, year = 2019
 * Output: "Saturday"
 *
 *
 * Example 2:
 *
 *
 * Input: day = 18, month = 7, year = 1999
 * Output: "Sunday"
 *
 *
 * Example 3:
 *
 *
 * Input: day = 15, month = 8, year = 1993
 * Output: "Sunday"
 *
 *
 *
 * Constraints:
 *
 *
 * The given dates are valid dates between the years 1971 and 2100.
 *
 *
 */
func dayOfTheWeek(day int, month int, year int) string {
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return d.Weekday().String()
}
