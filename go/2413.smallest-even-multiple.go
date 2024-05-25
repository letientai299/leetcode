package main

// Smallest Even Multiple
//
// # Easy
//
// https://leetcode.com/problems/smallest-even-multiple
//
// Given a **positive** integer `n`, return _the smallest positive integer that
// is a multiple of **both**_ `2` _and_ `n`.
//
// **Example 1:**
//
// ```
// Input: n = 5
// Output: 10
// Explanation: The smallest multiple of both 5 and 2 is 10.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 6
// Output: 6
// Explanation: The smallest multiple of both 6 and 2 is 6. Note that a number
// is a multiple of itself.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 150`
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}