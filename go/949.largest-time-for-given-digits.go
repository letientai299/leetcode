package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=949 lang=golang
 *
 * [949] Largest Time for Given Digits
 *
 * https://leetcode.com/problems/largest-time-for-given-digits/description/
 *
 * algorithms
 * Easy (34.60%)
 * Total Accepted:    57.4K
 * Total Submissions: 158.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array arr of 4 digits, find the latest 24-hour time that can be
 * made using each digit exactly once.
 *
 * 24-hour times are formatted as "HH:MM", where HH is between 00 and 23, and
 * MM is between 00 and 59. The earliest 24-hour time is 00:00, and the latest
 * is 23:59.
 *
 * Return the latest 24-hour time in "HH:MM" format.  If no valid time can be
 * made, return an empty string.
 *
 *
 * Example 1:
 *
 *
 * Input: A = [1,2,3,4]
 * Output: "23:41"
 * Explanation: The valid 24-hour times are "12:34", "12:43", "13:24", "13:42",
 * "14:23", "14:32", "21:34", "21:43", "23:14", and "23:41". Of these times,
 * "23:41" is the latest.
 *
 *
 * Example 2:
 *
 *
 * Input: A = [5,5,5,5]
 * Output: ""
 * Explanation: There are no valid 24-hour times as "55:55" is not valid.
 *
 *
 * Example 3:
 *
 *
 * Input: A = [0,0,0,0]
 * Output: "00:00"
 *
 *
 * Example 4:
 *
 *
 * Input: A = [0,0,1,0]
 * Output: "10:00"
 *
 *
 *
 * Constraints:
 *
 *
 * arr.length == 4
 * 0 <= arr[i] <= 9
 *
 *
 */

func largestTimeFromDigits(arr []int) string {
	sort.Ints(arr)
	if arr[0] > 2 || (arr[0] == 2 && arr[1] > 3) || // can't build hour
		arr[1] >= 6 { // can't build minute
		return ""
	}
	h := 0
	m := 0
	best := 0
	for i, x := range arr {
		if x > 2 {
			continue
		}

		h = x
		for j, y := range arr {
			if i == j || (h == 2 && y >= 4) {
				continue
			}
			h = h*10 + y
			for k, z := range arr {
				if k == i || k == j || z >= 6 {
					continue
				}
				m = z

				for l, u := range arr {
					if l == i || l == j || l == k {
						continue
					}
					m = 10*m + u

					if best < h*60+m {
						best = h*60 + m
					}
				}
				m /= 10
			}
			h /= 10
		}
		h = 0
	}

	if best == 0 && arr[3] != 0 {
		return ""
	}

	return fmt.Sprintf("%02d:%02d", best/60, best%60)
}
