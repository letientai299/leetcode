package main

func maxScore(s string) int {
	l := 0
	left := make([]int, len(s))
	for i := range s {
		if s[i] == '0' {
			left[i] = l + 1
			l = left[i]
		} else {
			left[i] = l
		}
	}

	m := 0
	r := 0
	for i := len(s) - 1; i >= 1; i-- {
		l := left[i-1]
		if s[i] == '1' {
			r++
		}
		if l+r > m {
			m = l + r
		}
	}

	return m
}
