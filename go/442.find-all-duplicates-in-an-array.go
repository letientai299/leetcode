package main

/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 *
 * https://leetcode.com/problems/find-all-duplicates-in-an-array/description/
 *
 * algorithms
 * Medium (63.00%)
 * Total Accepted:    246.2K
 * Total Submissions: 358.5K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements
 * appear twice and others appear once.
 *
 * Find all the elements that appear twice in this array.
 *
 * Could you do it without extra space and in O(n) runtime?
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [2,3]
 *
 */
func findDuplicates(nums []int) []int {
	var res []int
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] < 0 {
			res = append(res, x)
		}
		nums[x-1] *= -1
	}

	return res
}
