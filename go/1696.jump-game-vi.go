package main

import "sort"

// medium

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ks := make([]int, 0, k)
	ks = append(ks, nums[0])
	for i := 1; i < n; i++ {
		dp[i] = ks[len(ks)-1] + nums[i]
		if len(ks) == k {
			// remove furthest element from the sorted list
			r := sort.SearchInts(ks, dp[i-k])
			if r == 0 {
				ks = ks[1:]
			} else {
				ks = append(ks[:r], ks[r+1:]...)
			}
		}

		// insert new value into sorted list
		j := sort.SearchInts(ks, dp[i])
		if j == -1 {
			ks = append([]int{dp[i]}, ks...)
		} else {
			ks = append(ks[:j], append([]int{dp[i]}, ks[j:]...)...)
		}
	}
	return dp[n-1]
}
