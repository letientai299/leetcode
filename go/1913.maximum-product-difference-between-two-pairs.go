package main

import (
	"math"
)

// Maximum Product Difference Between Two Pairs
//
// Easy
//
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
//
// The **product difference** between two pairs `(a, b)` and `(c, d)` is defined
// as `(a * b) - (c * d)`.
//
// - For example, the product difference between `(5, 6)` and `(2, 7)` is `(5 *
// 6) - (2 * 7) = 16`.
//
// Given an integer array `nums`, choose four **distinct** indices `w`, `x`,
// `y`, and `z` such that the **product difference** between pairs `(nums[w],
// nums[x])` and `(nums[y], nums[z])` is **maximized**.
//
// Return _the **maximum** such product difference_.
//
// **Example 1:**
//
// ```
// Input: nums = [5,6,2,7,4]
// Output: 34
// Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and
// indices 2 and 4 for the second pair (2, 4).
// The product difference is (6 * 7) - (2 * 4) = 34.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [4,2,5,9,7,4,8]
// Output: 64
// Explanation: We can choose indices 3 and 6 for the first pair (9, 8) and
// indices 1 and 5 for the second pair (2, 4).
// The product difference is (9 * 8) - (2 * 4) = 64.
//
// ```
//
// **Constraints:**
//
// - `4 <= nums.length <= 104`
// - `1 <= nums[i] <= 104`
func maxProductDifference(nums []int) int {
	// sort is not the fastest way. We need top 2 and bottom 2, quick select
	// should works, but I'm lazy, so some for loops.

	t1, b1 := 0, math.MaxInt32
	countT1, countB1 := 0, 0
	for _, v := range nums {
		if v > t1 {
			t1 = v
			countT1 = 1
		} else if v == t1 {
			countT1++
		}

		if v < b1 {
			b1 = v
			countB1 = 1
		} else if v == b1 {
			countB1++
		}
	}

	t2, b2 := 0, math.MaxInt32
	for _, v := range nums {
		if v > t2 && v != t1 {
			t2 = v
		}
		if v < b2 && v != b1 {
			b2 = v
		}
	}

	if countT1 > 1 {
		t2 = t1
	}
	if countB1 > 1 {
		b2 = b1
	}

	return t2*t1 - b1*b2
}
