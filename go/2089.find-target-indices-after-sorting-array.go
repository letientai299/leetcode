package main

// Find Target Indices After Sorting Array
//
// Easy
//
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
//
// You are given a **0-indexed** integer array `nums` and a target element
// `target`.
//
// A **target index** is an index `i` such that `nums[i] == target`.
//
// Return _a list of the target indices of_ `nums` after _sorting_ `nums` _in
// **non-decreasing** order_. If there are no target indices, return _an
// **empty** list_. The returned list must be sorted in **increasing** order.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,5,2,3], target = 2
// Output: [1,2]
// Explanation: After sorting, nums is [1,2,2,3,5].
// The indices where nums[i] == 2 are 1 and 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,5,2,3], target = 3
// Output: [3]
// Explanation: After sorting, nums is [1,2,2,3,5].
// The index where nums[i] == 3 is 3.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,2,5,2,3], target = 5
// Output: [4]
// Explanation: After sorting, nums is [1,2,2,3,5].
// The index where nums[i] == 5 is 4.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 100`
// - `1 <= nums[i], target <= 100`
func targetIndices(nums []int, target int) []int {
	/*
		// trivial O(n logn)
		sort.Ints(nums)
		var r []int
		for i, n := range nums {
			if n == target {
				r = append(r, i)
			}
		}
		return r
	*/

	// note that the input range is very low, so we can use counting sort
	start := 0
	cnt := 0
	for _, n := range nums {
		if n < target {
			start++
		} else if n == target {
			cnt++
		}
	}

	r := make([]int, cnt)
	for i := range r {
		r[i] = i + start
	}

	return r
}
