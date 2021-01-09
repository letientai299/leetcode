package main

import "math"

// medium
// 1376. Time Needed to Inform All Employees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	if n <= 1 {
		return 0
	}
	m := make([][]int, n+1)
	for i, v := range manager {
		v = (v + n + 1) % (n + 1)
		m[v] = append(m[v], i)
	}

	var inform func(x int) int
	inform = func(x int) int {
		if len(m[x]) == 0 {
			return 0
		}

		t := math.MinInt32
		for _, y := range m[x] {
			v := inform(y)
			if v > t {
				t = v
			}
		}
		return t + informTime[x]
	}
	return inform(headID)
}
