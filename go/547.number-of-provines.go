package main

// medium
// https://leetcode.com/problems/number-of-provinces/

func findCircleNum(m [][]int) int {
	n := len(m)
	if n == 0 {
		return 0
	}

	res := 0
	visited := make(map[int]struct{}, n)
	var clear func(i int)
	clear = func(i int) {
		for j, v := range m[i] {
			if v == 0 {
				continue
			}

			m[i][j] = 0
			if _, ok := visited[j]; ok {
				continue
			}

			visited[j] = struct{}{}
			clear(j)
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		visited[i] = struct{}{}
		clear(i)
		res++
	}

	return res
}
