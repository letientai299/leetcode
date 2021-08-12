package main

import (
	"sort"
)

// Partition to K Equal Sum Subsets
//
// Medium
//
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
//
// Given an integer array `nums` and an integer `k`, return `true` if it is
// possible to divide this array into `k` non-empty subsets whose sums are all
// equal.
//
// **Example 1:**
//
// ```
// Input: nums = [4,3,2,3,5,2,1], k = 4
// Output: true
// Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3),
// (2,3) with equal sums.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,3,4], k = 3
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= k <= nums.length <= 16`
// - `1 <= nums[i] <= 104`
// - The frequency of each element is in the range `[1, 4]`.
func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	total := 0
	for _, v := range nums {
		total += v
	}

	if total%k != 0 || nums[len(nums)-1] > total/k {
		return false
	}

	want := total / k
	used := make([]bool, len(nums))
	var check func(i, cur, part int) bool

	check = func(i, cur, part int) bool {
		if part == 0 {
			return true
		}

		if cur == want {
			return check(0, 0, part-1)
		}

		for ; i < len(nums) && used[i]; i++ {
		}

		if i >= len(nums) || nums[i] > want-cur {
			return false
		}

		used[i] = true
		if check(i+1, cur+nums[i], part) {
			return true
		}

		used[i] = false
		return check(i+1, cur, part)
	}

	return check(0, 0, k)
}
