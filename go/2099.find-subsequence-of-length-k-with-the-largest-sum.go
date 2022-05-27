package main

import (
	"sort"
)

// Find Subsequence of Length K With the Largest Sum
//
// Easy
//
// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/
//
// You are given an integer array `nums` and an integer `k`. You want to find a
// **subsequence** of `nums` of length `k` that has the **largest** sum.
//
// Return _**any** such subsequence as an integer array of length_ `k`.
//
// A **subsequence** is an array that can be derived from another array by
// deleting some or no elements without changing the order of the remaining
// elements.
//
// **Example 1:**
//
// ```
// Input: nums = [2,1,3,3], k = 2
// Output: [3,3]
// Explanation:
// The subsequence has the largest sum of 3 + 3 = 6.
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [-1,-2,3,4], k = 3
// Output: [-1,3,4]
// Explanation:
// The subsequence has the largest sum of -1 + 3 + 4 = 6.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [3,4,3,3], k = 2
// Output: [3,4]
// Explanation:
// The subsequence has the largest sum of 3 + 4 = 7.
// Another possible subsequence is [4, 3].
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 1000`
// - `-105 <= nums[i] <= 105`
// - `1 <= k <= nums.length`
func maxSubsequence(nums []int, k int) []int {
	pair := make([][2]int, k)
	for i := range pair {
		pair[i] = [2]int{i, nums[i]}
	}

	sort.Slice(pair, func(i, j int) bool { return pair[i][1] <= pair[j][1] })

	for i := k; i < len(nums); i++ {
		if nums[i] <= pair[0][1] {
			continue // the new value is even smaller than min
		}

		x := sort.Search(k, func(j int) bool {
			return pair[j][1] > nums[i]
		})

		p := [2]int{i, nums[i]}
		if x >= k {
			pair = append(pair[1:], p)
		} else {
			copy(pair[0:x-1], pair[1:x])
			pair[x-1] = p
		}
	}

	sort.Slice(pair, func(i, j int) bool { return pair[i][0] <= pair[j][0] })
	r := make([]int, k)
	for i := range r {
		r[i] = pair[i][1]
	}

	return r
	// TODO (tai): can be faster
}
