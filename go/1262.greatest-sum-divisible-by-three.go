package main

import "math"

// Greatest Sum Divisible by Three
//
// Medium
//
// https://leetcode.com/problems/greatest-sum-divisible-by-three/
//
// Given an array `nums` of integers, we need to find the maximum possible sum
// of elements of the array such that it is divisible by three.
//
// **Example 1:**
//
// ```
// Input: nums = [3,6,5,1,8]
// Output: 18
// Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum
// divisible by 3).
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [4]
// Output: 0
// Explanation: Since 4 is not divisible by 3, do not pick any number.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,2,3,4,4]
// Output: 12
// Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum
// divisible by 3).
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 4 * 10^4`
// - `1 <= nums[i] <= 10^4`
func maxSumDivThree(nums []int) int {
	r := 0
	mins := make([][]int, 2)
	for _, v := range nums {
		r += v
		id := v % 3
		if id == 0 {
			continue
		}

		id--
		if len(mins[id]) < 2 {
			mins[id] = append(mins[id], v)
			continue
		}

		a, b := mins[id][0], mins[id][1]
		if v >= a && v >= b {
			continue
		}

		if a >= b && a >= v {
			mins[id][0] = v
		} else {
			mins[id][1] = v
		}
	}

	remain := r % 3
	if remain == 0 {
		return r
	}

	remain--
	x := math.MaxInt32
	for _, v := range mins[remain] {
		if v < x {
			x = v
		}
	}

	if len(mins[1-remain]) != 2 {
		return r - x
	}

	y := mins[1-remain][0] + mins[1-remain][1]
	if x < y {
		return r - x
	}

	return r - y
}
