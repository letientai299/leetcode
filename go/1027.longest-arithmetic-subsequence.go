package main

// Longest Arithmetic Subsequence
//
// Medium
//
// https://leetcode.com/problems/longest-arithmetic-subsequence/
//
// Given an array `nums` of integers, return the **length** of the longest
// arithmetic subsequence in `nums`.
//
// Recall that a _subsequence_ of an array `nums` is a list `nums[i1], nums[i2],
// ..., nums[ik]` with `0 <= i1 < i2 < ... < ik <= nums.length - 1`, and that a
// sequence `seq` is _arithmetic_ if `seq[i+1] - seq[i]` are all the same value
// (for `0 <= i < seq.length - 1`).
//
// **Example 1:**
//
// ```
// Input: nums = [3,6,9,12]
// Output: 4
// Explanation:
// The whole array is an arithmetic sequence with steps of length = 3.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [9,4,7,2,10]
// Output: 3
// Explanation:
// The longest arithmetic subsequence is [4,7,10].
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [20,1,15,3,10,5,8]
// Output: 4
// Explanation:
// The longest arithmetic subsequence is [20,15,10,5].
//
// ```
//
// **Constraints:**
//
// - `2 <= nums.length <= 1000`
// - `0 <= nums[i] <= 500`
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	mem := make([][1002]int, n)

	best := 2
	for i := 1; i < n; i++ {
		x := nums[i]
		for j := 0; j < i; j++ {
			y := nums[j]
			d := 500 + (x - y)
			if mem[j][d] > 0 {
				mem[i][d] = mem[j][d] + 1
			} else {
				mem[i][d] = 2
			}

			if best < mem[i][d] {
				best = mem[i][d]
			}
		}
	}

	return best
}
