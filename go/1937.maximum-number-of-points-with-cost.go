package main

import (
	"math"
)

// Maximum Number of Points with Cost
//
// Medium
//
// https://leetcode.com/problems/maximum-number-of-points-with-cost/
//
// You are given an `m x n` integer matrix `points` ( **0-indexed**). Starting
// with `0` points, you want to **maximize** the number of points you can get
// from the matrix.
//
// To gain points, you must pick one cell in **each row**. Picking the cell at
// coordinates `(r, c)` will **add** `points[r][c]` to your score.
//
// However, you will lose points if you pick a cell too far from the cell that
// you picked in the previous row. For every two adjacent rows `r` and `r + 1`
// (where `0 <= r < m - 1`), picking cells at coordinates `(r, c1)` and `(r + 1,
// c2)` will **subtract** `abs(c1 - c2)` from your score.
//
// Return _the **maximum** number of points you can achieve_.
//
// `abs(x)` is defined as:
//
// - `x` for `x >= 0`.
// - `-x` for `x < 0`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-40-26-diagram-drawio-diagrams-net.png)
//
// ```
// Input: points = [[1,2,3],[1,5,1],[3,1,1]]
// Output: 9
// Explanation:
// The blue cells denote the optimal cells to pick, which have coordinates (0,
// 2), (1, 1), and (2, 0).
// You add 3 + 5 + 3 = 11 to your score.
// However, you must subtract abs(2 - 1) + abs(1 - 0) = 2 from your score.
// Your final score is 11 - 2 = 9.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-42-14-diagram-drawio-diagrams-net.png)
//
// ```
// Input: points = [[1,5],[2,3],[4,2]]
// Output: 11
// Explanation:
// The blue cells denote the optimal cells to pick, which have coordinates (0,
// 1), (1, 1), and (2, 0).
// You add 5 + 3 + 4 = 12 to your score.
// However, you must subtract abs(1 - 1) + abs(1 - 0) = 1 from your score.
// Your final score is 12 - 1 = 11.
//
// ```
//
// **Constraints:**
//
// - `m == points.length`
// - `n == points[r].length`
// - `1 <= m, n <= 105`
// - `1 <= m * n <= 105`
// - `0 <= points[r][c] <= 105`
func maxPoints1937(points [][]int) int64 {
	dp := make([][]int64, 2)
	dp[0] = make([]int64, len(points[0]))
	dp[1] = make([]int64, len(points[0]))

	for i, v := range points[0] {
		dp[0][i] = int64(v)
	}

	for i := 1; i < len(points); i++ {
		for c, x := range points[i] {
			best := int64(math.MinInt64)
			for d, y := range dp[0] {
				d -= c
				if d < 0 {
					d = -d
				}
				other := int64(x) + y - int64(d)
				if other > best {
					best = other
				}
			}
			dp[1][c] = best
		}

		dp[0], dp[1] = dp[1], dp[0]
	}

	r := int64(math.MinInt64)
	for _, v := range dp[0] {
		if r < v {
			r = v
		}
	}

	return r
}

// TODO (tai): can be O(n^2), this one is O(n^3), basically brute force
