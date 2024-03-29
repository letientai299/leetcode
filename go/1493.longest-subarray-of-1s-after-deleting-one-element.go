package main

// Longest Subarray of 1's After Deleting One Element
//
// Medium
//
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
//
// Given a binary array `nums`, you should delete one element from it.
//
// Return the size of the longest non-empty subarray containing only 1's in the
// resulting array.
//
// Return 0 if there is no such subarray.
//
// **Example 1:**
//
// ```
// Input: nums = [1,1,0,1]
// Output: 3
// Explanation: After deleting the number in position 2, [1,1,1] contains 3
// numbers with value of 1's.
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,1,1,1,0,1,1,0,1]
// Output: 5
// Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1]
// longest subarray with value of 1's is [1,1,1,1,1].
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,1,1]
// Output: 2
// Explanation: You must delete one element.
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [1,1,0,0,1,1,1,0,1]
// Output: 4
//
// ```
//
// **Example 5:**
//
// ```
// Input: nums = [0,0,0]
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 10^5`
// - `nums[i]` is either `0` or `1`.
func longestSubarray(nums []int) int {
	i := 0
	n := len(nums)
	ones := 0

	best := 0
	cur := 0

	for ; i < n; i++ {
		if nums[i] == 1 {
			ones++
			cur++
			continue
		}

		if i >= n-1 || nums[i+1] == 0 {
			if cur > best {
				best = cur
			}
			cur = 0
			continue
		}

		more := 1
		ones++
		for i += 2; i < n && nums[i] == 1; i++ {
			more++
			ones++
		}

		if more+cur > best {
			best = more + cur
		}
		cur = more
		i--
	}

	if ones == n {
		return ones - 1
	}

	if ones == 0 {
		return 0
	}

	if best < cur {
		best = cur
	}

	return best
}
