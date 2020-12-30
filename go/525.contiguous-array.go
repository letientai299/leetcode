package main

/*
 * @lc app=leetcode id=525 lang=golang
 *
 * [525] Contiguous Array
 *
 * https://leetcode.com/problems/contiguous-array/description/
 *
 * algorithms
 * Medium (44.18%)
 * Total Accepted:    177.1K
 * Total Submissions: 409K
 * Testcase Example:  '[0,1]'
 *
 * Given a binary array, find the maximum length of a contiguous subarray with
 * equal number of 0 and 1.
 *
 *
 * Example 1:
 *
 * Input: [0,1]
 * Output: 2
 * Explanation: [0, 1] is the longest contiguous subarray with equal number of
 * 0 and 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [0,1,0]
 * Output: 2
 * Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal
 * number of 0 and 1.
 *
 *
 *
 * Note:
 * The length of the given binary array will not exceed 50,000.
 *
 */

func findMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	best := 0
	count := 0
	countMap := make(map[int]int)
	countMap[0] = -1
	for i, v := range nums {
		if v == 0 {
			count++
		} else {
			count--
		}
		prev, ok := countMap[count]

		if !ok {
			countMap[count] = i
			continue
		}

		if i-prev > best {
			best = i - prev
		}
	}

	return best
}
