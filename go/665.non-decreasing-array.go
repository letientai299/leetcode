package main

/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 *
 * https://leetcode.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (19.46%)
 * Total Accepted:    55.5K
 * Total Submissions: 285.1K
 * Testcase Example:  '[4,2,3]'
 *
 *
 * Given an array with n integers, your task is to check if it could become
 * non-decreasing by modifying at most 1 element.
 *
 *
 *
 * We define an array is non-decreasing if array[i]  holds for every i (1
 *
 * Example 1:
 *
 * Input: [4,2,3]
 * Output: True
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing
 * array.
 *
 *
 *
 * Example 2:
 *
 * Input: [4,2,1]
 * Output: False
 * Explanation: You can't get a non-decreasing array by modify at most one
 * element.
 *
 *
 *
 * Note:
 * The n belongs to [1, 10,000].
 *
 */
func checkPossibility(nums []int) bool {
	// location of the first number that larger than its following value
	peakId := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if peakId >= 0 {
				return false
			}
			peakId = i - 1
		}
	}

	// if the peak is the first element, we can set it lower as we need
	if peakId <= 0 {
		return true
	}

	// if the peak is the second to last, then we can modify the last value
	// to any higher value that we want
	if peakId == len(nums)-2 {
		return true
	}

	return nums[peakId-1] <= nums[peakId+1] || nums[peakId] <= nums[peakId+2]
}
