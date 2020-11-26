package main

// medium

func integerBreak(n int) int {
	// Base on AM-GM inequality, we need to break the integer into x group,
	// such that (n/x)^x is max.
	// Can't prove yet, but base on experiment with google chart,
	// the best x is n/e where e=2.7182818285
	// https://www.google.com/search?q=f%28x%29%3D%2815%2Fx%29%5Ex
	// That means our group component should be 3.
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	x := 3
	v := 1
	for n > x {
		v *= x
		n -= x
	}

	if v/x*(n+x) > v*n {
		v = v / x * (n + x)
	} else {
		v *= n
	}

	return v
}
