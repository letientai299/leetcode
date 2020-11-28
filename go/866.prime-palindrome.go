package main

import (
	"math"
)

/*
 * @lc app=leetcode id=866 lang=golang
 * https://leetcode.com/problems/prime-palindrome/description/
 * Medium (21.77%)
 */

func primePalindrome(n int) int {
	switch {
	case n <= 2:
		return 2
	case n <= 3:
		return 3
	case n <= 5:
		return 5
	case n <= 7:
		return 7
	case n <= 11:
		return 11
	case n <= 101:
		return 101
	}
	ds := nextPotentialPalindromeNum(digits(n))
	n = fromDigits(ds)
	for !isPrime(n) {
		ds = nextPotentialPalindromeNum(ds)
		n = fromDigits(ds)
	}

	return n
}

func fromDigits(ds []int) int {
	n := 0
	for _, d := range ds {
		n = 10*n + d
	}
	return n
}

func digits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var ds []int
	for n > 0 {
		ds = append(ds, n%10)
		n /= 10
	}
	for i := 0; i < len(ds)/2; i++ {
		ds[i], ds[len(ds)-1-i] = ds[len(ds)-1-i], ds[i]
	}
	return ds
}

// find the next potential palindrome number that greater or equal to n.
// assuming len(ds) >= 3, because in our solution above handles special case of
// n<=101
func nextPotentialPalindromeNum(ds []int) []int {
	n := len(ds)
	if n == 0 {
		return nil
	}

	notPrimeByDigitRules := func(ds []int) bool {
		if ds[0] == 5 || ds[len(ds)-1]%2 == 0 {
			return true
		}

		sum := 0
		s11 := 0
		s101 := 0
		for i, x := range ds {
			sum += x
			if i%2 == 0 {
				s11 += x
			} else {
				s11 -= x
			}
			if i%4 <= 1 {
				s101 += x
			} else {
				s101 -= x
			}
		}
		return sum%3 == 0 || s11%11 == 0 || s101%101 == 0
	}

	if ds[0]%2 == 0 {
		ds[0]++
		ds[n-1] = ds[0]
		for i := 1; i < n-1; i++ {
			ds[i] = 0
		}
		if notPrimeByDigitRules(ds) {
			return nextPotentialPalindromeNum(ds)
		}
		return ds
	}

	if n%2 != 0 {
		for i := n / 2; i >= 0; i-- {
			if ds[i] < 9 {
				ds[i]++
				ds[n-1-i] = ds[i]
				sumDigit := ds[i]
				if n%2 == 0 {
					sumDigit += ds[i]
				}
				for j := i + 1; j <= n/2; j++ {
					ds[j], ds[n-1-j] = 0, 0
				}
				previousDigit := ds[i]
				same := true
				for i--; i >= 0; i-- {
					ds[n-1-i] = ds[i]
					sumDigit += 2 * ds[i]
					if ds[i] != previousDigit {
						same = false
					}
				}

				if same || notPrimeByDigitRules(ds) {
					return nextPotentialPalindromeNum(ds)
				}

				return ds
			}
		}
	}

	// palindrome with even number of digits is divisible to 11
	if n%2 == 0 {
		ds = append([]int{1}, ds...)
	} else {
		ds = append([]int{1, 0}, ds...)
	}
	ds[n] = 1
	for i := 1; i < n; i++ {
		ds[i] = 0
	}

	if notPrimeByDigitRules(ds) {
		return nextPotentialPalindromeNum(ds)
	}
	return ds
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
