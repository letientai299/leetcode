package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	arr = arr[n/20 : n-n/20]
	s := 0
	for _, x := range arr {
		s += x
	}
	return float64(s) / float64(n-n/10)
}
