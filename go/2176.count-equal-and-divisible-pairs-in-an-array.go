package main

// Count Equal and Divisible Pairs in an Array
//
// Easy
//
// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
//
// Given a **0-indexed** integer array `nums` of length `n` and an integer `k`,
// return _the **number of pairs**_`(i, j)` _where_ `0 <= i < j < n`, _such
// that_ `nums[i] == nums[j]` _and_`(i * j)` _is divisible by_ `k`.
//
// **Example 1:**
//
// ```
// Input: nums = [3,1,2,2,2,1,3], k = 2
// Output: 4
// Explanation:
// There are 4 pairs that meet all the requirements:
// - nums[0] == nums[6], and 0 * 6 == 0, which is divisible by 2.
// - nums[2] == nums[3], and 2 * 3 == 6, which is divisible by 2.
// - nums[2] == nums[4], and 2 * 4 == 8, which is divisible by 2.
// - nums[3] == nums[4], and 3 * 4 == 12, which is divisible by 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,3,4], k = 1
// Output: 0
// Explanation: Since no value in nums is repeated, there are no pairs (i,j)
// that meet all the requirements.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 100`
// - `1 <= nums[i], k <= 100`
func countPairs(nums []int, k int) int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	r := 0
	for _, is := range m {
		for i := 0; i < len(is)-1; i++ {
			for j := i + 1; j < len(is); j++ {
				if is[i]*is[j]%k == 0 {
					r++
				}
			}
		}
	}
	return r
}
