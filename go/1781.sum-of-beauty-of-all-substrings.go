package main

import "math"

// Sum of Beauty of All Substrings
//
// Medium
//
// https://leetcode.com/problems/sum-of-beauty-of-all-substrings/
//
// The **beauty** of a string is the difference in frequencies between the most
// frequent and least frequent characters.
//
// - For example, the beauty of `"abaacc"` is `3 - 1 = 2`.
//
// Given a string `s`, return _the sum of **beauty** of all of its substrings._
//
// **Example 1:**
//
// ```
// Input: s = "aabcb"
// Output: 5
// Explanation: The substrings with non-zero beauty are
// ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.
// ```
//
// **Example 2:**
//
// ```
// Input: s = "aabcbaa"
// Output: 17
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 500`
// - `s` consists of only lowercase English letters.
func beautySum(s string) int {
	n := len(s)
	freqs := make([][26]int, n)
	freqs[0][int(s[0]-'a')] = 1

	for i := 1; i < n; i++ {
		copy(freqs[i][:], freqs[i-1][:])
		freqs[i][int(s[i]-'a')]++
	}

	beauty := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			min, max := 0, math.MinInt32
			a := freqs[i]
			var b [26]int
			if j > 0 {
				b = freqs[j-1]
			}
			for k := 0; k < 26; k++ {
				x := a[k] - b[k]
				if x == 0 {
					continue
				}

				if min > x || (min == 0 && x != 0) {
					min = x
				}

				if max < x {
					max = x
				}
			}

			if min > 0 {
				beauty += max - min
			}
		}
	}

	return beauty
}
