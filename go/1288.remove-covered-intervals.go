package main

import "sort"

// Remove Covered Intervals
//
// Medium
//
// https://leetcode.com/problems/remove-covered-intervals/
//
// Given an array `intervals` where `intervals[i] = [li, ri]` represent the
// interval `[li, ri)`, remove all intervals that are covered by another
// interval in the list.
//
// The interval `[a, b)` is covered by the interval `[c, d)` if and only if `c
// <= a` and `b <= d`.
//
// Return _the number of remaining intervals_.
//
// **Example 1:**
//
// ```
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
//
// ```
//
// **Example 2:**
//
// ```
// Input: intervals = [[1,4],[2,3]]
// Output: 1
//
// ```
//
// **Example 3:**
//
// ```
// Input: intervals = [[0,10],[5,12]]
// Output: 2
//
// ```
//
// **Example 4:**
//
// ```
// Input: intervals = [[3,10],[4,10],[5,11]]
// Output: 2
//
// ```
//
// **Example 5:**
//
// ```
// Input: intervals = [[1,2],[1,4],[3,4]]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= intervals.length <= 1000`
// - `intervals[i].length == 2`
// - `0 <= li <= ri <= 105`
// - All the given intervals are **unique**.
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] == b[0] {
			return a[1] > b[1] // tricky bit
		}

		return a[0] < b[0]
	})

	res := 1
	cur := intervals[0]
	for _, r := range intervals[1:] {
		if r[1] <= cur[1] {
			continue
		}

		cur = r
		res++
	}
	return res
}
