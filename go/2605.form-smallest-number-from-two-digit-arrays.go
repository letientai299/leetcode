package main

import (
	"sort"
)

// Form Smallest Number From Two Digit Arrays
//
// # Easy
//
// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays
//
// Given two arrays of **unique** digits `nums1` and `nums2`, return _the
// **smallest** number that contains **at least** one digit from each array_.
//
// **Example 1:**
//
// ```
// Input: nums1 = [4,1,3], nums2 = [5,7]
// Output: 15
// Explanation: The number 15 contains the digit 1 from nums1 and the digit 5
// from nums2. It can be proven that 15 is the smallest number we can have.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums1 = [3,5,2,6], nums2 = [3,1,7]
// Output: 3
// Explanation: The number 3 contains the digit 3 which exists in both arrays.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums1.length, nums2.length <= 9`
// - `1 <= nums1[i], nums2[i] <= 9`
// - All digits in each array are **unique**.
func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var common = 100000000
	for _, d := range nums1 {
		i := sort.SearchInts(nums2, d)
		if i < len(nums2) && nums2[i] == d {
			common = d
			break
		}
	}

	if common < 10 {
		return common
	}

	x, y := nums1[0], nums2[0]
	return min(x*10+y, y*10+x, common)
}
