package main

// Snakes and Ladders
//
// Medium
//
// https://leetcode.com/problems/snakes-and-ladders/
//
// You are given an `n x n` integer matrix `board` where the cells are labeled
// from `1` to `n2` in a [**Boustrophedon
// style**](https://en.wikipedia.org/wiki/Boustrophedon) starting from the
// bottom left of the board (i.e. `board[n - 1][0]`) and alternating direction
// each row.
//
// You start on square `1` of the board. In each move, starting from square
// `curr`, do the following:
//
// - Choose a destination square `next` with a label in the range `[curr + 1,
// min(curr + 6, n2)]`.
//
//
//     - This choice simulates the result of a standard **6-sided die roll**:
// i.e., there are always at most 6 destinations, regardless of the size of the
// board.
// - If `next` has a snake or ladder, you **must** move to the destination of
// that snake or ladder. Otherwise, you move to `next`.
// - The game ends when you reach the square `n2`.
//
// A board square on row `r` and column `c` has a snake or ladder if
// `board[r][c] != -1`. The destination of that snake or ladder is
// `board[r][c]`. Squares `1` and `n2` do not have a snake or ladder.
//
// Note that you only take a snake or ladder at most once per move. If the
// destination to a snake or ladder is the start of another snake or ladder, you
// do **not** follow the subsequent snake or ladder.
//
// - For example, suppose the board is `[[-1,4],[-1,3]]`, and on the first move,
// your destination square is `2`. You follow the ladder to square `3`, but do
// **not** follow the subsequent ladder to `4`.
//
// Return _the least number of moves required to reach the square_ `n2` _. If it
// is not possible to reach the square, return_`-1`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2018/09/23/snakes.png)
//
// ```
// Input: board =
// [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]
// Output: 4
// Explanation:
// In the beginning, you start at square 1 (at row 5, column 0).
// You decide to move to square 2 and must take the ladder to square 15.
// You then decide to move to square 17 and must take the snake to square 13.
// You then decide to move to square 14 and must take the ladder to square 35.
// You then decide to move to square 36, ending the game.
// This is the lowest possible number of moves to reach the last square, so
// return 4.
//
// ```
//
// **Example 2:**
//
// ```
// Input: board = [[-1,-1],[-1,3]]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `n == board.length == board[i].length`
// - `2 <= n <= 20`
// - `grid[i][j]` is either `-1` or in the range `[1, n2]`.
// - The squares labeled `1` and `n2` do not have any ladders or snakes.
func snakesAndLadders(board [][]int) int {
	n := len(board)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp {
			dp[i][j] = -1
		}
	}

	location := func(v int) (int, int) {
		v--
		y := n - v/n - 1
		x := v % n
		if (v/n)%2 == 1 {
			x = n - x - 1
		}
		return x, y
	}

	x, y := location(1)
	dp[y][x] = 0

	for i := 2; i <= n*n; i++ {
		x, y = location(i)

		minStep := -1
		for j := 1; j <= 6; j++ {
			pre := i - j
			if pre >= 1 && pre <= n*n {
				preX, preY := location(pre)
				preStep := dp[preY][preX]
				// preStep == -1 mean that cell is unreachable, or it's a snake,
				if (preStep != -1 && preStep < minStep) || minStep == -1 {
					minStep = preStep
				}
			}
		}

		if minStep >= 0 {
			minStep++
		}

		v := board[y][x]
		if dp[y][x] == -1 || dp[y][x] > minStep {
			dp[y][x] = minStep
			if v != -1 && v < i {
				// here, the min way to go to this cell is from its 6 previous cells.
				// adn this cell is a snake. Hence, we don't go to this cell.
				dp[y][x] = -1
			} else {
				dp[y][x] = minStep
			}
		}

		if v == -1 {
			continue
		}

		nextX, nextY := location(v)
		if dp[nextY][nextX] == -1 || dp[nextY][nextX] > minStep {
			dp[nextY][nextX] = minStep
		}
	}

	printMat(dp)
	printMat(board)
	x, y = location(n * n)
	return dp[y][x]
}

// TODO (tai): 193 / 212 test cases passed, fix the bugs, or perhaps rewrite the code
