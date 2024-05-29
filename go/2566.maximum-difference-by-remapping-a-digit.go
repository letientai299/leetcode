package main

// Maximum Difference by Remapping a Digit
//
// # Easy
//
// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit
//
// You are given an integer `num`. You know that Bob will sneakily **remap** one
// of the `10` possible digits (`0` to `9`) to another digit.
//
// Return _the difference between the maximum and minimum values Bob can make by
// remapping **exactly** **one** digit in_ `num`.
//
// **Notes:**
//
// - When Bob remaps a digit d1 to another digit d2, Bob replaces all
// occurrences of `d1` in `num` with `d2`.
// - Bob can remap a digit to itself, in which case `num` does not change.
// - Bob can remap different digits for obtaining minimum and maximum values
// respectively.
// - The resulting number after remapping can contain leading zeroes.
//
// **Example 1:**
//
// ```
// Input: num = 11891
// Output: 99009
// Explanation:
// To achieve the maximum value, Bob can remap the digit 1 to the digit 9 to
// yield 99899.
// To achieve the minimum value, Bob can remap the digit 1 to the digit 0,
// yielding 890.
// The difference between these two numbers is 99009.
//
// ```
//
// **Example 2:**
//
// ```
// Input: num = 90
// Output: 99
// Explanation:
// The maximum value that can be returned by the function is 99 (if 0 is
// replaced by 9) and the minimum value that can be returned by the function is
// 0 (if 9 is replaced by 0).
// Thus, we return 99.
// ```
//
// **Constraints:**
//
// - `1 <= num <= 108`
func minMaxDifference(num int) int {
	ds := make([]int, 0, 9)
	for num > 0 {
		d := num % 10
		num /= 10
		ds = append(ds, d)
	}

	n := len(ds)
	maxDigitID := n - 1
	for ; maxDigitID >= 0 && ds[maxDigitID] == 9; maxDigitID-- {
	}

	maxD := 9
	if maxDigitID >= 0 {
		maxD = ds[maxDigitID]
	}

	minD := ds[n-1]
	min := 0
	max := 0
	for i := n - 1; i >= 0; i-- {
		d := ds[i]
		if d != minD {
			min = min*10 + d
		} else {
			min = min * 10
		}

		if d == maxD {
			max = max*10 + 9
		} else {
			max = max*10 + d
		}
	}

	return max - min
}
