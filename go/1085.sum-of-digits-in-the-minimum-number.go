package main

// https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/

func sumOfDigits(a []int) int {
	m := a[0]
	for _, x := range a {
		if m > x {
			m = x
		}
	}

	s := 0
	for m > 0 {
		s += m % 10
		m /= 10
	}
	return (s + 1) % 2
}
