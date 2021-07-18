package main

import "math"

// Remove One Element to Make the Array Strictly Increasing
//
// Easy
//
// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
//
// Given a **0-indexed** integer array `nums`, return `true` _if it can be made
// **strictly increasing** after removing **exactly one** element, or_ `false`
// _otherwise. If the array is already strictly increasing, return_ `true`.
//
// The array `nums` is **strictly increasing** if `nums[i - 1] < nums[i]` for
// each index `(1 <= i < nums.length).`
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,10,5,7]
// Output: true
// Explanation: By removing 10 at index 2 from nums, it becomes [1,2,5,7].
// [1,2,5,7] is strictly increasing, so return true.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [2,3,1,2]
// Output: false
// Explanation:
// [3,1,2] is the result of removing the element at index 0.
// [2,1,2] is the result of removing the element at index 1.
// [2,3,2] is the result of removing the element at index 2.
// [2,3,1] is the result of removing the element at index 3.
// No resulting array is strictly increasing, so return false.
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,1,1]
// Output: false
// Explanation: The result of removing any element is [1,1].
// [1,1] is not strictly increasing, so return false.
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [1,2,3]
// Output: true
// Explanation: [1,2,3] is already strictly increasing, so return true.
//
// ```
//
// **Constraints:**
//
// - `2 <= nums.length <= 1000`
// - `1 <= nums[i] <= 1000`
func canBeIncreasing(nums []int) bool {
	removed := false
	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		cur := nums[i]
		if prev < cur {
			continue
		}

		if removed {
			return false
		}

		beforePrev := math.MinInt32
		next := math.MaxInt32
		if i >= 2 {
			beforePrev = nums[i-2]
		}
		if i < len(nums)-1 {
			next = nums[i+1]
		}

		if beforePrev < cur || prev < next {
			removed = true
		} else {
			return false
		}
	}

	return true
}

// Good for interview
