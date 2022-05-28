package main

import "math"

// Count Elements With Strictly Smaller and Greater Elements
//
// Easy
//
// https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
//
// Given an integer array `nums`, return _the number of elements that have
// **both** a strictly smaller and a strictly greater element appear in_ `nums`.
//
// **Example 1:**
//
// ```
// nums
// ```
//
// **Example 2:**
//
// ```
// nums
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 100`
// - `-105 <= nums[i] <= 105`
func countElements(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	min := math.MaxInt32
	minCnt := 0
	max := math.MinInt32
	maxCnt := 0
	for _, c := range nums {
		if c < min {
			min = c
			minCnt = 1
		} else if c == min {
			minCnt++
		}

		if c > max {
			max = c
			maxCnt = 1
		} else if c == max {
			maxCnt++
		}
	}

	v := len(nums) - minCnt - maxCnt
	if v < 0 {
		return 0
	}

	return v
}
