package main

/*
 * @lc app=leetcode id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 *
 * https://leetcode.com/problems/subarray-product-less-than-k/description/
 *
 * algorithms
 * Medium (38.12%)
 * Total Accepted:    89.8K
 * Total Submissions: 222.7K
 * Testcase Example:  '[10,5,2,6]\n100'
 *
 * Your are given an array of positive integers nums.
 * Count and print the number of (contiguous) subarrays where the product of
 * all the elements in the subarray is less than k.
 *
 * Example 1:
 *
 * Input: nums = [10, 5, 2, 6], k = 100
 * Output: 8
 * Explanation: The 8 subarrays that have product less than 100 are: [10], [5],
 * [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
 * Note that [10, 5, 2] is not included as the product of 100 is not strictly
 * less than k.
 *
 *
 *
 * Note:
 * 0 < nums.length .
 * 0 < nums[i] < 1000.
 * 0 .
 *
 */
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	r := 0
	p := 1
	i, j := 0, -1
	last := -1

	for {
		for p >= k && i < n {
			p /= nums[i]
			i++
		}

		for j++; j < n && p*nums[j] < k; j++ {
			p *= nums[j]
		}

		r += (j - i) * (j - i + 1) / 2
		if last != -1 {
			r -= (last - i) * (last - i + 1) / 2
		}
		last = j

		if j == n {
			break
		}

		p *= nums[j]
	}

	return r
}
