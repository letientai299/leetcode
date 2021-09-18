package main

// Single Number II
//
// Medium
//
// https://leetcode.com/problems/single-number-ii/
//
// Given an integer array `nums` where every element appears **three times**
// except for one, which appears **exactly once**. _Find the single element and
// return it_.
//
// You must implement a solution with a linear runtime complexity and use only
// constant extra space.
//
// **Example 1:**
//
// ```
// Input: nums = [2,2,3,2]
// Output: 3
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,1,0,1,0,1,99]
// Output: 99
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 3 * 104`
// - `-231 <= nums[i] <= 231 - 1`
// - Each element in `nums` appears exactly **three times** except for one
// element which appears **once**.
func singleNumber(nums []int) int {
	once, twice := 0, 0
	for _, v := range nums {
		once = (once ^ v) & (^twice)
		twice = (twice ^ v) & (^once)
	}

	return once
}

// See: https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers
