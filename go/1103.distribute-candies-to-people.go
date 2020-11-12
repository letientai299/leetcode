package main

func distributeCandies(n int, p int) []int {
	round := 0
	x := p * (p + 1) / 2
	for n > p*p*round+x {
		n -= p*p*round + x
		round++
	}
	res := make([]int, p)
	for i := range res {
		res[i] = p*round*(round-1)/2 + (i+1)*round
		if n >= round*p+(i+1) {
			res[i] += round*p + (i + 1)
			n -= round*p + (i + 1)
		} else if n > 0 {
			res[i] += n
			n = 0
		}
	}
	return res
}
