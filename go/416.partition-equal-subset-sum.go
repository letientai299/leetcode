package main

// Partition Equal Subset Sum
//
// Medium
//
// https://leetcode.com/problems/partition-equal-subset-sum/
//
// Given a **non-empty** array `nums` containing **only positive integers**,
// find if the array can be partitioned into two subsets such that the sum of
// elements in both subsets is equal.
//
// **Example 1:**
//
// ```
// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,3,5]
// Output: false
// Explanation: The array cannot be partitioned into equal sum subsets.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 200`
// - `1 <= nums[i] <= 100`
func canPartition(nums []int) bool {
	// like 473, this could be solved by using the general algorithm in 698,
	// however, the range of nums length hints that 698 solution will get TLE
	// for this problem. There must be a better way to solve this.
	all := 0
	for _, v := range nums {
		all += v
	}

	if all%2 != 0 {
		return false
	}

	half := all / 2

	dp := make([]int, half+1)
	dp[0] = 1
	for _, v := range nums {
		for j := half; j >= v; j-- {
			dp[j] |= dp[j-v]
		}
	}

	return dp[half] != 0
}

// Oohhh shit: https://leetcode.com/problems/partition-equal-subset-sum/discuss/462699/Whiteboard-Editorial.-All-Approaches-explained.
// https://leetcode.com/problems/partition-equal-subset-sum/discuss/1385473/Python-1-liner-100
// This is mentioned in the first problem in The Hitchhiker's Guide to the Programming Contest
// https://comscigate.com/Books/contests/icpc.pdf
