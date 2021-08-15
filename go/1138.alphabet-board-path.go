package main

import "strings"

// Alphabet Board Path
//
// Medium
//
// https://leetcode.com/problems/alphabet-board-path/
//
// On an alphabet board, we start at position `(0, 0)`, corresponding to
// character `board[0][0]`.
//
// Here, `board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]`, as shown
// in the diagram below.
//
// ![](https://assets.leetcode.com/uploads/2019/07/28/azboard.png)
//
// We may make the following moves:
//
// - `'U'` moves our position up one row, if the position exists on the board;
// - `'D'` moves our position down one row, if the position exists on the board;
// - `'L'` moves our position left one column, if the position exists on the
// board;
// - `'R'` moves our position right one column, if the position exists on the
// board;
// - `'!'` adds the character `board[r][c]` at our current position `(r, c)` to
// the answer.
//
// (Here, the only positions that exist on the board are positions with letters
// on them.)
//
// Return a sequence of moves that makes our answer equal to `target` in the
// minimum number of moves.  You may return any path that does so.
//
// **Example 1:**
//
// ```
// Input: target = "leet"
// Output: "DDR!UURRR!!DDD!"
//
// ```
//
// **Example 2:**
//
// ```
// Input: target = "code"
// Output: "RR!DDRR!UUL!R!"
//
// ```
//
// **Constraints:**
//
// - `1 <= target.length <= 100`
// - `target` consists only of English lowercase letters.
func alphabetBoardPath(target string) string {
	pre := ' '
	x, y := 0, 0
	var sb strings.Builder
	for _, c := range target {
		if pre == 'z' && pre != c {
			y--
			sb.WriteByte('U')
		}
		c -= 'a'

		u, v := int(c%5), int(c/5)

		for x < u {
			sb.WriteByte('R')
			x++
		}

		for u < x {
			sb.WriteByte('L')
			x--
		}

		for y < v {
			sb.WriteByte('D')
			y++
		}

		for y > v {
			sb.WriteByte('U')
			y--
		}

		sb.WriteByte('!')
		x, y = u, v
		pre = c + 'a'
	}
	return sb.String()
}
