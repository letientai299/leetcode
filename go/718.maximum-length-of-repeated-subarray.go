package main

// Maximum Length of Repeated Subarray
//
// Medium
//
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
//
// Given two integer arrays `nums1` and `nums2`, return _the maximum length of a
// subarray that appears in **both** arrays_.
//
// **Example 1:**
//
// ```
// Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// Output: 3
// Explanation: The repeated subarray with maximum length is [3,2,1].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// Output: 5
//
// ```
//
// **Constraints:**
//
// - `1 <= nums1.length, nums2.length <= 1000`
// - `0 <= nums1[i], nums2[i] <= 100`
func findLength(a []int, b []int) int {
	m := len(a)
	n := len(b)
	if m > n {
		m, n = n, m
		a, b = b, a
	}

	best := 0

	dp := make([][]int, 2)
	dp[0] = make([]int, m+1)
	dp[1] = make([]int, m+1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[j] != b[i] {
				dp[1][j+1] = 0
				continue
			}

			dp[1][j+1] = dp[0][j] + 1
			if best < dp[1][j+1] {
				best = dp[1][j+1]
			}
		}

		dp[0], dp[1] = dp[1], dp[0]
	}

	return best
}
