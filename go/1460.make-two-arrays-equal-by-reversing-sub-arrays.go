package main

import "sort"

func canBeEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}
