package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.48%)
 * Total Accepted:    366.5K
 * Total Submissions: 905.4K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	if right == 0 {
		return 0
	}

	for left < right-1 {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}

	if target <= nums[left] {
		return left
	}

	return left + 1
}
