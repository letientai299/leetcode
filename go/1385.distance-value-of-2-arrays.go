package main

func findTheDistanceValue(a []int, b []int, d int) int {
	s := 0
	for _, x := range a {
		for _, y := range b {
			dx := x - y
			if dx < 0 {
				dx = -dx
			}
			if dx <= d {
				s--
				break
			}
		}
		s++
	}
	return s
}
