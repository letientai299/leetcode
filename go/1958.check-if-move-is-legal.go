package main

// Check if Move is Legal
//
// Medium
//
// https://leetcode.com/problems/check-if-move-is-legal/
//
// You are given a **0-indexed** `8 x 8` grid `board`, where `board[r][c]`
// represents the cell `(r, c)` on a game board. On the board, free cells are
// represented by `'.'`, white cells are represented by `'W'`, and black cells
// are represented by `'B'`.
//
// Each move in this game consists of choosing a free cell and changing it to
// the color you are playing as (either white or black). However, a move is only
// **legal** if, after changing it, the cell becomes the **endpoint of a good
// line** (horizontal, vertical, or diagonal).
//
// A **good line** is a line of **three or more cells (including the
// endpoints)** where the endpoints of the line are **one color**, and the
// remaining cells in the middle are the **opposite color** (no cells in the
// line are free). You can find examples for good lines in the figure below:
//
// ![](https://assets.leetcode.com/uploads/2021/07/22/goodlines5.png)
//
// Given two integers `rMove` and `cMove` and a character `color` representing
// the color you are playing as (white or black), return `true` _if changing
// cell_`(rMove, cMove)` _to color_ `color` _is a **legal** move, or_ `false`
// _if it is not legal_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/10/grid11.png)
//
// ```
// Input: board =
// [[".",".",".","B",".",".",".","."],[".",".",".","W",".",".",".","."],[".",".",".","W",".",".",".","."],[".",".",".","W",".",".",".","."],["W","B","B",".","W","W","W","B"],[".",".",".","B",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","W",".",".",".","."]],
// rMove = 4, cMove = 3, color = "B"
// Output: true
// Explanation: '.', 'W', and 'B' are represented by the colors blue, white, and
// black respectively, and cell (rMove, cMove) is marked with an 'X'.
// The two good lines with the chosen cell as an endpoint are annotated above
// with the red rectangles.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/10/grid2.png)
//
// ```
// Input: board =
// [[".",".",".",".",".",".",".","."],[".","B",".",".","W",".",".","."],[".",".","W",".",".",".",".","."],[".",".",".","W","B",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".","B","W",".","."],[".",".",".",".",".",".","W","."],[".",".",".",".",".",".",".","B"]],
// rMove = 4, cMove = 4, color = "W"
// Output: false
// Explanation: While there are good lines with the chosen cell as a middle
// cell, there are no good lines with the chosen cell as an endpoint.
//
// ```
//
// **Constraints:**
//
// - `board.length == board[r].length == 8`
// - `0 <= rMove, cMove < 8`
// - `board[rMove][cMove] == '.'`
// - `color` is either `'B'` or `'W'`.
func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	m := len(board)
	n := len(board[0])

	board[rMove][cMove] = color
	mid := byte('W')

	if mid == color {
		mid = 'B'
	}

	get := func(x, y int) byte {
		if x < 0 || n <= x || y < 0 || m <= y {
			return 0
		}

		return board[y][x]
	}

	var walk func(x, y, dx, dy, step int) bool
	walk = func(x, y, dx, dy, step int) bool {
		next := get(x+dx, y+dy)
		if x == cMove && y == rMove {
			return next != '.' && next != 0 && walk(x+dx, y+dy, dx, dy, 2)
		}

		if next == color {
			return board[y][x] == mid && step >= 2
		}

		if next == 0 || next == '.' {
			return board[y][x] == color && step >= 3
		}

		return board[y][x] == mid && walk(x+dx, y+dy, dx, dy, step+1)
	}

	return walk(cMove, rMove, 1, 0, 1) || // right
		walk(cMove, rMove, -1, 0, 1) || // left
		walk(cMove, rMove, 0, 1, 1) || // down
		walk(cMove, rMove, 0, -1, 1) || // up
		walk(cMove, rMove, 1, 1, 1) || // down-right
		walk(cMove, rMove, 1, -1, 1) || // up-right
		walk(cMove, rMove, -1, -1, 1) || // up-left
		walk(cMove, rMove, -1, 1, 1) // down-left
}
