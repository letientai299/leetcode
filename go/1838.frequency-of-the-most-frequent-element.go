package main

import "sort"

// Frequency of the Most Frequent Element
//
// Medium
//
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
//
// The **frequency** of an element is the number of times it occurs in an array.
//
// You are given an integer array `nums` and an integer `k`. In one operation,
// you can choose an index of `nums` and increment the element at that index by
// `1`.
//
// Return _the **maximum possible frequency** of an element after performing
// **at most**_ `k` _operations_.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,4], k = 5
// Output: 3
// Explanation: Increment the first element three times and the second element
// two times to make nums = [4,4,4].
// 4 has a frequency of 3.
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,4,8,13], k = 5
// Output: 2
// Explanation: There are multiple optimal solutions:
// - Increment the first element three times to make nums = [4,4,8,13]. 4 has a
// frequency of 2.
// - Increment the second element four times to make nums = [1,8,8,13]. 8 has a
// frequency of 2.
// - Increment the third element five times to make nums = [1,4,13,13]. 13 has a
// frequency of 2.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [3,9,6], k = 2
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `1 <= nums[i] <= 105`
// - `1 <= k <= 105`
// TODO (tai): too slow, but still pass, try again with sliding windows
func maxFrequency(nums []int, k int) int {
	m := make(map[int]int)
	n := 0
	for _, v := range nums {
		if m[v] == 0 {
			nums[n] = v
			n++
		}
		m[v]++
	}
	nums = nums[:n]

	sort.Ints(nums)

	best := m[nums[0]]
	for i := 1; i < n; i++ {
		v := nums[i]
		freq := m[v]

		for rem, j := k, i-1; j >= 0; j-- {
			diff := v - nums[j]
			if rem < diff {
				break
			}

			f := rem / diff
			if f > m[nums[j]] {
				f = m[nums[j]]
			}

			freq += f
			rem -= diff * f
		}

		if best < freq {
			best = freq
		}
	}
	return best
}
