package main

// Average Value of Even Numbers That Are Divisible by Three
//
// # Easy
//
// https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three
//
// Given an integer array `nums` of **positive** integers, return _the average
// value of all even integers that are divisible by_ `3` _._
//
// Note that the **average** of `n` elements is the **sum** of the `n` elements
// divided by `n` and **rounded down** to the nearest integer.
//
// **Example 1:**
//
// ```
// Input: nums = [1,3,6,10,12,15]
// Output: 9
// Explanation: 6 and 12 are even numbers that are divisible by 3. (6 + 12) / 2
// = 9.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,4,7,10]
// Output: 0
// Explanation: There is no single number that satisfies the requirement, so
// return 0.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 1000`
// - `1 <= nums[i] <= 1000`
func averageValue(nums []int) int {
	agg := 0
	cnt := 0
	for _, v := range nums {
		if v%6 == 0 {
			cnt++
			agg += v
		}
	}

	if cnt == 0 {
		return 0
	}

	return agg / cnt
}
