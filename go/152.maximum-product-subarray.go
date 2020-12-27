package main

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (30.40%)
 * Total Accepted:    420.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 *
 * Example 1:
 *
 *
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 */
func maxProduct152(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	ps := make([]int, n)
	ns := make([]int, n)
	x := nums[0]
	best := 0
	if x > 0 {
		ps[0] = x
		best = x
	} else {
		ns[0] = x
	}

	for i := 1; i < n; i++ {
		x = nums[i]
		if x > 0 {
			if ps[i-1] == 0 {
				ps[i] = x
			} else {
				ps[i] = x * ps[i-1]
			}
			ns[i] = x * ns[i-1]
		} else if x < 0 {
			ps[i] = ns[i-1] * x
			if ps[i-1] != 0 {
				ns[i] = x * ps[i-1]
			} else {
				ns[i] = x
			}
		}

		if best < ps[i] {
			best = ps[i]
		}
	}

	return best
}
