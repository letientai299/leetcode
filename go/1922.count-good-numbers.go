package main

// Count Good Numbers
//
// Medium
//
// https://leetcode.com/problems/count-good-numbers/
//
// A digit string is **good** if the digits **(0-indexed)** at **even** indices
// are **even** and the digits at **odd** indices are **prime** (`2`, `3`, `5`,
// or `7`).
//
// - For example, `"2582"` is good because the digits (`2` and `8`) at even
// positions are even and the digits (`5` and `2`) at odd positions are prime.
// However, `"3245"` is **not** good because `3` is at an even index but is not
// even.
//
// Given an integer `n`, return _the **total** number of good digit strings of
// length_ `n`. Since the answer may be large, **return it modulo** `109 + 7`.
//
// A **digit string** is a string consisting of digits `0` through `9` that may
// contain leading zeros.
//
// **Example 1:**
//
// ```
// Input: n = 1
// Output: 5
// Explanation: The good numbers of length 1 are "0", "2", "4", "6", "8".
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 4
// Output: 400
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 50
// Output: 564908303
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 1015`
func countGoodNumbers(n int64) int {
	odds := n / 2

	// 10 digits: 5 evens are usable for evens indices
	// 5^evens * 4^odds
	r := 1
	const mod = 1e9 + 7

	cur := 20
	for odds > 0 {
		if odds%2 == 0 {
			odds /= 2
			cur *= cur
			cur %= mod
		} else {
			odds--
			r = (r * cur) % mod
		}
	}

	if n%2 == 1 {
		r = (r * 5) % mod
	}
	return r
}
