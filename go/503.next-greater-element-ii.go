package main

import (
	"math"
)

/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 *
 * https://leetcode.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (52.92%)
 * Total Accepted:    118.8K
 * Total Submissions: 204.8K
 * Testcase Example:  '[1,2,1]'
 *
 *
 * Given a circular array (the next element of the last element is the first
 * element of the array), print the Next Greater Number for every element. The
 * Next Greater Number of a number x is the first greater number to its
 * traversing-order next in the array, which means you could search circularly
 * to find its next greater number. If it doesn't exist, output -1 for this
 * number.
 *
 *
 * Example 1:
 *
 * Input: [1,2,1]
 * Output: [2,-1,2]
 * Explanation: The first 1's next greater number is 2; The number 2 can't find
 * next greater number; The second 1's next greater number needs to search
 * circularly, which is also 2.
 *
 *
 *
 * Note:
 * The length of given array won't exceed 10000.
 *
 */
// TODO (tai): can be faster, 87%, use stack
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	res := make([]int, n)
	max := math.MinInt32
	maxIndex := -1
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}

	for i, v := range nums {
		if v == max {
			res[i] = -1
		}
	}

	if res[n-1] != -1 {
		res[n-1] = maxIndex
	}

	for i := n - 2; i >= 0; i-- {
		v := nums[i]
		if v == max {
			continue
		}

		j := i + 1
		for v >= nums[j] {
			j = res[j]
		}
		res[i] = j
	}

	j := 0
	for j != -1 && nums[n-1] >= nums[j] {
		j = res[j]
	}
	res[n-1] = j

	for i := n - 2; i >= 0; i-- {
		v := nums[i]
		if v == max {
			continue
		}

		j := i + 1
		for v >= nums[j] {
			j = res[j]
		}
		res[i] = j
	}

	for i, v := range res {
		if v != -1 {
			res[i] = nums[v]
		}
	}

	return res
}
