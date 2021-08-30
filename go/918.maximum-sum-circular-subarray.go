package main

// Maximum Sum Circular Subarray
//
// Medium
//
// https://leetcode.com/problems/maximum-sum-circular-subarray/
//
// Given a **circular integer array** `nums` of length `n`, return _the maximum
// possible sum of a non-empty **subarray** of_ `nums`.
//
// A **circular array** means the end of the array connects to the beginning of
// the array. Formally, the next element of `nums[i]` is `nums[(i + 1) % n]` and
// the previous element of `nums[i]` is `nums[(i - 1 + n) % n]`.
//
// A **subarray** may only include each element of the fixed buffer `nums` at
// most once. Formally, for a subarray `nums[i], nums[i + 1], ..., nums[j]`,
// there does not exist `i <= k1`, `k2 <= j` with `k1 % n == k2 % n`.
//
// **Example 1:**
//
// ```
// Input: nums = [1,-2,3,-2]
// Output: 3
// Explanation: Subarray [3] has maximum sum 3
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [5,-3,5]
// Output: 10
// Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [3,-1,2,-1]
// Output: 4
// Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [3,-2,2,-3]
// Output: 3
// Explanation: Subarray [3] and [3,-2,2] both have maximum sum 3
//
// ```
//
// **Example 5:**
//
// ```
// Input: nums = [-2,-3,-1]
// Output: -1
// Explanation: Subarray [-1] has maximum sum -1
//
// ```
//
// **Constraints:**
//
// - `n == nums.length`
// - `1 <= n <= 3 * 104`
// - `-3 * 104 <= nums[i] <= 3 * 104`
func maxSubarraySumCircular(nums []int) int {
	// need to find min sub array sum, where the sub array can be empty

	curMin := nums[0]
	curMax := curMin
	all := curMin
	minSub := curMin
	maxSub := curMax

	for _, v := range nums[1:] {
		all += v

		if curMin > 0 && curMin > v {
			curMin = v
		} else {
			curMin += v
		}

		if curMax > 0 {
			curMax += v
		} else {
			curMax = v
		}

		if minSub > curMin {
			minSub = curMin
		}

		if maxSub < curMax {
			maxSub = curMax
		}
	}

	all -= minSub
	if all > maxSub && all > 0 {
		return all
	}

	return maxSub
}

// Good for interview, but quite tricky, should be 2nd level.
