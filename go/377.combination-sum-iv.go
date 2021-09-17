package main

// Combination Sum IV
//
// Medium
//
// https://leetcode.com/problems/combination-sum-iv/
//
// Given an array of **distinct** integers `nums` and a target integer `target`,
// return _the number of possible combinations that add up to_ `target`.
//
// The answer is **guaranteed** to fit in a **32-bit** integer.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,3], target = 4
// Output: 7
// Explanation:
// The possible combination ways are:
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// Note that different sequences are counted as different combinations.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [9], target = 3
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 200`
// - `1 <= nums[i] <= 1000`
// - All the elements of `nums` are **unique**.
// - `1 <= target <= 1000`
//
// **Follow up:** What if negative numbers are allowed in the given array? How
// does it change the problem? What limitation we need to add to the question to
// allow negative numbers?
func combinationSum4(nums []int, target int) int {
	if target == 0 {
		return 0
	}

	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, v := range nums {
			if t >= v {
				dp[t] += dp[t-v]
			}
		}
	}
	return dp[target]
}
