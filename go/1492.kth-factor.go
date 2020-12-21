package main

import (
	"math"
	"sort"
)

// medium

func kthFactor(n int, k int) int {
	if k == 1 {
		return 1
	}

	factors := []int{1, n}
	b := n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i != 0 {
			continue
		}

		b = n / i
		if b != i {
			factors = append(factors, i, b)
		} else {
			factors = append(factors, i)
		}
		if isPrime(b) {
			break
		}
	}

	sort.Ints(factors)

	if k > len(factors) {
		return -1
	}

	return factors[k-1]
}
