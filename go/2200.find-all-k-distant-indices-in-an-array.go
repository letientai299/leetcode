package main

// Find All K-Distant Indices in an Array
//
// Easy
//
// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/
//
// You are given a **0-indexed** integer array `nums` and two integers `key` and
// `k`. A **k-distant index** is an index `i` of `nums` for which there exists
// at least one index `j` such that `|i - j| <= k` and `nums[j] == key`.
//
// Return _a list of all k-distant indices sorted in **increasing order**_.
//
// **Example 1:**
//
// ```
// nums[2] == keynums[5] == key.
// - For index 0, |0 - 2| > k and |0 - 5| > k, so there is no j|0 - j| <=
// knums[j] == key. Thus, 0 is not a k-distant index.
// - For index 1, |1 - 2| <= k and nums[2] == key, so 1 is a k-distant index.
// - For index 2, |2 - 2| <= k and nums[2] == key, so 2 is a k-distant index.
// - For index 3, |3 - 2| <= k and nums[2] == key, so 3 is a k-distant index.
// - For index 4, |4 - 5| <= k and nums[5] == key, so 4 is a k-distant index.
// - For index 5, |5 - 5| <= k and nums[5] == key, so 5 is a k-distant index.
// - For index 6, |6 - 5| <= k and nums[5] == key, so 6 is a k-distant index.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [2,2,2,2,2], key = 2, k = 2
// Output: [0,1,2,3,4]
// Explanation: For all indices i in nums, there exists some index j such that
// |i - j| <= k and nums[j] == key, so every index is a k-distant index.
// Hence, we return [0,1,2,3,4].
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 1000`
// - `1 <= nums[i] <= 1000`
// - `key` is an integer from the array `nums`.
// - `1 <= k <= nums.length`
func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	r := make([]int, 0, n)

	for i, v := range nums {
		if v != key {
			continue
		}

		var j int
		if len(r) != 0 {
			j = r[len(r)-1] + 1
			if j < i-k {
				j = i - k
			}
		} else {
			j = i - k
			if j < 0 {
				j = 0
			}
		}

		for ; j <= i+k && j < n; j++ {
			r = append(r, j)
		}
	}

	return r
}
