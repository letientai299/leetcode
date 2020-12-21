package main

// medium
// https://leetcode.com/problems/number-of-substrings-with-only-1s/

func numSub(s string) int {
	prev := -1
	b := '0'
	r := 0
	mod := int(1e9 + 7)
	for i, c := range s {
		if c == b {
			continue
		}

		if c == '1' {
			prev = i
		} else {
			r += ((i - prev) * (i - prev + 1) / 2) % mod
			r %= mod
		}
		b = c
	}

	if b == '1' {
		i := len(s)
		r += ((i - prev) * (i - prev + 1) / 2) % mod
		r %= mod
	}

	return r
}
