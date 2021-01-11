package main

// medium
// https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	var v float64
	v = 1
	for n > 0 {
		t := x
		i := 1
		for i < n/2 {
			t *= t
			i *= 2
		}
		v *= t
		n -= i
	}
	return v
}
