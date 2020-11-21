package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 *
 * https://leetcode.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (29.25%)
 * Total Accepted:    275.1K
 * Total Submissions: 940.6K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * Given an array, rotate the array to the right by k steps, where k is
 * non-negative.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4,5,6,7] and k = 3
 * Output: [5,6,7,1,2,3,4]
 * Explanation:
 * rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * rotate 3 steps to the right: [5,6,7,1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: [-1,-100,3,99] and k = 2
 * Output: [3,99,-1,-100]
 * Explanation:
 * rotate 1 steps to the right: [99,-1,-100,3]
 * rotate 2 steps to the right: [3,99,-1,-100]
 *
 *
 * Note:
 *
 *
 * Try to come up as many solutions as you can, there are at least 3 different
 * ways to solve this problem.
 * Could you do it in-place with O(1) extra space?
 *
 */
func rotate_189(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k %= n
	reverseInts(nums, 0, n)
	reverseInts(nums, 0, k)
	reverseInts(nums, k, n)
}

func reverseInts(nums []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	for i := left; i < mid; i++ {
		ri := right + left - i - 1
		nums[i], nums[ri] = nums[ri], nums[i]
	}
}
