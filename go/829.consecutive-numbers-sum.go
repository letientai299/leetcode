package main

// Consecutive Numbers Sum
//
// Hard
//
// https://leetcode.com/problems/consecutive-numbers-sum/
//
// Given an integer `n`, return _the number of ways you can write_ `n` _as the
// sum of consecutive positive integers._
//
// **Example 1:**
//
// ```
// Input: n = 5
// Output: 2
// Explanation: 5 = 2 + 3
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 9
// Output: 3
// Explanation: 9 = 4 + 5 = 2 + 3 + 4
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 15
// Output: 4
// Explanation: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 109`
func consecutiveNumbersSum(n int) int {
	// n = k + (k+1) + ... (k+x) = (x+1)*(x+2k)/2
	// 2n = (x+1)*(x+2k)
	n *= 2
	r := 0
	// left side is x+1, min k must be 1, hence the stop condition
	for x := 0; (x+1)*(x+2) <= n; x++ {
		if n%(x+1) == 0 && (n/(x+1)-x)%2 == 0 {
			r++
		}
	}
	return r
}
