package main

// medium

func angleClock(hour int, minutes int) float64 {
	h := float64(hour)
	m := float64(minutes)
	ha := 30*h + m/2
	ma := m * 6

	x := ha - ma
	if x < 0 {
		x = -x
	}

	if x > 360-x {
		return 360 - x
	}
	return x
}
