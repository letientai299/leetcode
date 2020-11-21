package main

// medium

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			up := -1
			if y > 0 {
				up = grid[y-1][x]
			}
			left := -1
			if x > 0 {
				left = grid[y][x-1]
			}

			if (up != -1 && left != -1 && up > left) || (up == -1) {
				up = left
			}

			if up != -1 {
				grid[y][x] += up
			}
		}
	}

	return grid[m-1][n-1]
}
