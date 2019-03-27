package main

/*
* @lc app=leetcode id=283 lang=golang
*
* [283] Move Zeroes
*
* https://leetcode.com/problems/move-zeroes/description/
*
* algorithms
* Easy (53.84%)
* Total Accepted:    436.4K
* Total Submissions: 810.4K
* Testcase Example:  '[0,1,0,3,12]'
*
* Given an array nums, write a function to move all 0's to the end of it while
* maintaining the relative order of the non-zero elements.
*
* Example:
*
*
* Input: [0,1,0,3,12]
* Output: [1,3,12,0,0]
*
* Note:
*
*
* You must do this in-place without making a copy of the array.
* Minimize the total number of operations.
*
 */
func moveZeroes(nums []int) {
	i, j := 0, 0
	n := len(nums)
	for i < n && j < n {
		for i < n && nums[i] != 0 {
			i++
		}

		j = i + 1
		for j < n && nums[j] == 0 {
			j++
		}

		if j >= n {
			return
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j++
	}
}
