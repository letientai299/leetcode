package main

import "sort"

// Non-overlapping Intervals
//
// Medium
//
// https://leetcode.com/problems/non-overlapping-intervals/
//
// Given an array of intervals `intervals` where `intervals[i] = [starti,
// endi]`, return _the minimum number of intervals you need to remove to make
// the rest of the intervals non-overlapping_.
//
// **Example 1:**
//
// ```
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are
// non-overlapping.
//
// ```
//
// **Example 2:**
//
// ```
// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals
// non-overlapping.
//
// ```
//
// **Example 3:**
//
// ```
// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're
// already non-overlapping.
//
// ```
//
// **Constraints:**
//
// - `1 <= intervals.length <= 105`
// - `intervals[i].length == 2`
// - `-5 * 104 <= starti < endi <= 5 * 104`
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[1] < b[1]
	})

	res := 0
	preEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < preEnd {
			res++
		} else {
			preEnd = intervals[i][1]
		}
	}

	return res
}
