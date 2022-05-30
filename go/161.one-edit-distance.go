package main

// One Edit Distance
//
// Medium
//
// https://leetcode.com/problems/one-edit-distance/
func isOneEditDistance(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m > n+1 || n > m+1 {
		return false
	}

	if m == n {
		// check for replace
		diff := 0
		for i, x := range s {
			y := t[i]
			if byte(x) != y {
				diff++
			}
		}

		return diff == 1
	}

	if m > n {
		m, n = n, m
		s, t = t, s
	}

	// check for insert case, in which we only insert to t

	i := 0
	for ; i < m && s[i] == t[i]; i++ {
	}

	for ; i < m && s[i] == t[i+1]; i++ {
	}

	return i >= m
}
