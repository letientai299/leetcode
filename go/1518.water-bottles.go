package main

func numWaterBottles(n int, k int) int {
	s := 0
	empty := 0
	for n > 0 {
		s += n
		next := (n + empty) / k
		empty = (n + empty) % k
		n = next
	}
	return s
}
