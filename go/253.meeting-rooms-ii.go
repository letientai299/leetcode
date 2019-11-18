package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=253 lang=golang
 *
 * [253] Meeting Rooms II
 *
 * https://leetcode.com/problems/meeting-rooms-ii/description/
 *
 * algorithms
 * Medium (43.86%)
 * Total Accepted:    208.5K
 * Total Submissions: 472.8K
 * Testcase Example:  '[[0,30],[5,10],[15,20]]'
 *
 * Given an array of meeting time intervals consisting of start and end times
 * [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms
 * required.
 *
 * Example 1:
 *
 *
 * Input: [[0, 30],[5, 10],[15, 20]]
 * Output: 2
 *
 * Example 2:
 *
 *
 * Input: [[7,10],[2,4]]
 * Output: 1
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */
func minMeetingRooms(rev [][]int) int {
	n := len(rev)
	starts := make([]int, 0, n)
	ends := make([]int, 0, n)

	for _, r := range rev {
		starts = append(starts, r[0])
		ends = append(ends, r[1])
	}

	sort.Ints(starts)
	sort.Ints(ends)

	si := 0
	ei := 0
	m := 0
	cur := 0

	// 2 7
	// 3 10
	for si < n && ei < n {
		if starts[si] < ends[ei] {
			cur++
			if m < cur {
				m = cur
			}

			si++
			continue
		}

		cur--
		ei++
	}

	return m
}
