package main

// Largest Subarray Length K
//
// Easy
//
// https://leetcode.com/problems/largest-subarray-length-k/
func largestSubarray(nums []int, k int) []int {
	n := len(nums)
	best := 0
	for i := 1; i <= n-k; i++ {
		if nums[i] > nums[best] {
			best = i
		}
	}
	return nums[best : best+k]
}
