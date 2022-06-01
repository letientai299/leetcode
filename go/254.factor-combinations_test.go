package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFactors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			n: 12,
			want: [][]int{
				{2, 6},
				{2, 2, 3},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getFactors(tt.n), "getFactors(%v)", tt.n)
		})
	}
}
