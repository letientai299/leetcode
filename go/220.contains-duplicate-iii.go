package main

/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 *
 * https://leetcode.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (20.29%)
 * Total Accepted:    161K
 * Total Submissions: 755.5K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * Given an array of integers, find out whether there are two distinct indices
 * i and j in the array such that the absolute difference between nums[i] and
 * nums[j] is at most t and the absolute difference between i and j is at most
 * k.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,1], k = 3, t = 0
 * Output: true
 * Example 2:
 * Input: nums = [1,0,1,1], k = 1, t = 2
 * Output: true
 * Example 3:
 * Input: nums = [1,5,9,1,5,9], k = 2, t = 3
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 2 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 * 0 <= k <= 10^4
 * 0 <= t <= 2^31 - 1
 *
 *
 */

// TODO (tai): slow 40%, 288ms, can be 8ms
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 1; j <= k && i+j < n; j++ {
			//fmt.Println(nums[i], nums[i+j], abs(nums[i]-nums[i+j]))
			if abs(nums[i]-nums[i+j]) <= t {
				return true
			}
		}
	}
	return false
}
