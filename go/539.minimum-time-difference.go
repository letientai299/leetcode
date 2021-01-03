package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=539 lang=golang
 *
 * [539] Minimum Time Difference
 *
 * https://leetcode.com/problems/minimum-time-difference/description/
 *
 * algorithms
 * Medium (49.87%)
 * Total Accepted:    57.1K
 * Total Submissions: 109.6K
 * Testcase Example:  '["23:59","00:00"]'
 *
 * Given a list of 24-hour clock time points in "HH:MM" format, return the
 * minimum minutes difference between any two time-points in the list.
 *
 * Example 1:
 * Input: timePoints = ["23:59","00:00"]
 * Output: 1
 * Example 2:
 * Input: timePoints = ["00:00","23:59","00:00"]
 * Output: 0
 *
 *
 * Constraints:
 *
 *
 * 2 <= timePoints <= 2 * 10^4
 * timePoints[i] is in the format "HH:MM".
 *
 *
 */
func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	seconds := make([]int, n)
	parse := func(s string) int {
		h := int(s[0]-'0')*10 + int(s[1]-'0')
		m := int(s[3]-'0')*10 + int(s[4]-'0')
		return h*60 + m
	}

	for i, s := range timePoints {
		seconds[i] = parse(s)
	}
	const (
		half = 720
		all  = half * 2
	)
	m := all

	sort.Ints(seconds)
	for i := 0; i < n-1; i++ {
		a, b := seconds[i], seconds[i+1]
		v := b - a

		if v == 0 {
			return 0
		}

		if all-v < v {
			v = all - v
		}

		if v < m {
			m = v
		}

		if all-v < m {
			m = all - v
		}
	}

	v := seconds[n-1] - seconds[0]
	if v > all-v {
		v = all - v
	}

	if m > v {
		m = v
	}

	if m <= half {
		return  m
	}

	return m % half
}
