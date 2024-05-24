package main

import (
	"strconv"
	"strings"
)

// Count Days Spent Together
//
// # Easy
//
// https://leetcode.com/problems/count-days-spent-together
//
// Alice and Bob are traveling to Rome for separate business meetings.
//
// You are given 4 strings `arriveAlice`, `leaveAlice`, `arriveBob`, and
// `leaveBob`. Alice will be in the city from the dates `arriveAlice` to
// `leaveAlice` ( **inclusive**), while Bob will be in the city from the dates
// `arriveBob` to `leaveBob` ( **inclusive**). Each will be a 5-character string
// in the format `"MM-DD"`, corresponding to the month and day of the date.
//
// Return _the total number of days that Alice and Bob are in Rome together._
//
// You can assume that all dates occur in the **same** calendar year, which is
// **not** a leap year. Note that the number of days per month can be
// represented as: `[31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]`.
//
// **Example 1:**
//
// ```
// Input: arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16",
// leaveBob = "08-19"
// Output: 3
// Explanation: Alice will be in Rome from August 15 to August 18. Bob will be
// in Rome from August 16 to August 19. They are both in Rome together on August
// 16th, 17th, and 18th, so the answer is 3.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01",
// leaveBob = "12-31"
// Output: 0
// Explanation: There is no day when Alice and Bob are in Rome together, so we
// return 0.
//
// ```
//
// **Constraints:**
//
// - All dates are provided in the format `"MM-DD"`.
// - Alice and Bob's arrival dates are **earlier than or equal to** their
// leaving dates.
// - The given dates are valid dates of a **non-leap** year.
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	m1, d1 := countDaysTogetherParse(arriveAlice)
	m2, d2 := countDaysTogetherParse(leaveAlice)

	m, d := countDaysTogetherParse(arriveBob)
	if m2 < m || m2 == m && d2 < d {
		return 0 // Bob arrives after Alice leaves.
	}

	if m1 < m || m1 == m && d1 < d {
		m1, d1 = m, d
	}

	m, d = countDaysTogetherParse(leaveBob)
	if m1 > m || m1 == m && d1 > d {
		return 0 // Alice arrives after Bob leaves.
	}

	if m2 > m || m2 == m && d2 > d {
		m2, d2 = m, d
	}

	if m1 == m2 {
		return d2 - d1 + 1
	}

	months := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	res := 0
	for i := m1 + 1; i < m2; i++ {
		res += months[i]
	}

	res += months[m1] - d1 + 1
	res += d2
	return res
}

func countDaysTogetherParse(s string) (int, int) {
	month, day, _ := strings.Cut(s, "-")
	m, _ := strconv.Atoi(month)
	d, _ := strconv.Atoi(day)
	return m, d
}
