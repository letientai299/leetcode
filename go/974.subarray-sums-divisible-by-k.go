package main

/*
 * @lc app=leetcode id=974 lang=golang
 *
 * [974] Subarray Sums Divisible by K
 *
 * https://leetcode.com/problems/subarray-sums-divisible-by-k/description/
 *
 * algorithms
 * Medium (47.20%)
 * Total Accepted:    49.7K
 * Total Submissions: 98.5K
 * Testcase Example:  '[4,5,0,-2,-3,1]\n5'
 *
 * Given an array A of integers, return the number of (contiguous, non-empty)
 * subarrays that have a sum divisible by K.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [4,5,0,-2,-3,1], K = 5
 * Output: 7
 * Explanation: There are 7 subarrays with a sum divisible by K = 5:
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2,
 * -3]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 30000
 * -10000 <= A[i] <= 10000
 * 2 <= K <= 10000
 *
 *
 */

func subarraysDivByK(a []int, k int) int {
	m := make([]int, k)
	r := 0
	s := 0
	m[0]++
	for _, v := range a {
		s = ((s+v)%k + k) % k
		r += m[s]
		m[s]++
	}
	return r
}
