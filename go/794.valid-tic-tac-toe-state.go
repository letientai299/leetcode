package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=794 lang=golang
 *
 * [794] Valid Tic-Tac-Toe State
 *
 * https://leetcode.com/problems/valid-tic-tac-toe-state/description/
 *
 * algorithms
 * Medium (31.65%)
 * Total Accepted:    29.8K
 * Total Submissions: 88.7K
 * Testcase Example:  '["O  ","   ","   "]'
 *
 * A Tic-Tac-Toe board is given as a string array board. Return True if and
 * only if it is possible to reach this board position during the course of a
 * valid tic-tac-toe game.
 *
 * The board is a 3 x 3 array, and consists of characters " ", "X", and "O".
 * The " " character represents an empty square.
 *
 * Here are the rules of Tic-Tac-Toe:
 *
 *
 * Players take turns placing characters into empty squares (" ").
 * The first player always places "X" characters, while the second player
 * always places "O" characters.
 * "X" and "O" characters are always placed into empty squares, never filled
 * ones.
 * The game ends when there are 3 of the same (non-empty) character filling any
 * row, column, or diagonal.
 * The game also ends if all squares are non-empty.
 * No more moves can be played if the game is over.
 *
 *
 *
 * Example 1:
 * Input: board = ["O  ", "   ", "   "]
 * Output: false
 * Explanation: The first player always plays "X".
 *
 * Example 2:
 * Input: board = ["XOX", " X ", "   "]
 * Output: false
 * Explanation: Players take turns making moves.
 *
 * Example 3:
 * Input: board = ["XXX", "   ", "OOO"]
 * Output: false
 *
 * Example 4:
 * Input: board = ["XOX", "O O", "XOX"]
 * Output: true
 *
 *
 * Note:
 *
 *
 * board is a length-3 array of strings, where each string board[i] has length
 * 3.
 * Each board[i][j] is a character in the set {" ", "X", "O"}.
 *
 *
 */
func validTicTacToe(b []string) bool {
	stat := func(c byte) (count int, win bool) {
		for _, s := range b {
			count += strings.Count(s, string(c))
		}
		w := string([]byte{c, c, c})
		win = b[0] == w || b[1] == w || b[2] == w ||
			(b[0][0] == c && b[1][0] == c && b[2][0] == c) ||
			(b[0][1] == c && b[1][1] == c && b[2][1] == c) ||
			(b[0][2] == c && b[1][2] == c && b[2][2] == c) ||
			(b[0][0] == c && b[1][1] == c && b[2][2] == c) ||
			(b[2][0] == c && b[1][1] == c && b[0][2] == c)
		return
	}
	cx, wx := stat('X')
	co, wo := stat('O')

	if wx {
		return !wo && cx-co == 1
	}

	if wo {
		return !wx && cx-co == 0
	}

	// don't have 2 winners,
	// and number of moves are valid
	return cx-co == 1 || cx-co == 0
}
