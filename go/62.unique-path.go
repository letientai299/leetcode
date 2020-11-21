package main

// medium
// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	if m < n {
		m, n = n, m
	}
	m, n = m+n-2, n-1
	v := uint64(1)
	for i := m; i > m-n; i-- {
		v *= uint64(i)
	}
	for i := 2; i <= n; i++ {
		v /= uint64(i)
	}
	return int(v)
}
