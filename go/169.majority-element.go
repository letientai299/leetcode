package main

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (51.77%)
 * Total Accepted:    357.6K
 * Total Submissions: 690.6K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
func majorityElement169(nums []int) int {
	major := 0
	c := 0

	// this only works because the input appears more than half
	for _, x := range nums {
		if c == 0 {
			major = x
			c++
			continue
		}

		if x != major {
			c--
		} else {
			c++
		}
	}
	return major
}
