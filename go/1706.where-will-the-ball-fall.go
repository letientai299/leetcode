package main

// medium
// 1706. Where Will the Ball Fall
// TODO (tai): can be faster
func findBall(grid [][]int) []int {
	m := len(grid)
	if m == 0 {
		return nil
	}
	n := len(grid[0])
	if n == 0 {
		return nil
	}

	line := make([][]int, n)

	for i := range line {
		line[i] = []int{i, i, i} // original index, current index, status
	}

	for y := 0; y < m; y++ {
		for i, b := range line {
			if b[2] == -1 {
				continue
			}

			x := b[1]

			direction := grid[y][x]
			if direction == 1 {
				if x != n-1 && grid[y][x+1] != -1 {
					b[1]++
				} else {
					b[2] = -1
				}
			} else {
				if x != 0 && grid[y][x-1] != 1 {
					b[1]--
				} else {
					b[2] = -1
				}
			}

			line[i] = b
		}
	}

	for _, b := range line {
		if b[2] == -1 {
			grid[m-1][b[0]] = -1
		} else {
			grid[m-1][b[0]] = b[1]
		}
	}

	return grid[m-1]
}
