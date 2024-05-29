package main

import "sort"

// Split With Minimum Sum
//
// # Easy
//
// https://leetcode.com/problems/split-with-minimum-sum
//
// Given a positive integer `num`, split it into two non-negative integers
// `num1` and `num2` such that:
//
// - The concatenation of `num1` and `num2` is a permutation of `num`.
//
//   - In other words, the sum of the number of occurrences of each digit in
//
// `num1` and `num2` is equal to the number of occurrences of that digit in
// `num`.
// - `num1` and `num2` can contain leading zeros.
//
// Return _the **minimum** possible sum of_ `num1` _and_ `num2`.
//
// **Notes:**
//
// - It is guaranteed that `num` does not contain any leading zeros.
// - The order of occurrence of the digits in `num1` and `num2` may differ from
// the order of occurrence of `num`.
//
// **Example 1:**
//
// ```
// num1num2
// ```
//
// **Example 2:**
//
// ```
// num1num2
// ```
//
// **Constraints:**
//
// - `10 <= num <= 109`
func splitNum(num int) int {
	var ds []int
	for num > 0 {
		d := num % 10
		if d > 0 {
			ds = append(ds, d)
		}
		num /= 10
	}

	sort.Ints(ds)

	if len(ds)%2 == 1 {
		ds = append(ds, 0)
	}

	a, b := 0, 0
	for i := 0; i < len(ds); i += 2 {
		x, y := ds[i], ds[i+1]
		a = a*10 + x
		b = b*10 + y
	}

	if ds[len(ds)-1] == 0 {
		b /= 10
	}

	return a + b
}
