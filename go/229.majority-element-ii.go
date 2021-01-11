package main

import "math"

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 *
 * https://leetcode.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (33.60%)
 * Total Accepted:    184.2K
 * Total Submissions: 481.5K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an integer array of size n, find all elements that appear more than ⌊
 * n/3 ⌋ times.
 *
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,2,3]
 * Output: [3]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: [1]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2]
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

func majorityElement(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	marjors := [2]int{math.MaxInt32, math.MaxInt32}
	cnt := [2]int{0, 0}

	for i := 0; i < n; i++ {
		x := nums[i]
		switch {
		case x == marjors[0]:
			cnt[0]++
		case x == marjors[1]:
			cnt[1]++
		case cnt[0] == 0:
			marjors[0] = x
			cnt[0] = 1
		case cnt[1] == 0:
			marjors[1] = x
			cnt[1] = 1
		default:
			cnt[0]--
			cnt[1]--
		}
	}

	cnt[0] = 0
	cnt[1] = 0

	for _, x := range nums {
		if x == marjors[0] {
			cnt[0]++
		} else if x == marjors[1] {
			cnt[1]++
		}
	}

	var r []int
	if cnt[0] > n/3 {
		r = append(r, marjors[0])
	}
	if cnt[1] > n/3 {
		r = append(r, marjors[1])
	}

	return r
}
