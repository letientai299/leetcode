package main

// Number of Common Factors
//
// # Easy
//
// https://leetcode.com/problems/number-of-common-factors
//
// Given two positive integers `a` and `b`, return _the number of **common**
// factors of_ `a` _and_ `b`.
//
// An integer `x` is a **common factor** of `a` and `b` if `x` divides both `a`
// and `b`.
//
// **Example 1:**
//
// ```
// Input: a = 12, b = 6
// Output: 4
// Explanation: The common factors of 12 and 6 are 1, 2, 3, 6.
//
// ```
//
// **Example 2:**
//
// ```
// Input: a = 25, b = 30
// Output: 2
// Explanation: The common factors of 25 and 30 are 1, 5.
//
// ```
//
// **Constraints:**
//
// - `1 <= a, b <= 1000`
func commonFactors(a int, b int) int {
	if b < a {
		a, b = b, a
	}
	m := make(map[int]int)
	for i := 2; i <= a; i++ {
		x := 0
		for a%i == 0 {
			a /= i
			x++
		}

		y := 0
		for b%i == 0 {
			b /= i
			y++
		}

		if x > 0 && y > 0 {
			m[i] = min(x, y)
		}
	}

	res := 1
	for _, v := range m {
		res *= v + 1
	}
	return res
}
