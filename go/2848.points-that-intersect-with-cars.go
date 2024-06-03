package main

import (
	"sort"
)

// Points That Intersect With Cars
//
// # Easy
//
// https://leetcode.com/problems/points-that-intersect-with-cars
//
// You are given a **0-indexed** 2D integer array `nums` representing the
// coordinates of the cars parking on a number line. For any index `i`, `nums[i]
// = [starti, endi]` where `starti` is the starting point of the `ith` car and
// `endi` is the ending point of the `ith` car.
//
// Return _the number of integer points on the line that are covered with **any
// part** of a car._
//
// **Example 1:**
//
// ```
// Input: nums = [[3,6],[1,5],[4,7]]
// Output: 7
// Explanation: All the points from 1 to 7 intersect at least one car, therefore
// the answer would be 7.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [[1,3],[5,8]]
// Output: 7
// Explanation: Points intersecting at least one car are 1, 2, 3, 5, 6, 7, 8.
// There are a total of 7 points, therefore the answer would be 7.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 100`
// - `nums[i].length == 2`
// - `1 <= starti <= endi <= 100`
// O(n log(n))
func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return a[0] < b[0]
	})

	start, end := nums[0][0], nums[0][1]
	cnt := 0
	for _, b := range nums[1:] {
		if end < b[0] {
			cnt += end - start + 1
			start, end = b[0], b[1]
			continue
		}

		if end < b[1] {
			end = b[1]
		}
	}

	return cnt + end - start + 1
}
