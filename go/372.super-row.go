package main

// medium

func superPow(A int, b []int) int {
	const c = 1337
	a := uint64(A % c)
	if a <= 1 {
		return int(a)
	}
	if a%c == 0 {
		return 0
	}

	p := uint64(1)
	for _, x := range b {
		if p != 1 {
			p = (((p * p) % c) *
				((p * p) % c) *
				((p * p) % c) *
				((p * p) % c) *
				((p * p) % c)) % c
		}

		v := uint64(1)
		for i := uint64(0); i < uint64(x); i++ {
			v = (v * a) % c
		}

		p = (p * v) % c
	}

	return int(p)
}
