package main

// medium

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 0
	}
	s := 10
	for i := n; i > 1; i-- {
		v := 9
		for j := 1; j < i; j++ {
			v *= 10 - j
		}
		s += v
	}
	return s
}
