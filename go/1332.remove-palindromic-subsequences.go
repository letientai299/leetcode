package main

func removePalindromeSub(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return 2
		}
	}
	return 1
}
