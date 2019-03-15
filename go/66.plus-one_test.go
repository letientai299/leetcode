package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	tests := []int{109, 190, 199, 909, 99999991, 199999999, 0, 10, 1, 2, 3, 99, 91, 9999}

	for _, num := range tests {
		t.Run(strconv.Itoa(num), func(t *testing.T) {
			digits := makeDigits(num)
			x := makeNum(digits)
			fmt.Println(digits)
			fmt.Println(x)

			got := plusOne(digits)
			fmt.Println(got)
			actual := makeNum(got)
			assert.Equal(t, num, actual-1)
		})
	}
}

func makeDigits(n int) []int {
	var r []int
	base := 10
	for base <= n {
		base *= 10
	}
	base /= 10

	for {
		d := n / base
		r = append(r, d)
		n = n - d*base
		base /= 10
		if base == 0 {
			break
		}
	}

	return r
}

func makeNum(digits []int) int {
	base := 1
	n := 0
	for i := len(digits) - 1; i >= 0; i-- {
		n += digits[i] * base
		base *= 10
	}
	return n
}
