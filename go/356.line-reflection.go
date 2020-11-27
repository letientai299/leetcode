package main

import (
	"math"
	"sort"
)

// medium

func isReflected(points [][]int) bool {
	n := len(points)
	if n <= 1 {
		return true
	}

	min, max := math.MaxInt32, math.MinInt32
	ps := make([][]int, 0, len(points))
	ms := make(map[[2]int]bool, len(points))
	for _, p := range points {
		if min > p[0] {
			min = p[0]
		}
		if max < p[0] {
			max = p[0]
		}

		cp := [2]int{p[0], p[1]}
		if !ms[cp] {
			ps = append(ps, p)
			ms[cp] = true
		}
	}
	mid := min + max

	sort.Slice(ps, func(i, j int) bool {
		a, b := ps[i], ps[j]
		if a[0] == b[0] {
			if 2*a[0] > mid {
				return a[1] < b[1]
			}
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})

	n = len(ps)

	if n <= 1 {
		return true
	}

	for i := 0; i <= (len(ps)+1)/2; i++ {
		if mid != ps[i][0]+ps[n-1-i][0] {
			return false
		}

		if mid/2 != ps[i][0] && ps[i][1] != ps[n-1-i][1] {
			return false
		}
	}

	return true
}
