package main

// Maximum Element After Decreasing and Rearranging
//
// Medium
//
// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
//
// You are given an array of positive integers `arr`. Perform some operations
// (possibly none) on `arr` so that it satisfies these conditions:
//
// - The value of the **first** element in `arr` must be `1`.
// - The absolute difference between any 2 adjacent elements must be **less than
// or equal to** `1`. In other words, `abs(arr[i] - arr[i - 1]) <= 1` for each
// `i` where `1 <= i < arr.length` ( **0-indexed**). `abs(x)` is the absolute
// value of `x`.
//
// There are 2 types of operations that you can perform any number of times:
//
// - **Decrease** the value of any element of `arr` to a **smaller positive
// integer**.
// - **Rearrange** the elements of `arr` to be in any order.
//
// Return _the **maximum** possible value of an element in_ `arr` _after
// performing the operations to satisfy the conditions_.
//
// **Example 1:**
//
// ```
// arr[1,2,2,2,1]arr
// ```
//
// **Example 2:**
//
// ```
// arr[1,100,1000]arr = [1,2,3], which arr is 3.
// ```
//
// **Example 3:**
//
// ```
// Input: arr = [1,2,3,4,5]
// Output: 5
// Explanation: The array already satisfies the conditions, and the largest
// element is 5.
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 105`
// - `1 <= arr[i] <= 109`
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	counts := make([]int, n)
	under := 0
	for _, v := range arr {
		if v <= n {
			counts[v-1]++
			under++
		}
	}

	need := 0
	for _, v := range counts {
		if v == 0 {
			need++ // need something to fill the missing number
		} else {
			// might need to downgrade some of these numbers to satisfy the needs
			need -= v - 1
			if need < 0 {
				need = 0
			}
		}
	}

	return 2*n - under - need
}
