package main

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (52.98%)
 * Total Accepted:    143.2K
 * Total Submissions: 270.3K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
 * elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the
 * returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */
func findDisappearedNumbers(nums []int) []int {
	var res []int
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		old := nums[x-1]
		if old > 0 {
			nums[x-1] = -old
		}
	}

	for i, x := range nums {
		if x > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
