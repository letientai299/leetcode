package main

// medium

func maxRotateFunction(a []int) int {
	s := 0
	b := 0
	for i, x := range a {
		s += x
		b += i * x
	}

	m := b
	for i := len(a) - 1; i >= 0; i-- {
		b = b + s - len(a)*a[i]
		if b > m {
			m = b
		}
	}
	return m
}
