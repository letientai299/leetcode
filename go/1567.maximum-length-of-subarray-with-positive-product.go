package main

// Maximum Length of Subarray With Positive Product
//
// Medium
//
// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/
//
// Given an array of integers `nums`, find the maximum length of a subarray
// where the product of all its elements is positive.
//
// A subarray of an array is a consecutive sequence of zero or more values taken
// out of that array.
//
// Return _the maximum length of a subarray with positive product_.
//
// **Example 1:**
//
// ```
// Input: nums = [1,-2,-3,4]
// Output: 4
// Explanation: The array nums already has a positive product of 24.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,1,-2,-3,-4]
// Output: 3
// Explanation: The longest subarray with positive product is [1,-2,-3] which
// has a product of 6.
// Notice that we cannot include 0 in the subarray since that'll make the
// product 0 which is not positive.
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [-1,-2,-3,0,1]
// Output: 2
// Explanation: The longest subarray with positive product is [-1,-2] or
// [-2,-3].
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [-1,2]
// Output: 1
//
// ```
//
// **Example 5:**
//
// ```
// Input: nums = [1,2,3,5,-6,4,0,10]
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 10^5`
// - `-10^9 <= nums[i] <= 10^9`
func getMaxLen(nums []int) int {
	best := 0
	i := 0
	n := len(nums)
	for ; i < n; i++ {
		for i < n && nums[i] == 0 {
			i++
		}
		if i >= n {
			break
		}
		cur := 0
		sig := 1
		j := i
		for ; j < n && nums[j] != 0; j++ {
			cur++
			if nums[j] < 0 {
				sig *= -1
			}
			if sig > 0 && cur > best {
				best = cur
			}
		}

		for ; i < j && abs(cur) > best; i++ {
			cur--
			if nums[i] < 0 {
				sig *= -1
			}
			if sig > 0 && cur > best {
				best = cur
			}
		}
		i = j
	}
	return best
}

// TODO (tai): can be faster, 139 ms, 50.00%
