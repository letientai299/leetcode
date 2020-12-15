package main

import "strconv"

// medium

func simplifiedFractions(n int) []string {
	var res []string

	for i := 2; i <= n; i++ {
		res = append(res, "1/"+strconv.Itoa(i))
	}

	var primes []int
	primeFactors := func(x int) []int {
		var factors []int
		for _, v := range primes {
			if x%v == 0 {
				factors = append(factors, v)
			}
		}

		return factors
	}

	for i := 2; i < n; i++ {
		factors := primeFactors(i)
		if len(factors) == 0 {
			primes = append(primes, i)
			factors = append(factors, i)
		}

		for j := i + 1; j <= n; j++ {
			ok := true
			for _, f := range factors {
				if j%f == 0 {
					ok = false
					break
				}
			}
			if ok {
				res = append(res, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}

	return res
}
