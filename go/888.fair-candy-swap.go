package main

import "sort"

// https://leetcode.com/problems/fair-candy-swap/

func fairCandySwap(a []int, b []int) []int {
	diff := 0
	for _, x := range a {
		diff += x
	}
	for _, x := range b {
		diff -= x
	}

	sort.Ints(a)
	sort.Ints(b)
	var swap bool
	if diff < 0 {
		swap = true
		diff = -diff
		a, b = b, a
	}

	x, y := 0, 0

	diff = diff / 2

	for {
		d := a[x] - b[y]
		if d == diff {
			break // guaranteed to have answer
		}

		if d < diff {
			x++
		} else {
			y++
		}
	}

	if swap {
		a, b = b, a
		x, y = y, x
	}

	return []int{a[x], b[y]}
}
