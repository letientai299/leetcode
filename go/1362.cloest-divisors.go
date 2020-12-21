package main

import "math"

// medium

// TODO (tai): can be faster, this is just 50%
func closestDivisors(num int) []int {
	a, b := num+1, num+2
	find := func(x int) []int {
		s := int(math.Sqrt(float64(x)))
		if s*s == x {
			return []int{s, s}
		}

		i := 0
		s++
		if x%2 != s%2 {
			s--
		}

		for s+i <= x && i < s {
			if x%(s-i) == 0 {
				return []int{x / (s - i), s - i}
			}
			i++
		}

		return []int{1, x}
	}

	if a%2 != 0 {
		if isPrime(a) {
			return find(b)
		}
	} else if isPrime(b) {
		return find(a)
	}

	d1 := find(a)
	d2 := find(b)
	if abs(d1[0]-d1[1]) < abs(d2[0]-d2[1]) {
		return d1
	}

	return d2
}
