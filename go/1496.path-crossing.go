package main

func isPathCrossing(path string) bool {
	x, y := 0, 0
	m := make(map[[2]int]struct{})
	m[[2]int{0, 0}] = struct{}{}
	for _, c := range path {
		switch c {
		case 'N':
			y++
		case 'S':
			y--
		case 'W':
			x++
		case 'E':
			x--
		}
		a := [2]int{x, y}
		if _, ok := m[a]; ok {
			return true
		}
		m[a] = struct{}{}
	}
	return false
}
