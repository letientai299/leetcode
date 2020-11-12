package main

func countOdds(low int, high int) int {
	s := (high - low) / 2
	if low%2 == 1 || high%2 == 1 {
		s++
	}
	return s
}
