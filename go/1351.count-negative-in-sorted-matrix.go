package main

import "sort"

func countNegatives(grid [][]int) int {
	s := 0
	for _, arr := range grid {
		x := sort.Search(len(arr), func(i int) bool {
			return arr[i] < 0
		})
		for x < len(arr) && arr[x] >= 0 {
			x++
		}
		s += len(arr) - x
	}
	return s
}
