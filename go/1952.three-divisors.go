package main

import "math"

// Three Divisors
//
// Easy
//
// https://leetcode.com/problems/three-divisors/
//
// Given an integer `n`, return `true` _if_ `n` _has **exactly three positive
// divisors**. Otherwise, return_ `false`.
//
// An integer `m` is a **divisor** of `n` if there exists an integer `k` such
// that `n = k * m`.
//
// **Example 1:**
//
// ```
// Input: n = 2
// Output: false
// Explantion: 2 has only two divisors: 1 and 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 4
// Output: true
// Explantion: 4 has three divisors: 1, 2, and 4.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 104`
func isThree(n int) bool {
	if n <= 3 {
		return false
	}

	x := int(math.Sqrt(float64(n)))
	return x*x == n && isPrime(x)
}
