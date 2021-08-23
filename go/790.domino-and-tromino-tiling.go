package main

// Domino and Tromino Tiling
//
// Medium
//
// https://leetcode.com/problems/domino-and-tromino-tiling/
//
// You have two types of tiles: a `2 x 1` domino shape and a tromino shape. You
// may rotate these shapes.
//
// ![](https://assets.leetcode.com/uploads/2021/07/15/lc-domino.jpg)
//
// Given an integer n, return _the number of ways to tile an_ `2 x n` _board_.
// Since the answer may be very large, return it **modulo** `109 + 7`.
//
// In a tiling, every square must be covered by a tile. Two tilings are
// different if and only if there are two 4-directionally adjacent cells on the
// board such that exactly one of the tilings has both squares occupied by a
// tile.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/15/lc-domino1.jpg)
//
// ```
// Input: n = 3
// Output: 5
// Explanation: The five different ways are show above.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 1000`
func numTilings(n int) int {
	mem := []int{1, 2, 5}
	for n > 3 {
		v := (mem[2]*2 + mem[0]) % (1e9 + 7)
		mem[0], mem[1], mem[2] = mem[1], mem[2], v
		n--
	}

	// I looked at the sample output and guess the DP formula, still haven't
	// proved its correctness yet.
	return mem[n-1]
}
