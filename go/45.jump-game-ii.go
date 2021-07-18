package main

// Jump Game II
//
// Medium
//
// https://leetcode.com/problems/jump-game-ii/
//
// Given an array of non-negative integers `nums`, you are initially positioned
// at the first index of the array.
//
// Each element in the array represents your maximum jump length at that
// position.
//
// Your goal is to reach the last index in the minimum number of jumps.
//
// You can assume that you can always reach the last index.
//
// **Example 1:**
//
// ```
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1
// step from index 0 to 1, then 3 steps to the last index.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [2,3,0,1,4]
// Output: 2
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 104`
// - `0 <= nums[i] <= 1000`
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	maxBefore := nums[0]
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && i-j <= maxBefore; j-- {
			if nums[j] < i-j {
				continue
			}
			v := dp[j] + 1
			if dp[i] == 0 || dp[i] > v {
				dp[i] = v
			}
		}

		if maxBefore < nums[i] {
			maxBefore = nums[i]
		}
	}

	return dp[len(nums)-1]
}

// TODO (tai): slow, 80ms, 19.65%, can be 0ms
