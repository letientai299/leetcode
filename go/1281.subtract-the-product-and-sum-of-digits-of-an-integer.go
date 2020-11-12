package main

func subtractProductAndSum(n int) int {
	p := 1
	s := 0
	for n > 0 {
		p *= n % 10
		s += n % 10
		n /= 10
	}
	return p - s
}
