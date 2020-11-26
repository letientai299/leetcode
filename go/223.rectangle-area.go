package main

// medium

func computeArea(a int, b int, c int, d int, e int, f int, g int, h int) int {
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}

	s1 := abs(c-a) * abs(d-b)
	s2 := abs(g-e) * abs(h-f)
	if !isRectangleOverlap([]int{a, b, c, d}, []int{e, f, g, h}) {
		return s1 + s2
	}

	within := func(x, y int, m, n, p, q int, trick ...bool) bool {
		if len(trick) == 0 {
			return m <= x && x <= p && n <= y && y <= q
		}

		return m < x && x < p && n < y && y < q
	}

	// contains
	if within(a, b, e, f, g, h) && within(c, d, e, f, g, h) {
		return s2
	}

	if within(e, f, a, b, c, d) && within(g, h, a, b, c, d) {
		return s1
	}

	// point

	if within(a, b, e, f, g, h) && within(g, h, a, b, c, d) {
		return s1 + s2 - abs(a-g)*abs(b-h)
	}

	if within(c, d, e, f, g, h) && within(e, f, a, b, c, d) {
		return s1 + s2 - abs(c-e)*abs(d-f)
	}

	if within(c, b, e, f, g, h) && within(e, h, a, b, c, d) {
		return s1 + s2 - abs(c-e)*abs(b-h)
	}

	if within(a, d, e, f, g, h) && within(g, f, a, b, c, d) {
		return s1 + s2 - abs(a-g)*abs(d-f)
	}

	// side

	if within(a, b, e, f, g, h, true) && within(c, b, e, f, g, h, true) {
		return s1 + s2 - abs(a-c)*abs(b-h)
	}

	if within(c, d, e, f, g, h, true) && within(a, d, e, f, g, h, true) {
		return s1 + s2 - abs(a-c)*abs(d-f)
	}

	if within(a, b, e, f, g, h, true) && within(a, d, e, f, g, h, true) {
		return s1 + s2 - abs(d-b)*abs(a-g)
	}

	if within(c, b, e, f, g, h, true) && within(c, d, e, f, g, h, true) {
		return s1 + s2 - abs(d-b)*abs(c-e)
	}

	// side

	if within(e, f, a, b, c, d, true) && within(g, f, a, b, c, d, true) {
		return s1 + s2 - abs(e-g)*abs(f-d)
	}

	if within(e, h, a, b, c, d, true) && within(g, h, a, b, c, d, true) {
		return s1 + s2 - abs(e-g)*abs(h-b)
	}

	if within(e, f, a, b, c, d, true) && within(e, h, a, b, c, d, true) {
		return s1 + s2 - abs(f-h)*abs(e-c)
	}

	if within(g, f, a, b, c, d, true) && within(g, h, a, b, c, d, true) {
		return s1 + s2 - abs(f-h)*abs(g-a)
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	return s1 + s2 - min(abs(a-c), abs(e-g))*min(abs(b-d), abs(f-h))
}
