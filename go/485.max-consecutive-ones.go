package main

/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 *
 * https://leetcode.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (54.77%)
 * Total Accepted:    130.4K
 * Total Submissions: 238.1K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * Given a binary array, find the maximum number of consecutive 1s in this
 * array.
 *
 * Example 1:
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive
 * 1s.
 * ⁠   The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */
func findMaxConsecutiveOnes485(nums []int) int {
	max := 0
	cur := 0
	for _, i := range nums {
		if i == 1 {
			cur++
			continue
		}

		if max < cur {
			max = cur
		}
		cur = 0
	}
	if max < cur {
		return cur
	}
	return max
}
