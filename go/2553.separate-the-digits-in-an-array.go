package main

// Separate the Digits in an Array
//
// # Easy
//
// https://leetcode.com/problems/separate-the-digits-in-an-array
//
// Given an array of positive integers `nums`, return _an array_ `answer` _that
// consists of the digits of each integer in_ `nums` _after separating them in
// **the same order** they appear in_ `nums`.
//
// To separate the digits of an integer is to get all the digits it has in the
// same order.
//
// - For example, for the integer `10921`, the separation of its digits is
// `[1,0,9,2,1]`.
//
// **Example 1:**
//
// ```
// Input: nums = [13,25,83,77]
// Output: [1,3,2,5,8,3,7,7]
// Explanation:
// - The separation of 13 is [1,3].
// - The separation of 25 is [2,5].
// - The separation of 83 is [8,3].
// - The separation of 77 is [7,7].
// answer = [1,3,2,5,8,3,7,7]. Note that answer contains the separations in the
// same order.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [7,1,3,9]
// Output: [7,1,3,9]
// Explanation: The separation of each integer in nums is itself.
// answer = [7,1,3,9].
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 1000`
// - `1 <= nums[i] <= 105`
func separateDigits(nums []int) []int {
	res := make([]int, 0, len(nums)*4)
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		for n > 0 {
			d := n % 10
			n /= 10
			res = append(res, d)
		}
	}

	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}

	return res
}
