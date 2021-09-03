package main

// Game of Life
//
// Medium
//
// https://leetcode.com/problems/game-of-life/
//
// According to [Wikipedia's
// article](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life): "The **Game
// of Life**, also known simply as **Life**, is a cellular automaton devised by
// the British mathematician John Horton Conway in 1970."
//
// The board is made up of an `m x n` grid of cells, where each cell has an
// initial state: **live** (represented by a `1`) or **dead** (represented by a
// `0`). Each cell interacts with its [eight
// neighbors](https://en.wikipedia.org/wiki/Moore_neighborhood) (horizontal,
// vertical, diagonal) using the following four rules (taken from the above
// Wikipedia article):
//
// 1. Any live cell with fewer than two live neighbors dies as if caused by
// under-population.
// 2. Any live cell with two or three live neighbors lives on to the next
// generation.
// 3. Any live cell with more than three live neighbors dies, as if by
// over-population.
// 4. Any dead cell with exactly three live neighbors becomes a live cell, as if
// by reproduction.
//
// The next state is created by applying the above rules simultaneously to every
// cell in the current state, where births and deaths occur simultaneously.
// Given the current state of the `m x n` grid `board`, return _the next state_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/12/26/grid1.jpg)
//
// ```
// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/12/26/grid2.jpg)
//
// ```
// Input: board = [[1,1],[1,0]]
// Output: [[1,1],[1,1]]
//
// ```
//
// **Constraints:**
//
// - `m == board.length`
// - `n == board[i].length`
// - `1 <= m, n <= 25`
// - `board[i][j]` is `0` or `1`.
//
// **Follow up:**
//
// - Could you solve it in-place? Remember that the board needs to be updated
// simultaneously: You cannot update some cells first and then use their updated
// values to update other cells.
// - In this question, we represent the board using a 2D array. In principle,
// the board is infinite, which would cause problems when the active area
// encroaches upon the border of the array (i.e., live cells reach the border).
// How would you address these problems?
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			neighbors := 0
			if y > 0 {
				neighbors += board[y-1][x] & 1
				if x > 0 {
					neighbors += board[y-1][x-1] & 1
				}
				if x < n-1 {
					neighbors += board[y-1][x+1] & 1
				}
			}

			if y < m-1 {
				neighbors += board[y+1][x] & 1
				if x > 0 {
					neighbors += board[y+1][x-1] & 1
				}
				if x < n-1 {
					neighbors += board[y+1][x+1] & 1
				}
			}

			if x > 0 {
				neighbors += board[y][x-1] & 1
			}

			if x < n-1 {
				neighbors += board[y][x+1] & 1
			}

			live := 1 == (board[y][x] & 1)
			switch {
			case !live && neighbors == 3:
				board[y][x] |= 2
			case live && (neighbors == 2 || neighbors == 3):
				board[y][x] |= 2
			}
		}
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			board[y][x] >>= 1
		}
	}
}
