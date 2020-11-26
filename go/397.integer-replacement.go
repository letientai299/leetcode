package main

// medium

func integerReplacement(n int) int {
	cached := make(map[int]int)
	var find func(int) int

	find = func(n int) (c int) {
		if n == 1 {
			return 0
		}

		if v, ok := cached[n]; ok {
			return v
		}

		if n%2 == 0 {
			c = 1 + find(n/2)
			cached[n] = c
			return c
		}

		a, b := find(n+1), find(n-1)
		if a > b {
			c = b + 1
		} else {
			c = a + 1
		}

		cached[n] = c
		return c
	}

	return find(n)
}
