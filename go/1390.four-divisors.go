package main

import (
	"math"
)

// medium

// this one is 40ms. Faster solution use a precomputed list of primes.
func sumFourDivisors(nums []int) int {
	sum4 := func(n int) int {
		if n <= 5 {
			return 0 // they don't have 4 divisors
		}
		cnt := 2
		sum := n + 1
		for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
			d, m := n/i, n%i
			if m != 0 {
				continue
			}
			sum += i
			cnt++
			if d != i {
				sum += d
				cnt++
			}
			if cnt > 4 {
				return 0
			}
		}

		if cnt != 4 {
			return 0
		}

		return sum
	}

	s := 0
	for _, x := range nums {
		s += sum4(x)
	}
	return s
}
