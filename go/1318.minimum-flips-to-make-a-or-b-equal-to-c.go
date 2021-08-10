package main

// Minimum Flips to Make a OR b Equal to c
//
// Medium
//
// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
//
// Given 3 positives numbers `a`, `b` and `c`. Return the minimum flips required
// in some bits of `a` and `b` to make ( `a` OR `b` == `c` ). (bitwise OR
// operation).
//
// Flip operation consists of change **any** single bit 1 to 0 or change the bit
// 0 to 1 in their binary representation.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/01/06/sample_3_1676.png)
//
// ```
// abc
// ```
//
// **Example 2:**
//
// ```
// Input: a = 4, b = 2, c = 7
// Output: 1
//
// ```
//
// **Example 3:**
//
// ```
// Input: a = 1, b = 2, c = 3
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= a <= 10^9`
// - `1 <= b <= 10^9`
// - `1 <= c <= 10^9`
func minFlips(a int, b int, c int) int {
	r := 0
	for a != 0 || b != 0 || c != 0 {
		x := a & 1
		y := b & 1
		z := c & 1

		if x|y != z {
			r += z + x + y
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return r
}

// Good for interview, easy
