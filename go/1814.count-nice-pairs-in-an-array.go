package main

// Count Nice Pairs in an Array
//
// Medium
//
// https://leetcode.com/problems/count-nice-pairs-in-an-array/
//
// You are given an array `nums` that consists of non-negative integers. Let us
// define `rev(x)` as the reverse of the non-negative integer `x`. For example,
// `rev(123) = 321`, and `rev(120) = 21`. A pair of indices `(i, j)` is **nice**
// if it satisfies all of the following conditions:
//
// - `0 <= i < j < nums.length`
// - `nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])`
//
// Return _the number of nice pairs of indices_. Since that number can be too
// large, return it **modulo** `109 + 7`.
//
// **Example 1:**
//
// ```
// Input: nums = [42,11,1,97]
// Output: 2
// Explanation: The two pairs are:
//  - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
//  - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [13,10,35,24,76]
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `0 <= nums[i] <= 109`
func countNicePairs(nums []int) int {
	rev := func(n int) int {
		r := 0
		for n > 0 {
			r = 10*r + n%10
			n /= 10
		}
		return r
	}

	m := make(map[int]uint64, len(nums))
	res := uint64(0)
	for _, v := range nums {
		r := rev(v)
		res += m[v-r]
		m[v-r]++
	}

	return int(res % (1e9 + 7))
}
