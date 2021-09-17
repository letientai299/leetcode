package main

// Minimum Operations to Reduce X to Zero
//
// Medium
//
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
//
// You are given an integer array `nums` and an integer `x`. In one operation,
// you can either remove the leftmost or the rightmost element from the array
// `nums` and subtract its value from `x`. Note that this **modifies** the array
// for future operations.
//
// Return _the **minimum number** of operations to reduce_ `x` _to **exactly**_
// `0` _if it is possible_ _, otherwise, return_`-1`.
//
// **Example 1:**
//
// ```
// Input: nums = [1,1,4,2,3], x = 5
// Output: 2
// Explanation: The optimal solution is to remove the last two elements to
// reduce x to zero.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [5,6,7,8,9], x = 4
// Output: -1
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [3,2,20,1,1,3], x = 10
// Output: 5
// Explanation: The optimal solution is to remove the last three elements and
// the first two elements (5 operations in total) to reduce x to zero.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `1 <= nums[i] <= 104`
// - `1 <= x <= 109`
func minOperations(nums []int, x int) int {
	all := 0
	for _, v := range nums {
		all += v
	}

	if all < x {
		return -1
	}

	if all == x {
		return len(nums)
	}

	target := all - x
	sum := 0

	// now we find len of the longest sub array such that nums[i]-nums[j] == temp
	l, r := 0, 0
	best := -1
	for ; r < len(nums); r++ {
		sum += nums[r]

		for ; sum > target && l < r; l++ {
			sum -= nums[l]
		}

		if sum == target && best < r-l+1 {
			best = r - l + 1
		}
	}

	if best == -1 {
		return best
	}

	return len(nums) - best
}
