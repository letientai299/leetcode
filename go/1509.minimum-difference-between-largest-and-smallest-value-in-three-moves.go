package main

import (
	"sort"
)

// Minimum Difference Between Largest and Smallest Value in Three Moves
//
// Medium
//
// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
//
// Given an array `nums`, you are allowed to choose one element of `nums` and
// change it by any value in one move.
//
// Return the minimum difference between the largest and smallest value of
// `nums` after perfoming at most 3 moves.
//
// **Example 1:**
//
// ```
// Input: nums = [5,3,2,4]
// Output: 0
// Explanation: Change the array [5,3,2,4] to [2,2,2,2].
// The difference between the maximum and minimum is 2-2 = 0.
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,5,0,10,14]
// Output: 1
// Explanation: Change the array [1,5,0,10,14] to [1,1,0,1,1].
// The difference between the maximum and minimum is 1-0 = 1.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [6,6,0,1,1,4,6]
// Output: 2
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [1,5,6,14,15]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 10^5`
// - `-10^9 <= nums[i] <= 10^9`
func minDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	max4 := append([]int{}, nums[:4]...)
	sort.Ints(max4)
	min4 := append([]int{}, nums[:4]...)
	sort.Ints(min4)

	for _, v := range nums[4:] {
		i := 0
		for i < 4 && v >= max4[i] {
			i++
		}
		if i > 0 {
			copy(max4[:i-1], max4[1:i])
			max4[i-1] = v
		}

		i = 3
		for i >= 0 && v <= min4[i] {
			i--
		}
		if i < 3 {
			copy(min4[i+2:], min4[i+1:])
			min4[i+1] = v
		}
	}

	return min(
		max4[3]-min4[0],
		max4[2]-min4[0],
		max4[1]-min4[0],
		max4[0]-min4[0],

		max4[3]-min4[1],
		max4[2]-min4[1],
		max4[1]-min4[1],

		max4[3]-min4[2],
		max4[2]-min4[2],

		max4[3]-min4[3],
	)
}
