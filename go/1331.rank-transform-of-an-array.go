package main

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	ranked := make(map[int]int)
	r := 1
	prev := tmp[0]
	ranked[prev] = 1
	for _, x := range tmp {
		if x > prev {
			r++
			prev = x
			ranked[x] = r
		}
	}

	for i, x := range arr {
		arr[i] = ranked[x]
	}

	return arr
}
