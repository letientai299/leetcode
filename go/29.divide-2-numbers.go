package main

// medium

func divide(a int, b int) int {
	if -1<<31 == a && b == -1 {
		return 1<<31 - 1
	}

	if a == 0 {
		return 0
	}

	if (a > 0) != (b > 0) {
		return -divide(-a, b)
	}

	if a < 0 {
		return divide(-a, -b)
	}

	n := 1
	c := b
	for c <= a {
		c += c
		n <<= 1
	}

	for c > a {
		c -= b
		n--
	}

	return n
}
