package main

func average(salary []int) float64 {
	n := len(salary)
	if n < 3 {
		return 0
	}

	s := 0
	min := 1 << 31
	max := 0
	for _, x := range salary {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}

		s += x
	}
	s -= min + max
	return float64(s) / float64(n-2)
}
