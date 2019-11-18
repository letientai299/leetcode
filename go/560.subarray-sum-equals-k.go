package main

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 *
 * https://leetcode.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (43.02%)
 * Total Accepted:    161.2K
 * Total Submissions: 373.8K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * Given an array of integers and an integer k, you need to find the total
 * number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 *
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 *
 *
 *
 * Note:
 *
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the
 * integer k is [-1e7, 1e7].
 *
 *
 *
 */
func subarraySum(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	s := 0
	sums := make(map[int]int, n)
	count := 0
	sums[0] = 1
	for i := 0; i < n; i++ {
		s += nums[i]
		v := s - k
		if _, ok := sums[v]; ok {
			count += sums[v]
		}

		sums[s] += 1
	}

	return count
}
