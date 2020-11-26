package main

import "math"

// medium

func numSquares(n int) int {
	var find func(n int) int
	cached := make(map[int]int)
	find = func(n int) int {
		if n < 4 {
			return n
		}

		if v, ok := cached[n]; ok {
			return v
		}

		cap := int(math.Floor(math.Sqrt(float64(n))))
		min := math.MaxInt32
		for i := cap; i > cap/2; i-- {
			c := n/(i*i) + find(n%(i*i))
			if min > c {
				min = c
			}
		}
		cached[n] = min
		return min
	}

	return find(n)
}
