package main

import "math"

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

func primeFactor(n int) []int {
	var r []int
	i := 2
	if n < i {
		return r
	}

	for n >= i {
		if n%i == 0 {
			r = append(r, i)
			for n%i == 0 {
				n /= i
			}
		}
		if i > 2 {
			i += 2
		} else {
			i++
		}
	}

	return r
}

func isPrime(a int) bool {
	n := a
	if n%2 == 0 {
		return false
	}

	if n == 3 || n == 5 || n == 7 || n == 11 {
		return true
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
