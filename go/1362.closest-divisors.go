package main

import "math"

// medium

func closestDivisors(num int) []int {
	a, b := num+1, num+2
	find := func(x int) []int {
		s := int(math.Floor(math.Sqrt(float64(x))))
		for x%s != 0 {
			s--
		}
		return []int{x / s, s}
	}

	d1 := find(a)
	d2 := find(b)
	if d1[0]-d1[1] < d2[0]-d2[1] {
		return d1
	}

	return d2
}
