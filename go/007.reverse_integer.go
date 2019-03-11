package main

import "math"

// NOTE:
// using recursion (reverse(x) = - reverse(-x)) for negative number make this solution slow as fuck,
// from 4ms (100%) to 8ms (10.7%).
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	s := int64(0)
	a := int64(x)
	for a > 0 {
		lastDigit := a % 10
		s = 10*s + lastDigit
		if s > math.MaxInt32 {
			return 0
		}

		a = a / 10
	}
	return sign * int(s)
}
