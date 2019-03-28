package main

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (53.49%)
 * Total Accepted:    204.4K
 * Total Submissions: 382K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 *
 *
 * Note:
 *
 *
 * Each element in the result must be unique.
 * The result can be in any order.
 *
 *
 *
 *
 */
func intersection(nums1 []int, nums2 []int) []int {
	distinct := make(map[int]bool, len(nums1))
	for _, n := range nums1 {
		distinct[n] = true
	}

	res := make([]int, 0, len(distinct))
	for _, n := range nums2 {
		if distinct[n] {
			res = append(res, n)
			delete(distinct, n)
		}
	}
	return res
}
