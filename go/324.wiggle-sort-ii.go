package main

/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 *
 * https://leetcode.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (28.86%)
 * Total Accepted:    91.9K
 * Total Submissions: 301.6K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * Given an unsorted array nums, reorder it such that nums[0] < nums[1] >
 * nums[2] < nums[3]....
 *
 * Example 1:
 *
 *
 * Input: nums = [1, 5, 1, 1, 6, 4]
 * Output: One possible answer is [1, 4, 1, 5, 1, 6].
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 3, 2, 2, 3, 1]
 * Output: One possible answer is [2, 3, 1, 3, 1, 2].
 *
 * Note:
 * You may assume all input has valid answer.
 *
 * Follow Up:
 * Can you do it in O(n) time and/or in-place with O(1) extra space?
 *
 */
// TODO (tai): not done yet
func wiggleSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	up := true
	for i := 0; i < n-1; i++ {
		j := i + 1
		for ; j < n; j++ {
			if nums[i] == nums[j] {
				continue
			}

			if up == (nums[i] < nums[j]) {
				nums[i+1], nums[j] = nums[j], nums[i+1]
				break
			}
		}

		if j == n { // can't swap
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}

		up = !up
	}

	for i := 1; i < n-2; i++ {
		if nums[i] == nums[i+1] {
			nums[i-1], nums[i], nums[i+1], nums[i+2] = nums[i+1], nums[i+2], nums[i-1], nums[i]
			i += 2
		}
	}
}
