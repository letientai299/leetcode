package main

import "math"

// Coin Change
//
// Medium
//
// https://leetcode.com/problems/coin-change/
//
// You are given an integer array `coins` representing coins of different
// denominations and an integer `amount` representing a total amount of money.
//
// Return _the fewest number of coins that you need to make up that amount_. If
// that amount of money cannot be made up by any combination of the coins,
// return `-1`.
//
// You may assume that you have an infinite number of each kind of coin.
//
// **Example 1:**
//
// ```
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: coins = [2], amount = 3
// Output: -1
//
// ```
//
// **Example 3:**
//
// ```
// Input: coins = [1], amount = 0
// Output: 0
//
// ```
//
// **Example 4:**
//
// ```
// Input: coins = [1], amount = 1
// Output: 1
//
// ```
//
// **Example 5:**
//
// ```
// Input: coins = [1], amount = 2
// Output: 2
//
// ```
//
// **Constraints:**
//
// - `1 <= coins.length <= 12`
// - `1 <= coins[i] <= 231 - 1`
// - `0 <= amount <= 104`
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for n := 1; n <= amount; n++ {
		dp[n] = math.MaxInt32
		for i := 0; i < len(coins); i++ {
			x := coins[i]
			if n-x >= 0 {
				v := dp[n-x] + 1
				if v < dp[n] {
					dp[n] = v
				}
			}
		}
	}

	v := dp[amount]
	if v == math.MaxInt32 {
		return -1
	}
	return v
}
