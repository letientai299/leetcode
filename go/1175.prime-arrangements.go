package main

func numPrimeArrangements(n int) int {
	if n <= 2 {
		return 1
	}

	primes := []int{2}
	for i := 3; i <= n; i++ {
		isPrime := true
		for _, x := range primes {
			if x > i>>1 {
				break
			}
			if i%x == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	r := 1
	for i := 2; i <= len(primes); i++ {
		r *= i
		r %= 1e9 + 7
	}

	for i := 2; i <= n-len(primes); i++ {
		r *= i
		r %= 1e9 + 7
	}

	return r
}
