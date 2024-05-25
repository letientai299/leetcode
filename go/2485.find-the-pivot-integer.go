package main

import (
	"math"
)

// Find the Pivot Integer
//
// # Easy
//
// https://leetcode.com/problems/find-the-pivot-integer
//
// Given a positive integer `n`, find the **pivot integer** `x` such that:
//
// - The sum of all elements between `1` and `x` inclusively equals the sum of
// all elements between `x` and `n` inclusively.
//
// Return _the pivot integer_ `x`. If no such integer exists, return `-1`. It is
// guaranteed that there will be at most one pivot index for the given input.
//
// **Example 1:**
//
// ```
// Input: n = 8
// Output: 6
// Explanation: 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8
// = 21.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1
// Output: 1
// Explanation: 1 is the pivot integer since: 1 = 1.
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 4
// Output: -1
// Explanation: It can be proved that no such integer exist.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 1000`
func pivotInteger(n int) int {
	// solve this: x*(x+1) = (n+x)*(n-x+1)
	// 2x^2  = n2 + n
	v := n*n + n
	if v%2 == 1 {
		return -1
	}

	v /= 2
	x := int(math.Floor(math.Sqrt(float64(v))))
	if x*x == v {
		return x
	}

	return -1
}
