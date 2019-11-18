package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (57.06%)
 * Total Accepted:    338.7K
 * Total Submissions: 589.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array nums of n integers where n > 1,  return an array output such
 * that output[i] is equal to the product of all the elements of nums except
 * nums[i].
 *
 * Example:
 *
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does
 * not count as extra space for the purpose of space complexity analysis.)
 *
 */
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	// 2 3 4
	// res: 1 4 6

	r := nums[n-1] // 12
	for i := n - 2; i > 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}

	res[0] = r
	return res
}
