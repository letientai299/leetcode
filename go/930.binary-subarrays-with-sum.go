package main

// Binary Subarrays With Sum
//
// Medium
//
// https://leetcode.com/problems/binary-subarrays-with-sum/
//
// Given a binary array `nums` and an integer `goal`, return _the number of
// non-empty **subarrays** with a sum_ `goal`.
//
// A **subarray** is a contiguous part of the array.
//
// **Example 1:**
//
// ```
// Input: nums = [1,0,1,0,1], goal = 2
// Output: 4
// Explanation: The 4 subarrays are bolded and underlined below:
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,0,0,0,0], goal = 0
// Output: 15
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 3 * 104`
// - `nums[i]` is either `0` or `1`.
// - `0 <= goal <= nums.length`
func numSubarraysWithSum(nums []int, goal int) int {
	var ones []int
	for i, v := range nums {
		if v == 1 {
			ones = append(ones, i)
		}
	}

	if len(ones) < goal {
		return 0
	}

	n := len(nums)
	r := 0
	pre := 0

	if goal == 0 {
		for _, v := range ones {
			r += (v - pre) * (v - pre + 1) / 2
			pre = v + 1
		}

		r += (n - pre) * (n - pre + 1) / 2
		return r
	}

	ones = append(ones, n)
	for i := 0; i+goal < len(ones); i++ {
		a := ones[i] - pre
		b := ones[i+goal] - ones[i+goal-1] - 1
		r += (a + 1) * (b + 1)
		pre = ones[i] + 1
	}

	return r
}
