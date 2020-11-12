package main

func projectionArea(grid [][]int) int {
	s := 0
	n := len(grid)
	cols := make([]int, n)
	for y := 0; y < n; y++ {
		maxRow := 0
		for x := 0; x < n; x++ {
			if grid[y][x] != 0 {
				s++
			}

			if grid[y][x] > maxRow {
				maxRow = grid[y][x]
			}

			if grid[y][x] > cols[x] {
				cols[x] = grid[y][x]
			}
		}
		s += maxRow
	}

	for _, v := range cols {
		s += v
	}
	return s
}
