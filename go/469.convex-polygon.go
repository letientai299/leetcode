package main

// medium

func isConvex(ps [][]int) bool {
	n := len(ps)

	dx, dy := ps[1][0]-ps[0][0], ps[1][1]-ps[0][1]
	k := 1
	for i := 2; i < n; i++ {
		a := ps[(i+n)%n]
		b := ps[(i-1+n)%n]
		cx, cy := b[0]-a[0], b[1]-a[1]
		if dx*cy != cx*dy {
			k++
		}
		ps[k] = ps[i%n]
		dx, dy = cx, cy
	}

	ps = ps[:k+1]
	n = k + 1

	if (ps[k][0]-ps[k-1][0])*(ps[0][1]-ps[k][1]) == (ps[k][1]-ps[k-1][1])*(ps[0][0]-ps[k][0]) {
		ps = ps[:k]
		n--
	}

	if (ps[n-1][0]-ps[0][0])*(ps[0][1]-ps[1][1]) == (ps[n-1][1]-ps[0][1])*(ps[0][0]-ps[1][0]) {
		ps = ps[1:]
		n--
	}

	sign := func(a, b, p []int) int {
		return (b[0]-a[0])*(p[1]-a[1]) - (b[1]-a[1])*(p[0]-a[0])
	}

	// remove points that's on the straight line with other points
	for i := 0; i < len(ps); i++ {
		d := ps[i]
		for j := (i + 1) % n; j < (i+n-2)%n; j++ {
			a := ps[(j+n)%n]
			b := ps[(j+1+n)%n]
			c := ps[(j+2+n)%n]
			sig1 := sign(a, b, d)
			sig2 := sign(b, c, d)
			if (sig1 > 0) == (sig2 > 0) {
				continue
			}

			return false
		}
	}
	return true
}
