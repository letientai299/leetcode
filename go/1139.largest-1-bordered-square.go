package main

// Largest 1-Bordered Square
//
// Medium
//
// https://leetcode.com/problems/largest-1-bordered-square/
//
// Given a 2D `grid` of `0`s and `1`s, return the number of elements in the
// largest **square** subgrid that has all `1`s on its **border**, or `0` if
// such a subgrid doesn't exist in the `grid`.
//
// **Example 1:**
//
// ```
// Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
// Output: 9
//
// ```
//
// **Example 2:**
//
// ```
// Input: grid = [[1,1,0,0]]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= grid.length <= 100`
// - `1 <= grid[0].length <= 100`
// - `grid[i][j]` is `0` or `1`
func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	up := grid

	left := make([][]int, m)
	for i := range left {
		left[i] = make([]int, n)
	}

	left[0][0] = grid[0][0]
	best := left[0][0]

	for x := 1; x < n; x++ {
		if grid[0][x] == 1 {
			left[0][x] = left[0][x-1] + 1
			best = 1
		}
	}

	for y := 1; y < m; y++ {
		if grid[y][0] != 0 {
			up[y][0] = up[y-1][0] + 1
			left[y][0] = 1
			if best < 1 {
				best = 1
			}
		}

		for x := 1; x < n; x++ {
			if grid[y][x] == 0 {
				continue
			}

			up[y][x] = 1 + up[y-1][x]
			left[y][x] = 1 + left[y][x-1]
			a, b := up[y][x], left[y][x]

			side := a
			if side > b {
				side = b
			}

			for ; side > best; side-- {
				if grid[y-side+1][x-side+1] == 0 {
					continue
				}

				if left[y-side+1][x] < side || up[y][x-side+1] < side {
					continue
				}

				break
			}

			if best < side {
				best = side
			}
		}
	}

	return best * best
}
