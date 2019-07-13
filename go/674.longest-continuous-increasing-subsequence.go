package main

/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (44.50%)
 * Total Accepted:    70.2K
 * Total Submissions: 157.8K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 *
 * Given an unsorted array of integers, find the length of longest continuous
 * increasing subsequence (subarray).
 *
 *
 * Example 1:
 *
 * Input: [1,3,5,4,7]
 * Output: 3
 * Explanation: The longest continuous increasing subsequence is [1,3,5], its
 * length is 3.
 * Even though [1,3,5,7] is also an increasing subsequence, it's not a
 * continuous one where 5 and 7 are separated by 4.
 *
 *
 *
 * Example 2:
 *
 * Input: [2,2,2,2,2]
 * Output: 1
 * Explanation: The longest continuous increasing subsequence is [2], its
 * length is 1.
 *
 *
 *
 * Note:
 * Length of the array will not exceed 10,000.
 *
 */
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	max := 1
	cur := 1
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			cur++
			continue
		}

		if max < cur {
			max = cur
		}

		cur = 1
	}

	if max < cur {
		return cur
	}

	return max
}
