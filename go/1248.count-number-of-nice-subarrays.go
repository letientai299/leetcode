package main

// Count Number of Nice Subarrays
//
// # Medium
//
// https://leetcode.com/problems/count-number-of-nice-subarrays
//
// Given an array of integers `nums` and an integer `k`. A continuous subarray
// is called **nice** if there are `k` odd numbers on it.
//
// Return _the number of **nice** sub-arrays_.
//
// **Example 1:**
//
// ```
// Input: nums = [1,1,2,1,1], k = 3
// Output: 2
// Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and
// [1,2,1,1].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [2,4,6], k = 1
// Output: 0
// Explanation: There is no odd numbers in the array.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
// Output: 16
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 50000`
// - `1 <= nums[i] <= 10^5`
// - `1 <= k <= nums.length`
func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	l := 0

	// loop still first odd
	for ; l < n && nums[l]%2 == 0; l++ {
	}

	r := l
	cnt := 0

	// loop still cnt==k, r stops at an odd
	for ; r < n && cnt < k; r++ {
		if nums[r]%2 == 1 {
			cnt++
		}
		if cnt == k {
			break
		}
	}

	a := l + 1
	res := 0
	for l <= r-k+1 && r < n {
		// count evens from r and move to the next odd
		b := 1
		for r++; r < n && nums[r]%2 == 0; r++ {
			b++
		}

		res += a * b

		a = 0
		for l++; l < n && nums[l]%2 == 0; l++ {
			a++
		}
		a++
	}

	return res
}
