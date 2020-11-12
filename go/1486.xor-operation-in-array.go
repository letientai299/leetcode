package main

func xorOperation(n int, start int) int {
	r := start
	for i := 1; i < n; i++ {
		r ^= start + 2*i
	}
	return r
}
