package main

// this is in-place merge, part of merge sort.
func shuffle(a []int, n int) []int {
	for i := 0; i < n; i++ {
		v := a[n+i]
		for x := n + i; x > 2*i+1; x-- {
			a[x] = a[x-1]
		}
		a[2*i+1] = v
	}
	return a
}
