package main

// Flip String to Monotone Increasing
//
// Medium
//
// https://leetcode.com/problems/flip-string-to-monotone-increasing/
//
// A binary string is monotone increasing if it consists of some number of `0`'s
// (possibly none), followed by some number of `1`'s (also possibly none).
//
// You are given a binary string `s`. You can flip `s[i]` changing it from `0`
// to `1` or from `1` to `0`.
//
// Return _the minimum number of flips to make_ `s` _monotone increasing_.
//
// **Example 1:**
//
// ```
// Input: s = "00110"
// Output: 1
// Explanation: We flip the last digit to get 00111.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "010110"
// Output: 2
// Explanation: We flip to get 011111, or alternatively 000111.
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "00011000"
// Output: 2
// Explanation: We flip to get 00000000.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 105`
// - `s[i]` is either `'0'` or `'1'`.
func minFlipsMonoIncr(s string) int {
	ones := 0
	for _, c := range s {
		ones += int(c - '0')
	}

	n := len(s)
	if ones == 0 || ones == n {
		return 0 // no need to flip any bit
	}

	left := 0 // count 1
	best := ones
	if n-ones < best {
		best = n - ones
	}

	for i, c := range s {
		left += int(c - '0')
		flips := left + (n - i - 1 - (ones - left))
		if best > flips {
			best = flips
		}
	}
	return best
}
