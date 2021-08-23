package main

import "fmt"

// Sudoku Solver
//
// Hard
//
// https://leetcode.com/problems/sudoku-solver/
//
// Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy **all of the following rules**:
//
// 1. Each of the digits `1-9` must occur exactly once in each row.
// 2. Each of the digits `1-9` must occur exactly once in each column.
// 3. Each of the digits `1-9` must occur exactly once in each of the 9 `3x3`
// sub-boxes of the grid.
//
// The `'.'` character indicates empty cells.
//
// **Example 1:**
//
// ![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)
//
// ```
// Input: board =
// [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
// Output:
// [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
// Explanation: The input board is shown above and the only valid solution is
// shown below:
//
// ```
//
// **Constraints:**
//
// - `board.length == 9`
// - `board[i].length == 9`
// - `board[i][j]` is a digit or `'.'`.
// - It is **guaranteed** that the input board has only one solution.
func solveSudoku(board [][]byte) {
	rows := make([]int, 9)
	cols := make([]int, 9)
	regions := make([]int, 9)

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] != '.' {
				v := board[y][x] - '0'
				rows[y] |= 1 << v
				cols[x] |= 1 << v
				regionID := y/3*3 + x/3
				regions[regionID] |= 1 << v
			}
		}
	}

	var walk func(y, x int) bool

	walk = func(y, x int) bool {
		if y >= 9 {
			return true
		}

		if x >= 9 {
			return walk(y+1, 0)
		}

		v := board[y][x]
		if v != '.' {
			return walk(y, x+1)
		}

		r := y/3*3 + x/3
		for v = 1; v <= 9; v++ {
			mask := 1 << v
			if mask&regions[r] == 0 &&
				mask&rows[y] == 0 &&
				mask&cols[x] == 0 {

				regions[r] |= mask
				rows[y] |= mask
				cols[x] |= mask

				board[y][x] = v + '0'
				if walk(y, x+1) {
					return true
				}

				board[y][x] = '.'
				regions[r] &= ^mask
				rows[y] &= ^mask
				cols[x] &= ^mask
			}
		}

		return false
	}

	fmt.Println(walk(0, 0))
}
