package main

func numberOfSteps(n int) int {
	s := 0
	for n > 0 {
		s += n%2 + 1
		n /= 2
		if n == 0 {
			s--
		}
	}
	return s
}
