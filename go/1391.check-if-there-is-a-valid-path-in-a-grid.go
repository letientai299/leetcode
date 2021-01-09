package main

// medium
// https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid/

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	if m == 0 {
		return false
	}

	n := len(grid[0])
	if n == 0 {
		return false
	}

	if m == 1 && n == 1 {
		return true
	}

	x, y := 0, 0
	i, j := 0, 0
	switch grid[y][x] {
	case 1:
		i = 1
	case 2:
		j = 1
	case 3:
		j = 1
	case 4:
		if m == 1 {
			i = 1
		} else if n == 1 {
			j = 1
		} else if grid[0][1] == 2 || grid[0][1] == 6 {
			j = 1
		} else if grid[1][0] == 1 || grid[1][0] == 3 {
			i = 1
		} else {
			// try go left first
			i = 1
		}
	case 5:
		return false
	case 6:
		i = 1
	}

	attempt := 0
again:
	ok := true
	for !(x < 0 || x >= n || y < 0 || y >= m) && ok {
		v := grid[y][x]
		switch {
		case v == 1:
			ok = j == 0
			j = 0
		case v == 2:
			ok = i == 0
			i = 0
		case v == 3:
			if x != 0 || y != 0 {
				ok = !(i == -1 || j == 1)
				i, j = j, i
			}
		case v == 4:
			if x != 0 || y != 0 {
				ok = !(i == 1 || j == 1)
				i, j = -j, -i
			}
		case v == 5:
			ok = !(i == -1 || j == -1)
			i, j = -j, -i
		case v == 6:
			if x != 0 || y != 0 {
				ok = !(i == 1 || j == -1)
				i, j = j, i
			}
		}
		if (x == n-1) && (y == m-1) {
			break
		}
		x += i
		y += j
	}

	ok = ok && !(x < 0 || x >= n || y < 0 || y >= m)
	if !ok && grid[0][0] == 4 && attempt == 0 {
		x, y = 0, 0
		j = 1
		i = 0
		attempt++
		goto again
	}

	return ok && x == n-1 && y == m-1
}
