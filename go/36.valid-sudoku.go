package main

// medium
// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m := make(map[byte]bool)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[3*i+x][3*j+y] == '.' {
						continue
					}

					if m[board[3*i+x][3*j+y]] {
						return false
					}
					m[board[3*i+x][3*j+y]] = true
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		r := make(map[byte]bool)
		c := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if r[board[i][j]] {
					return false
				}
				r[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if c[board[j][i]] {
					return false
				}
				c[board[j][i]] = true
			}
		}
	}

	return true
}
