package main

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 *
 * https://leetcode.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (39.72%)
 * Total Accepted:    183.4K
 * Total Submissions: 453K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given an integer array nums, return true if there exists a triple of indices
 * (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such
 * indices exists, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,4,5]
 * Output: true
 * Explanation: Any triplet where i < j < k is valid.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [5,4,3,2,1]
 * Output: false
 * Explanation: No triplet exists.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [2,1,5,0,4,6]
 * Output: true
 * Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4]
 * == 4 < nums[5] == 6.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you implement a solution that runs in O(n) time complexity
 * and O(1) space complexity?
 */
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	higher := math.MaxInt32
	lower := math.MaxInt32
	for i := 0; i < n-1; i++ {
		if nums[i] > higher {
			return true
		}

		if nums[i] >= nums[i+1] {
			continue
		}

		if nums[i] > lower || nums[i+1] > higher {
			return true
		}

		lower = nums[i]
		higher = nums[i+1]
	}

	return false
}
