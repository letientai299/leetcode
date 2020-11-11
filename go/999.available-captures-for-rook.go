package main

/*
 * @lc app=leetcode id=999 lang=golang
 *
 * [999] Available Captures for Rook
 *
 * https://leetcode.com/problems/available-captures-for-rook/description/
 * Essy
 */
func numRookCaptures(board [][]byte) int {
	// find the Rook
	var x, y int
	for i, xs := range board {
		for j, c := range xs {
			if c == 'R' {
				x = j
				y = i
				break
			}
		}
	}

	capture := 0
	check := func(xDiff, yDiff int) {
		for i, j := x+xDiff, y+yDiff; j >= 0 && j < len(board) && i >= 0 && i < len(board); {
			switch board[j][i] {
			case '.': // empty
				i += xDiff
				j += yDiff
				continue
			case 'p':
				capture++
				return
			default:
				return
			}
		}
	}
	check(1, 0)
	check(-1, 0)
	check(0, 1)
	check(0, -1)
	return capture
}
