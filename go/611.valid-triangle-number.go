package main

import "sort"

/*
 * @lc app=leetcode id=611 lang=golang
 *
 * [611] Valid Triangle Number
 *
 * https://leetcode.com/problems/valid-triangle-number/description/
 *
 * algorithms
 * Medium (46.81%)
 * Total Accepted:    74.3K
 * Total Submissions: 151.5K
 * Testcase Example:  '[2,2,3,4]'
 *
 * Given an array consists of non-negative integers,  your task is to count the
 * number of triplets chosen from the array that can make triangles if we take
 * them as side lengths of a triangle.
 *
 * Example 1:
 *
 * Input: [2,2,3,4]
 * Output: 3
 * Explanation:
 * Valid combinations are:
 * 2,3,4 (using the first 2)
 * 2,3,4 (using the second 2)
 * 2,2,3
 *
 *
 *
 * Note:
 *
 * The length of the given array won't exceed 1000.
 * The integers in the given array are in the range of [0, 1000].
 *
 *
 *
 */

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for k := n - 1; k > 1; k-- {
		c := nums[k]
		for i, j := 0, k-1; i < j; {
			if nums[i]+nums[j] <= c {
				i++
			} else {
				res += j - i
				j--
			}
		}
	}
	return res
}
