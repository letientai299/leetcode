package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numsSameConsecDiff(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want []int
	}{
		{n: 2, k: 1, want: []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
		{n: 2, k: 0, want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
		{n: 3, k: 7, want: []int{181, 292, 707, 818, 929}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numsSameConsecDiff(tt.n, tt.k)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
