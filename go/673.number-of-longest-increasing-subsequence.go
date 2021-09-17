package main

// Number of Longest Increasing Subsequence
//
// Medium
//
// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
//
// Given an integer array `nums`, return _the number of longest increasing
// subsequences._
//
// **Notice** that the sequence has to be **strictly** increasing.
//
// **Example 1:**
//
// ```
// Input: nums = [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1,
// 3, 5, 7].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [2,2,2,2,2]
// Output: 5
// Explanation: The length of longest continuous increasing subsequence is 1,
// and there are 5 subsequences' length is 1, so output 5.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 2000`
// - `-106 <= nums[i] <= 106`
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	counts := make([]int, n)

	for i, v := range nums {
		dp[i] = 1
		counts[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= v {
				continue
			}

			if dp[j] == dp[i]-1 {
				counts[i] += counts[j]
			} else if dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
				counts[i] = counts[j]
			}
		}
	}

	r := 0
	best := 1
	for i, v := range dp {
		if best < v {
			best = v
			r = counts[i]
		} else if best == v {
			r += counts[i]
		}
	}

	return r
}

// TODO (tai): try to understand this O(nlogn) solution:
//  https://leetcode.com/problems/number-of-longest-increasing-subsequence/discuss/107295/9ms-C%2B%2B-Explanation%3A-DP-%2B-Binary-search-%2B-prefix-sums-O(NlogN)-time-O(N)-space
