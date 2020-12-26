package main

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (24.63%)
 * Total Accepted:    272.3K
 * Total Submissions: 938.7K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 *
 */
func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	var markAlive func(x, y int)
	markAlive = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m {
			return
		}

		if board[y][x] != 'O' {
			return
		}

		board[y][x] = '2'

		markAlive(x, y-1)
		markAlive(x, y+1)
		markAlive(x-1, y)
		markAlive(x+1, y)
	}

	// mark alive cell on border
	for r := 0; r < m; r++ {
		markAlive(0, r)
		markAlive(n-1, r)
	}

	for c := 0; c < n; c++ {
		markAlive(c, 0)
		markAlive(c, m-1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == '2' {
				board[r][c] = 'O'
			} else {
				board[r][c] = 'X'
			}
		}
	}

	return
}
