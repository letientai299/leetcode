package main

import "sort"

// Check if All the Integers in a Range Are Covered
//
// Easy
//
// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/
//
// You are given a 2D integer array `ranges` and two integers `left` and
// `right`. Each `ranges[i] = [starti, endi]` represents an **inclusive**
// interval between `starti` and `endi`.
//
// Return `true` _if each integer in the inclusive range_`[left, right]` _is
// covered by **at least one** interval in_ `ranges`. Return `false`
// _otherwise_.
//
// An integer `x` is covered by an interval `ranges[i] = [starti, endi]` if
// `starti <= x <= endi`.
//
// **Example 1:**
//
// ```
// Input: ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
// Output: true
// Explanation: Every integer between 2 and 5 is covered:
// - 2 is covered by the first range.
// - 3 and 4 are covered by the second range.
// - 5 is covered by the third range.
//
// ```
//
// **Example 2:**
//
// ```
// Input: ranges = [[1,10],[10,20]], left = 21, right = 21
// Output: false
// Explanation: 21 is not covered by any range.
//
// ```
//
// **Constraints:**
//
// - `1 <= ranges.length <= 50`
// - `1 <= starti <= endi <= 50`
// - `1 <= left <= right <= 50`
func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		a, b := ranges[i], ranges[j]
		if a[0] == b[0] {
			return a[1] <= b[1]
		}

		return a[0] <= b[0]
	})

	for _, a := range ranges {
		if left > right {
			return true
		}

		if a[0] <= left && a[1] >= left {
			left = a[1]+1
		}
	}

	return left > right
}
