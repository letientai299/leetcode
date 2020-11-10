package main

func minStartValue(a []int) int {
	min := 1 << 31
	sum := 0
	for _, n := range a {
		sum += n
		if sum < min {
			min = sum
		}
	}
	min = 1 - min
	if min < 1 {
		return 1
	}
	return min
}
