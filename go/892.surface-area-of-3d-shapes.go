package main

func surfaceArea(grid [][]int) int {
	n := len(grid)
	s := n * n * 2
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := grid[y][x]
			if v == 0 {
				s -= 2
				continue
			}

			// up
			if y == 0 {
				s += v
			} else if grid[y-1][x] < v {
				s += v - grid[y-1][x]
			}

			// down
			if y == n-1 {
				s += v
			} else if grid[y+1][x] < v {
				s += v - grid[y+1][x]
			}

			// left
			if x == 0 {
				s += v
			} else if grid[y][x-1] < v {
				s += v - grid[y][x-1]
			}

			// right
			if x == n-1 {
				s += v
			} else if grid[y][x+1] < v {
				s += v - grid[y][x+1]
			}
		}
	}

	return s
}
