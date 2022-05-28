package main

import "sort"

func countElements1426(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	sort.Ints(arr)
	s := 0

	prevCount := 1
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == prev {
			prevCount++
			continue
		}

		if arr[i] == prev+1 {
			s += prevCount
		}
		prevCount = 1
		prev = arr[i]
	}
	return s
}
