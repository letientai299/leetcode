package main

import (
	"math"
)

// Sum of Subarray Minimums
//
// Medium
//
// https://leetcode.com/problems/sum-of-subarray-minimums/
//
// Given an array of integers arr, find the sum of `min(b)`, where `b` ranges
// over every (contiguous) subarray of `arr`. Since the answer may be large,
// return the answer **modulo** `109 + 7`.
//
// **Example 1:**
//
// ```
// Input: arr = [3,1,2,4]
// Output: 17
// Explanation:
// Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4],
// [3,1,2,4].
// Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
// Sum is 17.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [11,81,94,43,3]
// Output: 444
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 3 * 104`
// - `1 <= arr[i] <= 3 * 104`
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	all := 0

	// 3 1 2 4
	//   1      6
	// 3        3
	//     2    4
	//       4  4

	l, r := 0, 0 // find range [l, r] where arr[i] is the min
	pre := math.MinInt32

	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if v <= pre {
			l++
			for i-1-l >= 0 && arr[i-1-l] >= v {
				l++
			}
		} else {
			l = 0
			r = 0
		}

		for i+1+r < n && arr[i+1+r] > v {
			r++
		}

		all = (all + (l+1)*(r+1)*v) % (1e9 + 7)
		pre = v
	}

	return all
}

// TODO (tai): can be O(n) with stack
