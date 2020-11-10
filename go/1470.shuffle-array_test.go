package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shuffle(t *testing.T) {
	for n := 1; n < 10; n++ {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			var a []int
			for i := 0; i < 2*n; i++ {
				a = append(a, i)
			}

			got := shuffle(a, n)
			expected := make([]int, 2*n)
			for i := 0; i < n; i++ {
				expected[2*i] = i
				expected[2*i+1] = i + n
			}

			for i := 0; i < n; i++ {
				if !assert.Truef(t, got[2*i] == i, "got=%v, want=%v", got, expected) {
					return
				}
			}

			for i := 0; i < n; i++ {
				if !assert.Truef(t, got[2*i+1] == n+i, "got=%v, want=%v", got, expected) {
					return
				}
			}
		})
	}
}
