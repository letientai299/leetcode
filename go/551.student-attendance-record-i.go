package main

/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 *
 * https://leetcode.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (45.26%)
 * Total Accepted:    51.9K
 * Total Submissions: 114.6K
 * Testcase Example:  '"PPALLP"'
 *
 * You are given a string representing an attendance record for a student. The
 * record only contains the following three characters:
 *
 *
 *
 * 'A' : Absent.
 * 'L' : Late.
 * ⁠'P' : Present.
 *
 *
 *
 *
 * A student could be rewarded if his attendance record doesn't contain more
 * than one 'A' (absent) or more than two continuous 'L' (late).
 *
 * You need to return whether the student could be rewarded according to his
 * attendance record.
 *
 * Example 1:
 *
 * Input: "PPALLP"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "PPALLL"
 * Output: False
 *
 *
 *
 *
 *
 */
func checkRecord(s string) bool {
	continuouslyLateCount := 0
	absentCount := 0
	for _, c := range s {
		switch c {
		case 'A':
			absentCount++
			continuouslyLateCount = 0
		case 'L':
			continuouslyLateCount++
		default:
			continuouslyLateCount = 0
		}

		if continuouslyLateCount > 2 {
			return false
		}
		if absentCount > 1 {
			return false
		}
	}

	return true
}
