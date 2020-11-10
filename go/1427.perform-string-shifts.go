package main

func stringShift(s string, shift [][]int) string {
	step := 0
	for _, a := range shift {
		if a[0] == 0 {
			step -= a[1]
		} else {
			step += a[1]
		}
	}

	step %= len(s)

	if step > 0 {
		return s[len(s)-step:] + s[:len(s)-step]
	}

	return s[-step:] + s[:-step]
}
