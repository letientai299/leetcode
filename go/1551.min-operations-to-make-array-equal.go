package main

// medium
// https://leetcode.com/problems/minimum-operations-to-make-array-equal/

func minOperations1551(n int) int {
	// even: 1 + 3 + ... + (n-1)
	// odd: 0 + 2 + ... + (n -1)
	return n / 2 * (n + n%2) / 2
}
