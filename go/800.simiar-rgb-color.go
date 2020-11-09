package main

import "fmt"

// https://leetcode.com/problems/similar-rgb-color/
func similarRGB(s string) string {
	val := func(high, low uint8) int {
		if low >= '0' && low <= '9' {
			low -= '0'
		} else {
			low = low - 'a' + 10
		}

		if high >= '0' && high <= '9' {
			high -= '0'
		} else {
			high = high - 'a' + 10
		}
		return int(high<<4 + low)
	}

	min := func(a, b, c int) int {
		if a <= b && a <= c {
			return a
		}
		if b <= a && b <= c {
			return b
		}
		return c
	}

	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	nearest := func(n int) int {
		i := n >> 4
		a, b, c := (i-1)*17, i*17, (i+1)*17
		x, y, z := abs(a-n), abs(b-n), abs(c-n)
		m := min(x, y, z)
		switch m {
		case x:
			return a
		case y:
			return b
		default:
			return c
		}
	}

	a := nearest(val(s[1], s[2]))
	b := nearest(val(s[3], s[4]))
	c := nearest(val(s[5], s[6]))
	return fmt.Sprintf("#%02x%02x%02x", a, b, c)
}
