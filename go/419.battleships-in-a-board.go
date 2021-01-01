package main

/*
 * @lc app=leetcode id=419 lang=golang
 *
 * [419] Battleships in a Board
 *
 * https://leetcode.com/problems/battleships-in-a-board/description/
 *
 * algorithms
 * Medium (67.16%)
 * Total Accepted:    106.1K
 * Total Submissions: 149.8K
 * Testcase Example:  '[["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]'
 *
 * Given an 2D board, count how many battleships are in it. The battleships are
 * represented with 'X's, empty slots are represented with '.'s. You may assume
 * the following rules:
 *
 *
 * You receive a valid board, made of only battleships or empty slots.
 * Battleships can only be placed horizontally or vertically. In other words,
 * they can only be made of the shape 1xN (1 row, N columns) or Nx1 (N rows, 1
 * column), where N can be of any size.
 * At least one horizontal or vertical cell separates between two battleships -
 * there are no adjacent battleships.
 *
 *
 * Example:
 * X..X
 * ...X
 * ...X
 *
 * In the above board there are 2 battleships.
 *
 * Invalid Example:
 * ...X
 * XXXX
 * ...X
 *
 * This is an invalid board that you will not receive - as battleships will
 * always have a cell separating between them.
 *
 * Follow up:Could you do it in one-pass, using only O(1) extra memory and
 * without modifying the value of the board?
 */
func countBattleships(board [][]byte) int {
	m := len(board)
	if m == 0 {
		return 0
	}

	n := len(board[0])
	if n == 0 {
		return 0
	}

	res := 0
	check := func(x, y int) {
		if board[y][x] == '.' {
			return
		}

		board[y][x] = '.'
		res++
		for x < n-1 && board[y][x+1] == 'X' {
			x++
			board[y][x] = '.'
		}

		for y < m-1 && board[y+1][x] == 'X' {
			y++
			board[y][x] = '.'
		}
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			check(x, y)
		}
	}

	return res
}
