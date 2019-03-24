package main

import (
	"strconv"
	"testing"
)

func Test_isHappy(t *testing.T) {
	sumSquareDigit := func(n int) int {
		s := 0

		for n > 0 {
			d := n % 10
			s += d * d
			n /= 10
		}

		return s
	}
	isItHappy := func(n int) bool {
		m := make(map[int]bool)
		for {
			x := sumSquareDigit(n)
			if m[x] == false {
				m[x] = true
				n = x
			} else {
				return x == 1
			}
		}
	}

	for n := 0; n < 100; n++ {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			got := isHappy(n)
			actual := isItHappy(n)
			if got != actual {
				t.Errorf("isHappy(%v) = %v, want %v", n, got, actual)
			}
		})
	}
}
