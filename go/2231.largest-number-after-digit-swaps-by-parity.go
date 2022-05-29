package main

import "sort"

// Largest Number After Digit Swaps by Parity
//
// Easy
//
// https://leetcode.com/problems/largest-number-after-digit-swaps-by-parity/
//
// You are given a positive integer `num`. You may swap any two digits of `num`
// that have the same **parity** (i.e. both odd digits or both even digits).
//
// Return _the **largest** possible value of_ `num` _after **any** number of
// swaps._
//
// **Example 1:**
//
// ```
// Input: num = 1234
// Output: 3412
// Explanation: Swap the digit 3 with the digit 1, this results in the number
// 3214.
// Swap the digit 2 with the digit 4, this results in the number 3412.
// Note that there may be other sequences of swaps but it can be shown that 3412
// is the largest possible number.
// Also note that we may not swap the digit 4 with the digit 1 since they are of
// different parities.
//
// ```
//
// **Example 2:**
//
// ```
// Input: num = 65875
// Output: 87655
// Explanation: Swap the digit 8 with the digit 6, this results in the number
// 85675.
// Swap the first digit 5 with the digit 7, this results in the number 87655.
// Note that there may be other sequences of swaps but it can be shown that
// 87655 is the largest possible number.
//
// ```
//
// **Constraints:**
//
// - `1 <= num <= 109`
func largestInteger(num int) int {
	all := make([]int, 0, 10)
	odds := make([]int, 0, 10)
	evens := make([]int, 0, 10)
	for num > 0 {
		d := num % 10
		num /= 10
		all = append(all, d)
		if d%2 == 0 {
			evens = append(evens, d)
		} else {
			odds = append(odds, d)
		}
	}
	sort.Ints(odds)
	sort.Ints(evens)
	n := 0

	o := len(odds) - 1
	e := len(evens) - 1
	for i := len(all) - 1; i >= 0; i-- {
		d := all[i]
		if d%2 == 0 {
			n = 10*n + evens[e]
			e--
		} else {
			n = 10*n + odds[o]
			o--
		}
	}
	return n
}
