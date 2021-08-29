package main

// Knight Probability in Chessboard
//
// Medium
//
// https://leetcode.com/problems/knight-probability-in-chessboard/
//
// On an `n x n` chessboard, a knight starts at the cell `(row, column)` and
// attempts to make exactly `k` moves. The rows and columns are **0-indexed**,
// so the top-left cell is `(0, 0)`, and the bottom-right cell is `(n - 1, n -
// 1)`.
//
// A chess knight has eight possible moves it can make, as illustrated below.
// Each move is two cells in a cardinal direction, then one cell in an
// orthogonal direction.
//
// ![](https://assets.leetcode.com/uploads/2018/10/12/knight.png)
//
// Each time the knight is to move, it chooses one of eight possible moves
// uniformly at random (even if the piece would go off the chessboard) and moves
// there.
//
// The knight continues moving until it has made exactly `k` moves or has moved
// off the chessboard.
//
// Return _the probability that the knight remains on the board after it has
// stopped moving_.
//
// **Example 1:**
//
// ```
// Input: n = 3, k = 2, row = 0, column = 0
// Output: 0.06250
// Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight
// on the board.
// From each of those positions, there are also two moves that will keep the
// knight on the board.
// The total probability the knight stays on the board is 0.0625.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1, k = 0, row = 0, column = 0
// Output: 1.00000
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 25`
// - `0 <= k <= 100`
// - `0 <= row, column <= n`
func knightProbability(n int, k int, row int, column int) float64 {
	if row == n || column == n {
		return 0
	}

	directions := []int{
		// dx, dy
		-1, -2,
		1, -2,
		2, -1,
		2, 1,
		1, 2,
		-1, 2,
		-2, 1,
		-2, -1,
	}

	// once again, slice/array is much faster than map.
	// if mem is []map[int]float64, this took 23ms, but with [][]float64, it's 4ms.
	mem := make([][]float64, k+1)
	for i := range mem {
		mem[i] = make([]float64, 31*31)
	}

	var walk func(k int, x int, y int) float64
	walk = func(k int, x int, y int) float64 {
		if k == 0 {
			return 1 // can't move, thus, can't go out of the board
		}

		key := y*n + x
		if v := mem[k][key]; v != 0 {
			return v
		}

		var on float64
		for i := 0; i < len(directions); i += 2 {
			dx, dy := directions[i], directions[i+1]
			u := x + dx
			v := y + dy
			if 0 <= u && u < n && 0 <= v && v < n {
				on += walk(k-1, u, v)
			}
		}

		v := on / 8.0
		mem[k][key] = v
		return v
	}
	return walk(k, column, row)
}
