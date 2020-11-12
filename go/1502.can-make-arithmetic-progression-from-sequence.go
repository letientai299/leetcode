package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		t := arr[i] - arr[i-1]
		if t != d {
			return false
		}
	}

	return true
}
