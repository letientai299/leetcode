package main

import "sort"

// Count Special Quadruplets
//
// Easy
//
// https://leetcode.com/problems/count-special-quadruplets/
//
// Given a **0-indexed** integer array `nums`, return _the number of
// **distinct** quadruplets_`(a, b, c, d)` _such that:_
//
// - `nums[a] + nums[b] + nums[c] == nums[d]`, and
// - `a < b < c < d`
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,3,6]
// Output: 1
// Explanation: The only quadruplet that satisfies the requirement is (0, 1, 2,
// 3) because 1 + 2 + 3 == 6.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [3,3,6,4,5]
// Output: 0
// Explanation: There are no such quadruplets in [3,3,6,4,5].
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,1,1,3,5]
// Output: 4
// Explanation: The 4 quadruplets that satisfy the requirement are:
// - (0, 1, 2, 3): 1 + 1 + 1 == 3
// - (0, 1, 3, 4): 1 + 1 + 3 == 5
// - (0, 2, 3, 4): 1 + 1 + 3 == 5
// - (1, 2, 3, 4): 1 + 1 + 3 == 5
//
// ```
//
// **Constraints:**
//
// - `4 <= nums.length <= 50`
// - `1 <= nums[i] <= 100`
func countQuadruplets(nums []int) int {
	n := len(nums)

	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	r := 0

	for i := 0; i < n-3; i++ {
		a := nums[i]
		for j := i + 1; j < n-2; j++ {
			b := nums[j]
			for k := j + 1; k < n-1; k++ {
				c := nums[k]
				d := a + b + c
				ids, ok := m[d]

				if !ok {
					continue
				}

				h := sort.SearchInts(ids, k)
				if h < len(ids) && ids[h] == k {
					h++
				}

				r += len(ids)-h
			}
		}
	}

	return r
}
