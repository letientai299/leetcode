package main

func max(a int, arr ...int) int {
	for _, v := range arr {
		if a < v {
			a = v
		}
	}

	return a
}

func min(a int, arr ...int) int {
	for _, v := range arr {
		if a > v {
			a = v
		}
	}

	return a
}

func gcd(a, b int) int {
	for a != 0 && a != 1 {
		a, b = b%a, a
	}

	if a == 0 {
		return b
	}

	return 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func numPalindrome(n int) bool {
	v := n
	r := 0
	for v > 0 {
		r = 10*r + v%10
		v /= 10
	}
	return r == n
}

