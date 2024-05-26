package main

// Maximum Gap
//
// # Medium
//
// https://leetcode.com/problems/maximum-gap
//
// Given an integer array `nums`, return _the maximum difference between two
// successive elements in its sorted form_. If the array contains less than two
// elements, return `0`.
//
// You must write an algorithm that runs in linear time and uses linear extra
// space.
//
// **Example 1:**
//
// ```
// Input: nums = [3,6,9,1]
// Output: 3
// Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9)
// has the maximum difference 3.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [10]
// Output: 0
// Explanation: The array contains less than 2 elements, therefore return 0.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `0 <= nums[i] <= 109`
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	m := nums[0]
	for _, v := range nums[1:] {
		m = max(m, v)
	}

	tmp := make([]int, len(nums))
	for exp := 1; exp <= m; exp *= 10 {
		var count [10]int
		for _, v := range nums {
			d := (v % (exp * 10)) / exp
			count[d]++ // count occurrence of digits
		}

		for i, v := range count[:9] {
			count[i+1] += v
		}

		for i := len(nums) - 1; i >= 0; i-- {
			v := nums[i]
			d := (v % (exp * 10)) / exp
			tmp[count[d]-1] = v
			count[d]--
		}

		tmp, nums = nums, tmp
	}

	best := 0
	for i, v := range nums[1:] {
		best = max(best, v-nums[i])
	}
	return best
}
