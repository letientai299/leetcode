package main

// medium
// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/

func concatenatedBinary(n int) int {
	mod := int(1e9 + 7)
	r := 0
	shift := 0
	for i := 1; i <= n; i++ {
		if (i-1)&i == 0 {
			shift++
		}
		r = (r<<shift + i) % mod
	}
	return r
}
