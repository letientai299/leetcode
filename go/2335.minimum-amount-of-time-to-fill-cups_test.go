package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fillCups(t *testing.T) {
	tests := []struct {
		amount []int
		want   int
	}{
		{
			amount: []int{4, 5, 0},
			want:   5,
		},

		{
			amount: []int{1, 4, 2},
			want:   4,
		},

		{
			amount: []int{5, 4, 4},
			want:   7,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equalf(t, tt.want, fillCups(tt.amount), "fillCups(%v)", tt.amount)
		})
	}
}
