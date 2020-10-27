package main

import "strconv"

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 *
 * https://leetcode.com/problems/summary-ranges/description/
 *
 * algorithms
 * Medium (37.41%)
 * Total Accepted:    176.4K
 * Total Submissions: 440.2K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * You are given a sorted unique integer array nums.
 *
 * Return the smallest sorted list of ranges that cover all the numbers in the
 * array exactly. That is, each element of nums is covered by exactly one of
 * the ranges, and there is no integer x such that x is in one of the ranges
 * but not in nums.
 *
 * Each range [a,b] in the list should be output as:
 *
 *
 * "a->b" if a != b
 * "a" if a == b
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: The ranges are:
 * [0,2] --> "0->2"
 * [4,5] --> "4->5"
 * [7,7] --> "7"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: The ranges are:
 * [0,0] --> "0"
 * [2,4] --> "2->4"
 * [6,6] --> "6"
 * [8,9] --> "8->9"
 *
 *
 * Example 3:
 *
 *
 * Input: nums = []
 * Output: []
 *
 *
 * Example 4:
 *
 *
 * Input: nums = [-1]
 * Output: ["-1"]
 *
 *
 * Example 5:
 *
 *
 * Input: nums = [0]
 * Output: ["0"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 20
 * -2^31 <= nums[i] <= 2^31 - 1
 * All the values of nums are unique.
 *
 *
 */
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	fmtRange := func(a ...int) string {
		if len(a) == 1 || a[0] == a[1] {
			return strconv.Itoa(a[0])
		}
		return strconv.Itoa(a[0]) + "->" + strconv.Itoa(a[1])
	}

	var res []string
	a := nums[0]
	b := nums[0]

	for i := 1; i < len(nums); i++ {
		c := nums[i]
		if b == c-1 {
			b = c
			continue
		}

		res = append(res, fmtRange(a, b))
		a = c
		b = c
	}
	res = append(res, fmtRange(a, b))
	return res
}
