package main

// medium

// TODO: subscribe leetcode and submit this.
func maxA(n int) int {
	if n <= 3 {
		return n
	}

	n -= 3
	r := 3
	t := 1
	for n >= 5 {
		n -= 5
		t = r
		r *= 4
	}

	switch n {
	case 1:
		r += t
	case 2:
		r += 2 * t
	case 3:
		r *= 2
	case 4:
		r *= 3
	}

	return r
}
