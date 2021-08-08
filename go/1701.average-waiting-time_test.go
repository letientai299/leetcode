package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_averageWaitingTime(t *testing.T) {
	tests := []struct {
		name      string
		customers [][]int
		want      float64
	}{
		{
			customers: [][]int{
				{2, 3}, {6, 3}, {7, 5}, {11, 3}, {15, 2}, {18, 1},
			},
			want: 4.16667,
		},

		{
			customers: [][]int{
				{1, 2},
				{2, 5},
				{4, 3},
			},
			want: 5.0,
		},

		{
			customers: [][]int{
				{5, 2},
				{5, 4},
				{10, 3},
				{20, 1},
			},
			want: 3.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageWaitingTime(tt.customers)
			assert.LessOrEqual(t, math.Abs(got-tt.want), 1e-5)
		})
	}
}
