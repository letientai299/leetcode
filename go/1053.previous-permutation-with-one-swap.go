package main

// Previous Permutation With One Swap
//
// Medium
//
// https://leetcode.com/problems/previous-permutation-with-one-swap/
//
// Given an array of positive integers `arr` (not necessarily distinct), return
// _the lexicographically largest permutation that is smaller than_ `arr`, that
// can be **made with exactly one swap** (A _swap_ exchanges the positions of
// two numbers `arr[i]` and `arr[j]`). If it cannot be done, then return the
// same array.
//
// **Example 1:**
//
// ```
// Input: arr = [3,2,1]
// Output: [3,1,2]
// Explanation: Swapping 2 and 1.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [1,1,5]
// Output: [1,1,5]
// Explanation: This is already the smallest permutation.
//
// ```
//
// **Example 3:**
//
// ```
// Input: arr = [1,9,4,6,7]
// Output: [1,7,4,6,9]
// Explanation: Swapping 9 and 7.
//
// ```
//
// **Example 4:**
//
// ```
// Input: arr = [3,1,1,3]
// Output: [1,3,1,3]
// Explanation: Swapping 1 and 3.
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 104`
// - `1 <= arr[i] <= 104`
func prevPermOpt1(arr []int) []int {
	n := len(arr)
	i := n - 1
	for ; i > 0 && arr[i] >= arr[i-1]; i-- {
	}

	if i == 0 {
		return arr
	}

	i-- // first value to swap
	j := n - 1
	for j > i+1 && (arr[j] >= arr[i] || arr[j] == arr[j-1]) {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
