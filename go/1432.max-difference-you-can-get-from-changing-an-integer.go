package main

import "math"

// Max Difference You Can Get From Changing an Integer
//
// Medium
//
// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
//
// You are given an integer `num`. You will apply the following steps exactly
// **two** times:
//
// - Pick a digit `x (0 <= x <= 9)`.
// - Pick another digit `y (0 <= y <= 9)`. The digit `y` can be equal to `x`.
// - Replace all the occurrences of `x` in the decimal representation of `num`
// by `y`.
// - The new integer **cannot** have any leading zeros, also the new integer
// **cannot** be 0.
//
// Let `a` and `b` be the results of applying the operations to `num` the first
// and second times, respectively.
//
// Return _the max difference_ between `a` and `b`.
//
// **Example 1:**
//
// ```
// Input: num = 555
// Output: 888
// Explanation: The first time pick x = 5 and y = 9 and store the new integer in
// a.
// The second time pick x = 5 and y = 1 and store the new integer in b.
// We have now a = 999 and b = 111 and max difference = 888
//
// ```
//
// **Example 2:**
//
// ```
// Input: num = 9
// Output: 8
// Explanation: The first time pick x = 9 and y = 9 and store the new integer in
// a.
// The second time pick x = 9 and y = 1 and store the new integer in b.
// We have now a = 9 and b = 1 and max difference = 8
//
// ```
//
// **Example 3:**
//
// ```
// Input: num = 123456
// Output: 820000
//
// ```
//
// **Example 4:**
//
// ```
// Input: num = 10000
// Output: 80000
//
// ```
//
// **Example 5:**
//
// ```
// Input: num = 9288
// Output: 8700
//
// ```
//
// **Constraints:**
//
// - `1 <= num <= 10^8`
func maxDiff(num int) int {
	digits := make([]int, 0, 9)
	mask := 0
	for num > 0 {
		d := num % 10
		mask |= 1 << d
		digits = append(digits, d)
		num /= 10
	}

	replace := func(a, b int) int {
		r := 0
		pow := 1
		for _, v := range digits {
			if v == a {
				v = b
			}

			r += pow * v
			pow *= 10
		}
		return r
	}

	down := math.MaxInt32
	up := math.MinInt32
	for a := 0; a <= 9; a++ {
		if mask&(1<<a) == 0 {
			continue
		}

		for b := 0; b <= 9; b++ {
			if b == 0 && a == digits[len(digits)-1] {
				continue
			}

			r := replace(a, b)
			if r > up {
				up = r
			}

			if r < down {
				down = r
			}
		}
	}

	return up - down
}
